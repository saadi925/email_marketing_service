package userprofiles

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	r.Route("/userprofile", func(r chi.Router) {
		r.Post("/", middlewares.AuthMiddleware(createUserProfile, db))
		r.Put("/", middlewares.AuthMiddleware(updateUserProfile, db))
		r.Get("/", middlewares.AuthMiddleware(getUserProfile, db))
	})
}
