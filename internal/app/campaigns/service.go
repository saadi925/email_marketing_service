package campaigns

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type CampaignService interface {
	CreateCampaign(ctx context.Context, userID uuid.UUID, name, subject, body string, scheduledAt sql.NullTime) (*Campaign, error)
	GetCampaignByID(ctx context.Context, campaignID int64) (*Campaign, error)
	GetCampaignsByUserID(ctx context.Context, userID uuid.UUID) ([]*Campaign, error)
	UpdateCampaign(ctx context.Context, campaignID int64, name, subject, body string, scheduledAt sql.NullTime) error
	DeleteCampaign(ctx context.Context, campaignID int64) error
}

type campaignService struct {
	db *database.Queries
}

func NewCampaignService(db *database.Queries) CampaignService {
	return &campaignService{
		db: db,
	}
}

func (s *campaignService) CreateCampaign(ctx context.Context, userID uuid.UUID, name, subject, body string, scheduledAt sql.NullTime) (*Campaign, error) {
	dbCampaign, err := s.db.CreateCampaign(ctx, database.CreateCampaignParams{
		UserID:      userID,
		Name:        name,
		Subject:     subject,
		Body:        body,
		ScheduledAt: scheduledAt,
	})
	if err != nil {
		return nil, err
	}

	return dbCampaignToModel(dbCampaign), nil
}

func (s *campaignService) GetCampaignByID(ctx context.Context, campaignID int64) (*Campaign, error) {
	dbCampaign, err := s.db.GetCampaignByID(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	return dbCampaignToModel(dbCampaign), nil
}

func (s *campaignService) GetCampaignsByUserID(ctx context.Context, userID uuid.UUID) ([]*Campaign, error) {
	dbCampaigns, err := s.db.GetCampaignsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	campaigns := make([]*Campaign, len(dbCampaigns))
	for i, dbCampaign := range dbCampaigns {
		campaigns[i] = dbCampaignToModel(dbCampaign)
	}

	return campaigns, nil
}

func (s *campaignService) UpdateCampaign(ctx context.Context, campaignID int64, name, subject, body string, scheduledAt sql.NullTime) error {
	return s.db.UpdateCampaign(ctx, database.UpdateCampaignParams{
		ID:          campaignID,
		Name:        name,
		Subject:     subject,
		Body:        body,
		ScheduledAt: scheduledAt,
	})
}

func (s *campaignService) DeleteCampaign(ctx context.Context, campaignID int64) error {
	return s.db.DeleteCampaign(ctx, campaignID)
}

func dbCampaignToModel(dbCampaign database.Campaign) *Campaign {
	return &Campaign{
		ID:          dbCampaign.ID,
		UserID:      dbCampaign.UserID,
		Name:        dbCampaign.Name,
		Subject:     dbCampaign.Subject,
		Body:        dbCampaign.Body,
		ScheduledAt: dbCampaign.ScheduledAt.Time,
		CreatedAt:   dbCampaign.CreatedAt.Time,
		UpdatedAt:   dbCampaign.UpdatedAt.Time,
	}
}
