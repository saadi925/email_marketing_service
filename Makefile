include .env
export $(shell sed 's/=.*//' .env)

build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) ./cmd/server

seed:
	@echo "Seeding..."
	@go run cmd/seed/main.go

db-status:
	@echo "Checking database status..."
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose -dir=$(GOOSE_MIGRATION_DIR) status

up:
	@echo "Running migrations..."
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose -dir=$(GOOSE_MIGRATION_DIR) up

down:
	@echo "Rolling back migrations..."
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset:
	@echo "Resetting migrations..."
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DB_URL) goose -dir=$(GOOSE_MIGRATION_DIR) reset
