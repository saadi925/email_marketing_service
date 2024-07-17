package main

import (
	"context"
	"log"

	"github.com/saadi925/email_marketing_api/internal/app"
	"github.com/saadi925/email_marketing_api/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	conn, err := app.DBInit()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	queries := database.New(conn)
	err = seedUser(queries)
	if err != nil {
		log.Fatalf("Could not seed user: %v", err)
	}

	log.Println("Database seeded successfully")
}

func seedUser(queries *database.Queries) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testuser12"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = queries.CreateUser(context.Background(), database.CreateUserParams{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: string(hashedPassword),
		Verified: true,
	})
	return err
}
