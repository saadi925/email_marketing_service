package app

import (
	"log"
	"net/http"

	"github.com/saadi925/email_marketing_api/internal/database"
)

func App() {
	db, err := DBInit()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	apiConf := apiConfig{
		DB: database.New(db),
	}
	defer db.Close()
	r := bootstrapRoutes(apiConf)
	port := getEnv("PORT", "8080")

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
