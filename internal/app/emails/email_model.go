package emails

import (
	"time"
)

type Email struct {
	ID             int       `json:"id"`
	CampaignID     int       `json:"campaign_id"`
	RecipientEmail string    `json:"recipient_email"`
	Status         string    `json:"status"`
	SentAt         time.Time `json:"sent_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
