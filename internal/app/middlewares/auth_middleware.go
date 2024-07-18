package middlewares

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

type AuthHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

// These returns a handler function with the current user id.
// Make sure to pass a handler that returns the handler function with the user id
func WithAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := extractUserIDFromToken(r)
		if err != nil {
			log.Println(err)
			utils.RespondJSON(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		handler(w, r, userId)
	}
}
