package email_logs

import (
	"time"
)

type EmailLog struct {
	ID        int       `json:"id"`
	EmailID   int       `json:"email_id"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
