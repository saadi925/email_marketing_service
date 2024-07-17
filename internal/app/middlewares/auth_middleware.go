package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthHandler func(http.ResponseWriter, *http.Request, uuid.UUID, *database.Queries)

func AuthMiddleware(handler AuthHandler, queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := extractUserIDFromToken(r)
		if err != nil {
			log.Println(err)
			utils.RespondJSON(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		handler(w, r, userId, queries)
	}
}
