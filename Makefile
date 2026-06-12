DB_URL=postgres://authguard:authguard@localhost:5432/authguard?sslmode=disable
MIGRATIONS_DIR=./migrations
GOOSE=go run github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: postgres-up postgres-down migrate-create migrate-up migrate-down migrate-status migrate-version migrate-redo run-api test tidy

postgres-up:
	docker compose up -d

postgres-down:
	docker compose down

migrate-create:
	$(GOOSE) -dir $(MIGRATIONS_DIR) create table sql

migrate-up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

migrate-version:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" version

migrate-redo:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" redo

run-api:
	go run ./cmd/auth-api

test:
	go test ./...

tidy:
	go mod tidy