package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/saadi925/email_marketing_api/internal/app/users"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

func SignInHandler(service AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signInRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		err = utils.Validate.Struct(creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		accessToken, refreshToken, err := service.SignIn(r.Context(), creds)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, err.Error())
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

func SignUpHandler(service AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signUpRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		err = utils.Validate.Struct(creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = service.SignUp(r.Context(), creds)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondJSON(w, http.StatusCreated, utils.MessageResponse{
			Message: "User created successfully. Please verify your email.",
		})
	}
}

func VerifyEmailHandler(service AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		code := r.URL.Query().Get("code")

		if email == "" || code == "" {
			utils.RespondError(w, http.StatusBadRequest, "Email and verification code are required")
			return
		}

		err := service.VerifyEmail(r.Context(), email, code)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Email verified successfully",
		})
	}
}

func ForgotPasswordHandler(service AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")

		if email == "" {
			utils.RespondError(w, http.StatusBadRequest, "Email is required")
			return
		}

		err := service.ForgotPassword(r.Context(), email)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error sending password reset email")
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Password reset email sent",
		})
	}
}

func ChangePasswordHandler(service AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var changePasswordReq changePasswordRequest
		err := utils.ParseJSON(r, &changePasswordReq)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		err = utils.Validate.Struct(changePasswordReq)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = service.ChangePassword(r.Context(), changePasswordReq)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error changing password")
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.MessageResponse{
			Message: "Password changed successfully",
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
