package app

import (
	"log"

	"github.com/saadi925/email_marketing_api/internal/database"
)

func App() {
	db, err := dbInit()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	apiConf := ApiConfig{
		DB: database.New(db),
	}
	defer db.Close()
	bootstrapRoutes(apiConf)
}
