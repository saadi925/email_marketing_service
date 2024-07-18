package subscriptions

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        int       `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
