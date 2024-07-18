package email_templates

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type EmailTemplate struct {
	ID               uuid.UUID      `json:"id"`
	UserID           uuid.UUID      `json:"user_id"`
	TemplateName     string         `json:"template_name"`
	SubjectLine      string         `json:"subject_line"`
	PreviewText      sql.NullString `json:"preview_text"`
	FromEmail        string         `json:"from_email"`
	FromName         string         `json:"from_name"`
	ReplyToEmail     sql.NullString `json:"reply_to_email"`
	CustomizeToField sql.NullBool   `json:"customize_to_field"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
