package auth

import (
	"net/http"
	"time"

	"os"

	"github.com/golang-jwt/jwt"
	"github.com/saadi925/email_marketing_api/internal/app/users"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpiry = time.Hour * 24

func signIn(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signInRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		user, err := queries.GetUserByEmail(r.Context(), creds.Email)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		tokenString, err := generateJWT(users.User{
			Name:     user.Name,
			Email:    user.Name,
			Password: user.Password,
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    tokenString,
			Expires:  time.Now().Add(tokenExpiry),
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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		_, err = queries.CreateUser(r.Context(), database.CreateUserParams{
			Name:     creds.Name,
			Password: string(hashedPassword),
			Email:    creds.Email,
		})
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error saving user")
			return
		}

		utils.RespondJSON(w, http.StatusCreated, utils.MessageResponse{
			Message: "User has been created successfully",
		})
	}
}

func verifyEmail(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement email verification logic here
	}
}

func resetPassword(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement password reset logic here
	}
}

func verifyOneTimeToken(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implement one-time token verification logic here
	}
}

func generateJWT(user users.User) (string, error) {
	claims := jwt.MapClaims{
		"userID":     user.ID,
		"isVerified": true,
		"exp":        time.Now().Add(tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
