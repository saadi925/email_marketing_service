package email_logs

import (
	"time"

	"github.com/google/uuid"
)

type EmailLog struct {
	ID        uuid.UUID `json:"id"`
	EmailID   uuid.UUID `json:"email_id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
