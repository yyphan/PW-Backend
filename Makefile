-include .env

DB_DSN := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATION_DIR := migrations

new:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(DB_DSN) up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(DB_DSN) down

migrate-force:
	migrate -path $(MIGRATION_DIR) -database $(DB_DSN) force $(v)
