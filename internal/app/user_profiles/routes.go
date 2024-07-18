package userprofiles

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	profileService := NewUserProfileService(db)
	r.Route("/userprofile", func(r chi.Router) {
		r.Post("/", middlewares.WithAuth(createUserProfile(profileService)))
		r.Put("/", middlewares.WithAuth(updateUserProfile(profileService)))
		r.Get("/", middlewares.WithAuth(getUserProfile(profileService)))
	})
}
