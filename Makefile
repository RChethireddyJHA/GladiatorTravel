APP_NAME=gladiatortravel
DB_URL?=postgres://postgres:postgres@localhost:5432/gladiatortravel?sslmode=disable

.PHONY: run
run:
	PORT=8080 DATABASE_URL=$(DB_URL) go run ./cmd/api

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

.PHONY: db-seed
db-seed:
	psql "$(DB_URL)" -f seed/001_base.sql
	psql "$(DB_URL)" -f seed/002_scores_and_core_data.sql
	psql "$(DB_URL)" -f seed/003_itinerary_samples.sql

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...
