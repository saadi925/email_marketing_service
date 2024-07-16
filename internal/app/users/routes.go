package users

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", middlewares.AuthMiddleware(getUser))
	})

}
