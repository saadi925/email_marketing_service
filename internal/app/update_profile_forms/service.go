package updateprofileforms

import (
	"context"
	"encoding/json"

	"github.com/saadi925/email_marketing_api/internal/database"
)

type UpdateProfileFormService interface {
	CreateUpdateProfileForm(ctx context.Context, name string, fields map[string]interface{}) (*UpdateProfileForm, error)
	GetUpdateProfileFormByID(ctx context.Context, id int64) (*UpdateProfileForm, error)
	UpdateUpdateProfileForm(ctx context.Context, id int64, name string, fields map[string]interface{}) error
	DeleteUpdateProfileForm(ctx context.Context, id int64) error
}

type updateProfileFormService struct {
	db *database.Queries
}

func NewUpdateProfileFormService(db *database.Queries) UpdateProfileFormService {
	return &updateProfileFormService{
		db: db,
	}
}

func (s *updateProfileFormService) CreateUpdateProfileForm(ctx context.Context, name string, fields map[string]interface{}) (*UpdateProfileForm, error) {
	fieldsJSON, err := json.Marshal(fields)
	if err != nil {
		return nil, err
	}

	dbForm, err := s.db.CreateUpdateProfileForm(ctx, database.CreateUpdateProfileFormParams{
		Name:   name,
		Fields: fieldsJSON,
	})
	if err != nil {
		return nil, err
	}

	return dbFormToModel(dbForm), nil
}

func (s *updateProfileFormService) GetUpdateProfileFormByID(ctx context.Context, id int64) (*UpdateProfileForm, error) {
	dbForm, err := s.db.GetUpdateProfileFormByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dbFormToModel(dbForm), nil
}

func (s *updateProfileFormService) UpdateUpdateProfileForm(ctx context.Context, id int64, name string, fields map[string]interface{}) error {
	fieldsJSON, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return s.db.UpdateUpdateProfileForm(ctx, database.UpdateUpdateProfileFormParams{
		ID:     id,
		Name:   name,
		Fields: fieldsJSON,
	})
}

func (s *updateProfileFormService) DeleteUpdateProfileForm(ctx context.Context, id int64) error {
	return s.db.DeleteUpdateProfileForm(ctx, id)
}

func dbFormToModel(dbForm database.UpdateProfileForm) *UpdateProfileForm {
	var fields map[string]interface{}
	json.Unmarshal(dbForm.Fields, &fields)

	return &UpdateProfileForm{
		ID:        dbForm.ID,
		Name:      dbForm.Name,
		Fields:    fields,
		CreatedAt: dbForm.CreatedAt.Time,
		UpdatedAt: dbForm.UpdatedAt.Time,
	}
}
