package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"os"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/saadi925/email_marketing_api/internal/app/users"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenExpiry        = time.Hour * 1
	passwordResetTokenExpiry = time.Hour * 1
	refreshTokenExpiry       = time.Hour * 24 * 30
)

func signIn(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signInRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		err = validate.Struct(creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := queries.GetUserByEmail(r.Context(), creds.Email)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		if !user.Verified {
			utils.RespondError(w, http.StatusForbidden, "Account not verified. Please verify your email.")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		accessToken, err := generateJWT(users.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}, accessTokenExpiry)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error generating access token")
			return
		}

		refreshToken, err := generateJWT(users.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		}, refreshTokenExpiry)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error generating refresh token")
			return
		}

		// Store refresh token in the database
		_, err = queries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
			UserID:    user.ID,
			Token:     refreshToken,
			ExpiresAt: time.Now().Add(refreshTokenExpiry),
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error saving refresh token")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Expires:  time.Now().Add(accessTokenExpiry),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Expires:  time.Now().Add(refreshTokenExpiry),
			HttpOnly: true,
		})

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Signed in successfully",
		})
	}
}

func signUp(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signUpRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		err = validate.Struct(creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Validation failed")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		_, err = queries.CreateUser(r.Context(), database.CreateUserParams{
			Name:     creds.Name,
			Password: string(hashedPassword),
			Email:    creds.Email,
			Verified: false,
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error saving user")
			return
		}

		err = sendVerificationEmail(queries, creds.Email)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error sending verification email")
			return
		}
		utils.RespondJSON(w, http.StatusCreated, utils.MessageResponse{
			Message: "User has been created successfully. Please check your email to verify your account.",
		})
	}
}

func generateJWT(user users.User, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func verifyEmail(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		code := r.URL.Query().Get("code")

		// Validate email and code parameters
		if email == "" || code == "" {
			utils.RespondError(w, http.StatusBadRequest, "Email and verification code are required")
			return
		}

		// Fetch email verification record from the database
		emailVerify, err := queries.GetEmailVerifyByEmail(context.Background(), email)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error fetching email verification record")
			return
		}

		// Check if verification code matches
		if emailVerify.Code != code {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid verification code")
			return
		}

		// Update user's verification status in the users table
		err = queries.VerifyUserByEmail(context.Background(), email)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error verifying user")
			return
		}

		// Optionally, delete or mark the email verification record as used
		err = queries.DeleteEmailVerifyByEmail(context.Background(), email)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error deleting email verification record")
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Email verified successfully",
		})
	}
}

func forgotPassword(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := chi.URLParam(r, "email")

		user, err := queries.GetUserByEmail(r.Context(), email)
		if err != nil {
			utils.RespondError(w, http.StatusNotFound, "User not found")
			return
		}

		// Generate a password reset token
		token, err := generateJWT(users.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}, passwordResetTokenExpiry)
		if err != nil {
			fmt.Println("error generatring token :", err)
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Store the password reset token in the database
		err = queries.CreatePasswordResetToken(r.Context(), database.CreatePasswordResetTokenParams{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: time.Now().Add(passwordResetTokenExpiry),
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating password reset token")
			return
		}

		// Send the password reset email with the token
		err = sendPasswordResetEmail(user.Email, token)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error sending password reset email")
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Password reset email has been sent. Please check your email.",
		})
	}
}

func changePassword(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var changePasswordRequest struct {
			Token    string `json:"token" validate:"required"`
			Password string `json:"password" validate:"required,min=8"`
		}

		err := utils.ParseJSON(r, &changePasswordRequest)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		// Validate the token and fetch the associated user
		tokenInfo, err := queries.GetPasswordResetToken(r.Context(), changePasswordRequest.Token)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Check if the token is expired
		if tokenInfo.ExpiresAt.Before(time.Now()) {
			utils.RespondError(w, http.StatusUnauthorized, "Token has expired")
			return
		}

		// Generate hashed password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating password")
			return
		}

		// Update user's password
		err = queries.UpdateUserPassword(r.Context(), database.UpdateUserPasswordParams{
			ID:       tokenInfo.UserID,
			Password: string(hashedPassword),
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error updating password")
			return
		}

		// Delete the used password reset token
		err = queries.DeletePasswordResetToken(r.Context(), changePasswordRequest.Token)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error deleting password reset token")
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Password has been changed successfully",
		})
	}
}
