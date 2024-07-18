package updateprofileforms

import "time"

type UpdateProfileForm struct {
	ID        int64                  `json:"id"`
	Name      string                 `json:"name"`
	Fields    map[string]interface{} `json:"fields"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
