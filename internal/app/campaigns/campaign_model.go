package campaigns

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
