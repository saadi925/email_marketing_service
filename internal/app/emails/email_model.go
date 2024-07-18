package emails

import (
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type Email struct {
	ID             uuid.UUID  `json:"id"`
	CampaignID     *uuid.UUID `json:"campaign_id"`
	RecipientEmail string     `json:"recipient_email"`
	Status         string     `json:"status"`
	SubscriptionID *int64     `json:"subscription_id"`
	SentAt         time.Time  `json:"sent_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func dbEmailToModel(dbEmail database.Email) *Email {
	var campaignId *uuid.UUID
	if dbEmail.CampaignID.Valid {
		campaignId = &dbEmail.CampaignID.UUID
	}
	email := &Email{
		ID:             dbEmail.ID,
		CampaignID:     campaignId,
		RecipientEmail: dbEmail.RecipientEmail,
	}

	if dbEmail.Status.Valid {
		email.Status = dbEmail.Status.String
	}

	if dbEmail.SentAt.Valid {
		email.SentAt = dbEmail.SentAt.Time
	}

	if dbEmail.CreatedAt.Valid {
		email.CreatedAt = dbEmail.CreatedAt.Time
	}

	if dbEmail.UpdatedAt.Valid {
		email.UpdatedAt = dbEmail.UpdatedAt.Time
	}
	if dbEmail.SubscriptionID.Valid {
		email.SubscriptionID = &dbEmail.SubscriptionID.Int64
	}

	return email
}
