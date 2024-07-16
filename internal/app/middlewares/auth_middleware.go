package middlewares

import (
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

func AuthMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := extractUserIDFromToken(r)
		if err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		handler(w, r, userId)
	}
}
