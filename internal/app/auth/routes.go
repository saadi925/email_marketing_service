package auth

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	authService := NewAuthService(db)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signin", SignInHandler(authService))
		r.Post("/signup", SignUpHandler(authService))
		r.Put("/verify-email", VerifyEmailHandler(authService))
		r.Get("/forgot-password", ForgotPasswordHandler(authService))
		r.Post("/change-password", ChangePasswordHandler(authService))
	})
}
