package campaign_options

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/saadi925/email_marketing_api/internal/database"
	"github.com/sqlc-dev/pqtype"
)

type CampaignOptionService interface {
	CreateCampaignOption(ctx context.Context, campaignID int64, enableGoogleAnalytics bool, updateProfileFormID *int64, tags []string, attachments []string) (*CampaignOption, error)
	GetCampaignOptionsByCampaignID(ctx context.Context, campaignID int64) ([]*CampaignOption, error)
	UpdateCampaignOption(ctx context.Context, id int64, enableGoogleAnalytics bool, updateProfileFormID *int64, tags []string, attachments []string) error
}

type campaignOptionService struct {
	db *database.Queries
}

func NewCampaignOptionService(db *database.Queries) CampaignOptionService {
	return &campaignOptionService{
		db: db,
	}
}

func (s *campaignOptionService) CreateCampaignOption(ctx context.Context, campaignID int64, enableGoogleAnalytics bool, updateProfileFormID *int64, tags []string, attachments []string) (*CampaignOption, error) {
	tagsRaw, err := json.Marshal(tags)
	if err != nil {
		return nil, err
	}
	attachmentsRaw, err := json.Marshal(attachments)
	if err != nil {
		return nil, err
	}

	dbCampaignOption, err := s.db.CreateCampaignOption(ctx, database.CreateCampaignOptionParams{
		CampaignID:            campaignID,
		EnableGoogleAnalytics: sql.NullBool{Bool: enableGoogleAnalytics, Valid: true},
		UpdateProfileFormID: sql.NullInt64{
			Int64: 0,
			Valid: updateProfileFormID != nil,
		},
		Tags: pqtype.NullRawMessage{
			RawMessage: tagsRaw,
			Valid:      true,
		},
		Attachments: pqtype.NullRawMessage{
			RawMessage: attachmentsRaw,
			Valid:      true,
		},
	})
	if err != nil {
		return nil, err
	}

	return dbCampaignOptionToModel(dbCampaignOption), nil
}

func (s *campaignOptionService) GetCampaignOptionsByCampaignID(ctx context.Context, campaignID int64) ([]*CampaignOption, error) {
	dbCampaignOptions, err := s.db.GetCampaignOptionsByCampaignID(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	campaignOptions := make([]*CampaignOption, len(dbCampaignOptions))
	for i, dbCampaignOption := range dbCampaignOptions {
		campaignOptions[i] = dbCampaignOptionToModel(dbCampaignOption)
	}

	return campaignOptions, nil
}

func (s *campaignOptionService) UpdateCampaignOption(ctx context.Context, id int64, enableGoogleAnalytics bool, updateProfileFormID *int64, tags []string, attachments []string) error {
	tagsRaw, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	attachmentsRaw, err := json.Marshal(attachments)
	if err != nil {
		return err
	}

	updateFormID := sql.NullInt64{}
	if updateProfileFormID != nil {
		updateFormID = sql.NullInt64{
			Int64: *updateProfileFormID,
			Valid: true,
		}
	}

	return s.db.UpdateCampaignOption(ctx, database.UpdateCampaignOptionParams{
		ID: id,
		EnableGoogleAnalytics: sql.NullBool{
			Bool:  enableGoogleAnalytics,
			Valid: true,
		},
		UpdateProfileFormID: updateFormID,
		Tags: pqtype.NullRawMessage{
			RawMessage: tagsRaw,
			Valid:      true,
		},
		Attachments: pqtype.NullRawMessage{
			RawMessage: attachmentsRaw,
			Valid:      true,
		},
	})
}

func dbCampaignOptionToModel(dbCampaignOption database.CampaignOption) *CampaignOption {
	var tags, attachments []string

	if dbCampaignOption.Tags.Valid {
		_ = json.Unmarshal(dbCampaignOption.Tags.RawMessage, &tags)
	}
	if dbCampaignOption.Attachments.Valid {
		_ = json.Unmarshal(dbCampaignOption.Attachments.RawMessage, &attachments)
	}

	return &CampaignOption{
		ID:                    dbCampaignOption.ID,
		CampaignID:            dbCampaignOption.CampaignID,
		EnableGoogleAnalytics: dbCampaignOption.EnableGoogleAnalytics.Bool,
		UpdateProfileFormID:   nullableInt64ToPointer(dbCampaignOption.UpdateProfileFormID),
		Tags:                  tags,
		Attachments:           attachments,
		CreatedAt:             dbCampaignOption.CreatedAt.Time,
		UpdatedAt:             dbCampaignOption.UpdatedAt.Time,
	}
}

func nullableInt64ToPointer(ni sql.NullInt64) *int64 {
	if ni.Valid {
		return &ni.Int64
	}
	return nil
}
