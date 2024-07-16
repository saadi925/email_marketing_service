package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

type AuthHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

func AuthMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := ExtractUserIDFromContext(w, r)
		if err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, "unauthorzed")
			return
		}
		handler(w, r, userId)
	}
}
