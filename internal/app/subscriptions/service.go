package subscriptions

import (
	"context"
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

var (
	ErrInvalidEmail = errors.New("invalid email address")
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, userID uuid.UUID, email string) (*Subscription, error)
	GetSubscriptionByID(ctx context.Context, id int64) (*Subscription, error)
	GetSubscriptionsByUserID(ctx context.Context, userID uuid.UUID) ([]*Subscription, error)
	UpdateSubscriptionEmail(ctx context.Context, id int64, email string) error
	DeleteSubscription(ctx context.Context, id int64) error
}

type subscriptionService struct {
	db *database.Queries
}

func NewSubscriptionService(db *database.Queries) SubscriptionService {
	return &subscriptionService{
		db: db,
	}
}

func (s *subscriptionService) CreateSubscription(ctx context.Context, userID uuid.UUID, email string) (*Subscription, error) {
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	dbSubscription, err := s.db.CreateSubscription(ctx, database.CreateSubscriptionParams{
		UserID: userID,
		Email:  email,
	})
	if err != nil {
		return nil, err
	}

	subscription := dbSubscriptionToModel(dbSubscription)
	return subscription, nil
}

func (s *subscriptionService) GetSubscriptionByID(ctx context.Context, id int64) (*Subscription, error) {
	dbSubscription, err := s.db.GetSubscriptionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	subscription := dbSubscriptionToModel(dbSubscription)
	return subscription, nil
}

func (s *subscriptionService) GetSubscriptionsByUserID(ctx context.Context, userID uuid.UUID) ([]*Subscription, error) {
	dbSubscriptions, err := s.db.GetSubscriptionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	subscriptions := make([]*Subscription, len(dbSubscriptions))
	for i, dbSubscription := range dbSubscriptions {
		subscriptions[i] = dbSubscriptionToModel(dbSubscription)
	}
	return subscriptions, nil
}

func (s *subscriptionService) UpdateSubscriptionEmail(ctx context.Context, id int64, email string) error {
	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	return s.db.UpdateSubscriptionEmail(ctx, database.UpdateSubscriptionEmailParams{
		ID:    id,
		Email: email,
	})
}

func (s *subscriptionService) DeleteSubscription(ctx context.Context, id int64) error {
	return s.db.DeleteSubscription(ctx, id)
}

func isValidEmail(email string) bool {
	// Basic regex for validating an email
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func dbSubscriptionToModel(dbSubscription database.Subscription) *Subscription {
	return &Subscription{
		ID:        dbSubscription.ID,
		UserID:    dbSubscription.UserID,
		Email:     dbSubscription.Email,
		CreatedAt: dbSubscription.CreatedAt.Time,
		UpdatedAt: dbSubscription.UpdatedAt.Time,
	}
}
