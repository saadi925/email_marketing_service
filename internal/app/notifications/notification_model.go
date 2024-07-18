package notifications

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        int32     `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Message   string    `json:"message"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}
