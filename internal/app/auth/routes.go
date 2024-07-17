package auth

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signin", signIn(db))
		r.Post("/signup", signUp(db))
		r.Put("/verify-email", verifyEmail(db))
		r.Get("/forgot-password", forgotPassword(db))
		r.Post("/change-password", changePassword(db))
	})
}
