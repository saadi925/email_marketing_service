package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Verified  string    `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DBUserToUser converts a database User model to an application User model.
func DBUserToUser(dbUser database.User) User {
	var user User
	utils.DbModelToModel(dbUser, &user)
	return user
}
