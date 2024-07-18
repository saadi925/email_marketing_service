package notifications

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type NotificationService interface {
	createNotification(ctx context.Context, userID uuid.UUID, message string) (Notification, error)
	updateNotificationReadStatus(ctx context.Context, notificationID int, read bool) error
	getNotificationsByUserID(ctx context.Context, userID uuid.UUID) ([]Notification, error)
}

type notificationService struct {
	db *database.Queries
}

func NewNotificationService(db *database.Queries) NotificationService {
	return &notificationService{db: db}
}

func (s *notificationService) createNotification(ctx context.Context, userID uuid.UUID, message string) (Notification, error) {
	createdAt := time.Now()
	notification := Notification{
		UserID:    userID,
		Message:   message,
		Read:      false,
		CreatedAt: createdAt,
	}

	dbNotification := database.CreateNotificationParams{
		UserID:  userID,
		Message: message,
	}

	if _, err := s.db.CreateNotification(ctx, dbNotification); err != nil {
		return Notification{}, err
	}

	return notification, nil
}

func (s *notificationService) updateNotificationReadStatus(ctx context.Context, notificationID int, read bool) error {
	params := database.UpdateNotificationReadStatusParams{
		ID: int32(notificationID),
		Read: sql.NullBool{
			Valid: true,
			Bool:  read,
		},
	}

	return s.db.UpdateNotificationReadStatus(ctx, params)
}

func (s *notificationService) getNotificationsByUserID(ctx context.Context, userID uuid.UUID) ([]Notification, error) {
	dbNotifications, err := s.db.GetNotificationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var notifications []Notification
	for _, dbNotification := range dbNotifications {
		notifications = append(notifications, Notification{
			ID:      dbNotification.ID,
			UserID:  dbNotification.UserID,
			Message: dbNotification.Message,
		})
	}

	return notifications, nil
}
