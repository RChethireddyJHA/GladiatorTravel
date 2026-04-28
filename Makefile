APP_NAME=gladiatortravel
DB_URL?=postgres://postgres:postgres@localhost:5432/gladiatortravel?sslmode=disable
API_PORT?=8080
UI1_PORT?=5173
UI2_PORT?=5174

.PHONY: help
help:
	@echo "Targets:"
	@echo "  run-api       Run Go API"
	@echo "  build-api     Build Go API binary"
	@echo "  run-ui1       Run UI 1 (web)"
	@echo "  run-ui2       Run UI 2 (frontend)"
	@echo "  build-ui1     Build UI 1 (web)"
	@echo "  build-ui2     Build UI 2 (frontend)"
	@echo "  install-ui1   Install UI 1 dependencies"
	@echo "  install-ui2   Install UI 2 dependencies"
	@echo "  up            Start docker services"
	@echo "  down          Stop docker services"

.PHONY: run
run:
	PORT=$(API_PORT) DATABASE_URL=$(DB_URL) go run ./cmd/api

.PHONY: run-api
run-api: run

.PHONY: build-api
build-api:
	go build -o $(APP_NAME).exe ./cmd/api

.PHONY: install-ui1
install-ui1:
	npm --prefix ./web install

.PHONY: install-ui2
install-ui2:
	npm --prefix ./frontend install

.PHONY: run-ui1
run-ui1:
	npm --prefix ./web run dev -- --port $(UI1_PORT)

.PHONY: run-ui2
run-ui2:
	npm --prefix ./frontend run dev -- --port $(UI2_PORT)

.PHONY: build-ui1
build-ui1:
	npm --prefix ./web run build

.PHONY: build-ui2
build-ui2:
	npm --prefix ./frontend run build

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

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
