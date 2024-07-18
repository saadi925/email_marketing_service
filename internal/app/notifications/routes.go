package notifications

import (
	"github.com/go-chi/chi"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func Routes(r chi.Router, db *database.Queries) {
	notificationService := NewNotificationService(db)
	r.Route("/notifications", func(r chi.Router) {
		r.Post("/", middlewares.WithAuth(CreateNotificationHandler(notificationService)))
		r.Put("/read-status", middlewares.WithAuth(UpdateNotificationReadStatusHandler(notificationService)))
		r.Get("/", middlewares.WithAuth(GetNotificationsHandler(notificationService)))
	})
}
