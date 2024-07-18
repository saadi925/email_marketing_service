package email_templates

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type EmailTemplateService interface {
	CreateEmailTemplate(ctx context.Context, userID uuid.UUID, templateName, subjectLine, previewText, fromEmail, fromName, replyToEmail string, customizeToField bool) (*EmailTemplate, error)
	GetEmailTemplateByID(ctx context.Context, id uuid.UUID) (*EmailTemplate, error)
	GetEmailTemplatesByUserID(ctx context.Context, userID uuid.UUID) ([]*EmailTemplate, error)
	UpdateEmailTemplate(ctx context.Context, id uuid.UUID, templateName, subjectLine, previewText, fromEmail, fromName, replyToEmail string, customizeToField bool) error
	DeleteEmailTemplate(ctx context.Context, id uuid.UUID) error
}

type emailTemplateService struct {
	db *database.Queries
}

func NewEmailTemplateService(db *database.Queries) EmailTemplateService {
	return &emailTemplateService{
		db: db,
	}
}

func (s *emailTemplateService) CreateEmailTemplate(ctx context.Context, userID uuid.UUID, templateName, subjectLine, previewText, fromEmail, fromName, replyToEmail string, customizeToField bool) (*EmailTemplate, error) {
	dbEmailTemplate, err := s.db.CreateEmailTemplate(ctx, database.CreateEmailTemplateParams{
		UserID:           userID,
		TemplateName:     templateName,
		SubjectLine:      subjectLine,
		PreviewText:      sql.NullString{String: previewText, Valid: previewText != ""},
		FromEmail:        fromEmail,
		FromName:         fromName,
		ReplyToEmail:     sql.NullString{String: replyToEmail, Valid: replyToEmail != ""},
		CustomizeToField: sql.NullBool{Bool: customizeToField, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return dbEmailTemplateToModel(dbEmailTemplate), nil
}

func (s *emailTemplateService) GetEmailTemplateByID(ctx context.Context, id uuid.UUID) (*EmailTemplate, error) {
	dbEmailTemplate, err := s.db.GetEmailTemplateByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dbEmailTemplateToModel(dbEmailTemplate), nil
}

func (s *emailTemplateService) GetEmailTemplatesByUserID(ctx context.Context, userID uuid.UUID) ([]*EmailTemplate, error) {
	dbEmailTemplates, err := s.db.GetEmailTemplatesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var emailTemplates []*EmailTemplate
	for _, dbEmailTemplate := range dbEmailTemplates {
		emailTemplates = append(emailTemplates, dbEmailTemplateToModel(dbEmailTemplate))
	}

	return emailTemplates, nil
}

func (s *emailTemplateService) UpdateEmailTemplate(ctx context.Context, id uuid.UUID, templateName, subjectLine, previewText, fromEmail, fromName, replyToEmail string, customizeToField bool) error {
	err := s.db.UpdateEmailTemplate(ctx, database.UpdateEmailTemplateParams{
		ID:               id,
		TemplateName:     templateName,
		SubjectLine:      subjectLine,
		PreviewText:      sql.NullString{String: previewText, Valid: previewText != ""},
		FromEmail:        fromEmail,
		FromName:         fromName,
		ReplyToEmail:     sql.NullString{String: replyToEmail, Valid: replyToEmail != ""},
		CustomizeToField: sql.NullBool{Bool: customizeToField, Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *emailTemplateService) DeleteEmailTemplate(ctx context.Context, id uuid.UUID) error {
	err := s.db.DeleteEmailTemplate(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func dbEmailTemplateToModel(dbEmailTemplate database.EmailTemplate) *EmailTemplate {
	return &EmailTemplate{
		ID:               dbEmailTemplate.ID,
		UserID:           dbEmailTemplate.UserID,
		TemplateName:     dbEmailTemplate.TemplateName,
		SubjectLine:      dbEmailTemplate.SubjectLine,
		PreviewText:      dbEmailTemplate.PreviewText,
		FromEmail:        dbEmailTemplate.FromEmail,
		FromName:         dbEmailTemplate.FromName,
		ReplyToEmail:     dbEmailTemplate.ReplyToEmail,
		CustomizeToField: dbEmailTemplate.CustomizeToField,
		CreatedAt:        dbEmailTemplate.CreatedAt.Time,
		UpdatedAt:        dbEmailTemplate.UpdatedAt.Time,
	}
}
