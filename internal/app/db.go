package app

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func dbInit() (*sql.DB, error) {
	dsn := getEnv("DB_URL", "postgres://postgres:Saadsaad1@localhost:5432/email_marketing?sslmode=disable")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}

// getEnv retrieves the value of the environment variable named by the key or returns the defaultValue if the variable is not present
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
