package email_lists

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type EmailList struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     uuid.UUID `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EmailListService interface {
	CreateEmailList(ctx context.Context, name, description string, ownerID uuid.UUID) (*EmailList, error)
	GetEmailListByID(ctx context.Context, id uuid.UUID) (*EmailList, error)
	GetEmailListsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*EmailList, error)
	UpdateEmailList(ctx context.Context, id uuid.UUID, name, description string) (*EmailList, error)
	DeleteEmailList(ctx context.Context, id uuid.UUID) error
}

type emailListService struct {
	db *database.Queries
}

func NewEmailListService(db *database.Queries) EmailListService {
	return &emailListService{
		db: db,
	}
}

func (s *emailListService) CreateEmailList(ctx context.Context, name, description string, ownerID uuid.UUID) (*EmailList, error) {
	dbEmailList, err := s.db.CreateEmailList(ctx, database.CreateEmailListParams{
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		OwnerID: ownerID,
	})
	if err != nil {
		return nil, err
	}

	return dbEmailListToModel(dbEmailList), nil
}

func (s *emailListService) GetEmailListByID(ctx context.Context, id uuid.UUID) (*EmailList, error) {
	dbEmailList, err := s.db.GetEmailListByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dbEmailListToModel(dbEmailList), nil
}

func (s *emailListService) GetEmailListsByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*EmailList, error) {
	dbEmailLists, err := s.db.GetEmailListsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	var emailLists []*EmailList
	for _, dbEmailList := range dbEmailLists {
		emailLists = append(emailLists, dbEmailListToModel(dbEmailList))
	}

	return emailLists, nil
}

func (s *emailListService) UpdateEmailList(ctx context.Context, id uuid.UUID, name, description string) (*EmailList, error) {
	dbEmailList, err := s.db.UpdateEmailList(ctx, database.UpdateEmailListParams{
		ID:   id,
		Name: name,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	})
	if err != nil {
		return nil, err
	}

	return dbEmailListToModel(dbEmailList), nil
}

func (s *emailListService) DeleteEmailList(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteEmailList(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbEmailListToModel(dbEmailList database.EmailList) *EmailList {
	return &EmailList{
		ID:          dbEmailList.ID,
		Name:        dbEmailList.Name,
		Description: dbEmailList.Description.String,
		OwnerID:     dbEmailList.OwnerID,
		CreatedAt:   dbEmailList.CreatedAt.Time,
		UpdatedAt:   dbEmailList.UpdatedAt.Time,
	}
}
