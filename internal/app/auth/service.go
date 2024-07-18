// service.go
package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/saadi925/email_marketing_api/internal/app/users"
	"github.com/saadi925/email_marketing_api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenExpiry        = time.Hour * 1
	passwordResetTokenExpiry = time.Hour * 1
	refreshTokenExpiry       = time.Hour * 24 * 30
)

var (
	errAccountNotVerified        = errors.New("account not verified")
	errInvalidCredentials        = errors.New("invalid email or password")
	errInvalidVerificationCode   = errors.New("invalid verification code")
	errExpiredPasswordResetToken = errors.New("password reset token has expired")
)

type AuthService interface {
	SignIn(ctx context.Context, creds signInRequest) (string, string, error)
	SignUp(ctx context.Context, creds signUpRequest) error
	VerifyEmail(ctx context.Context, email string, code string) error
	ForgotPassword(ctx context.Context, email string) error
	ChangePassword(ctx context.Context, changePasswordRequest changePasswordRequest) error
}

type authService struct {
	db  *database.Queries
	jwt *jwt.SigningMethodHMAC
}

func NewAuthService(db *database.Queries) AuthService {
	return &authService{
		db:  db,
		jwt: jwt.SigningMethodHS256,
	}
}

func (s *authService) SignIn(ctx context.Context, creds signInRequest) (string, string, error) {

	user, err := s.db.GetUserByEmail(ctx, creds.Email)
	if err != nil {
		return "", "", err
	}

	if !user.Verified {
		return "", "", errAccountNotVerified
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return "", "", errInvalidCredentials
	}

	accessToken, err := generateJWT(users.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, accessTokenExpiry)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateJWT(users.User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}, refreshTokenExpiry)
	if err != nil {
		return "", "", err
	}

	// Store refresh token in the database
	_, err = s.db.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(refreshTokenExpiry),
	})
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) SignUp(ctx context.Context, creds signUpRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.db.CreateUser(ctx, database.CreateUserParams{
		Name:     creds.Name,
		Password: string(hashedPassword),
		Email:    creds.Email,
		Verified: false,
	})
	if err != nil {
		return err
	}

	err = sendVerificationEmail(s.db, creds.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) VerifyEmail(ctx context.Context, email string, code string) error {
	emailVerify, err := s.db.GetEmailVerifyByEmail(ctx, email)
	if err != nil {
		return err
	}

	if emailVerify.Code != code {
		return errInvalidVerificationCode
	}

	err = s.db.VerifyUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	err = s.db.DeleteEmailVerifyByEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	token, err := generateJWT(users.User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}, passwordResetTokenExpiry)
	if err != nil {
		return err
	}

	err = s.db.CreatePasswordResetToken(ctx, database.CreatePasswordResetTokenParams{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(passwordResetTokenExpiry),
	})
	if err != nil {
		return err
	}

	err = sendPasswordResetEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ChangePassword(ctx context.Context, changePasswordRequest changePasswordRequest) error {
	tokenInfo, err := s.db.GetPasswordResetToken(ctx, changePasswordRequest.Token)
	if err != nil {
		return err
	}

	if tokenInfo.ExpiresAt.Before(time.Now()) {
		return errExpiredPasswordResetToken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.db.UpdateUserPassword(ctx, database.UpdateUserPasswordParams{
		ID:       tokenInfo.UserID,
		Password: string(hashedPassword),
	})
	if err != nil {
		return err
	}

	err = s.db.DeletePasswordResetToken(ctx, changePasswordRequest.Token)
	if err != nil {
		return err
	}

	return nil
}
func generateJWT(user users.User, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
