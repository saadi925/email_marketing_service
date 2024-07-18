package campaign_options

import (
	"time"
)

type CampaignOption struct {
	ID                    int64     `json:"id"`
	CampaignID            int64     `json:"campaign_id"`
	EnableGoogleAnalytics bool      `json:"enable_google_analytics"`
	UpdateProfileFormID   *int64    `json:"update_profile_form_id,omitempty"`
	Tags                  []string  `json:"tags,omitempty"`
	Attachments           []string  `json:"attachments,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
