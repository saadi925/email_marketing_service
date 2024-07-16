package campaign_options

import (
	"time"
)

type CampaignOption struct {
	ID                    int       `json:"id"`
	CampaignID            int       `json:"campaign_id"`
	EnableGoogleAnalytics bool      `json:"enable_google_analytics"`
	UpdateProfileFormID   *int      `json:"update_profile_form_id,omitempty"`
	Tags                  []string  `json:"tags,omitempty"`
	Attachments           []string  `json:"attachments,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
