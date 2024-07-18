package contacts

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Subscribed     bool      `json:"subscribed"`
	Blocklisted    bool      `json:"blocklisted"`
	Email          string    `json:"email"`
	Whatsapp       string    `json:"whatsapp"`
	LandlineNumber string    `json:"landline_number"`
	LastChanged    time.Time `json:"last_changed"`
	DateAdded      time.Time `json:"date_added"`
}
