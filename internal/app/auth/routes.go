package auth

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/signin", signIn(db))
		r.Get("/signup", signUp(db))
		r.Get("/verify-email", verifyEmail(db))
		r.Get("/verify-email-token", verifyOneTimeToken(db))
		r.Get("/reset-password", resetPassword(db))
	})
}
