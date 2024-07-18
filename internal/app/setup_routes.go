package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/saadi925/email_marketing_api/internal/app/auth"
	"github.com/saadi925/email_marketing_api/internal/app/users"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

// enableCors sets up CORS middleware
func enableCors() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "application/json"},
		AllowCredentials: true,
	})
}

func bootstrapRoutes(config apiConfig) *chi.Mux {
	r := chi.NewRouter()
	r.Use(enableCors())
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondJSON(w, http.StatusOK, "Hi Buddy")
	})
	auth.Routes(r, config.DB)

	users.Routes(r, config.DB)
	return r
}
