package emails

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type EmailService interface {
	CreateEmail(ctx context.Context, campaignID uuid.UUID, recipientEmail string) (*Email, error)
	GetEmailByID(ctx context.Context, id uuid.UUID) (*Email, error)
	UpdateEmailStatus(ctx context.Context, id uuid.UUID, status string) error
	GetEmailsByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*Email, error)
	// Add more methods as needed for your application
}

type emailService struct {
	db *database.Queries
}

func NewEmailService(db *database.Queries) EmailService {
	return &emailService{
		db: db,
	}
}

func (s *emailService) CreateEmail(ctx context.Context, campaignID uuid.UUID, recipientEmail string) (*Email, error) {
	dbEmail, err := s.db.CreateEmail(ctx, database.CreateEmailParams{
		CampaignID:     campaignID,
		RecipientEmail: recipientEmail,
	})
	if err != nil {
		return nil, err
	}

	email := dbEmailToModel(dbEmail)
	return email, nil
}
func (s *emailService) GetEmailByID(ctx context.Context, id uuid.UUID) (*Email, error) {
	dbEmail, err := s.db.GetEmailByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbEmailToModel(dbEmail), nil
}

func (s *emailService) UpdateEmailStatus(ctx context.Context, id uuid.UUID, status string) error {
	var validStatus sql.NullString
	if status != "" {
		validStatus.String = status
		validStatus.Valid = true
	}

	return s.db.UpdateEmailStatus(ctx, database.UpdateEmailStatusParams{
		ID:     id,
		Status: validStatus,
	})
}

func (s *emailService) GetEmailsByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*Email, error) {
	dbEmails, err := s.db.GetEmailsByCampaignID(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	emails := make([]*Email, len(dbEmails))
	for i, dbEmail := range dbEmails {
		emails[i] = dbEmailToModel(dbEmail)
	}
	return emails, nil
}
