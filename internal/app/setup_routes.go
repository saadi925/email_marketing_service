package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/saadi925/email_marketing_api/internal/app/auth"
	"github.com/saadi925/email_marketing_api/internal/app/users"
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

func bootstrapRoutes(config ApiConfig) *chi.Mux {
	r := chi.NewRouter()
	r.Use(enableCors())
	// Auth Routes
	// A Local Guard Middleware will be best in production , so it will redirect , if the user is loggedIn
	auth.Routes(r, config.DB)
	users.Routes(r, config.DB)
	return r
}
