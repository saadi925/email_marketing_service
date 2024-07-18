package auth

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	authService := newAuthService(db)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signin", signin(authService))
		r.Post("/signup", signup(authService))
		r.Put("/verify-email", verifyEmail(authService))
		r.Get("/forgot-password", forgotPassword(authService))
		r.Post("/change-password", changePassword(authService))
	})
}
