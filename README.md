# GladiatorTravel API (PoC)

Go proof-of-concept API for food-and-bar-driven travel discovery and itinerary generation.

## Stack
- Go + chi
- PostgreSQL
- SQL migrations + seed scripts

## Project Structure (template-aligned)
- `cmd/api`: API bootstrap
- `internal/handler`: HTTP handlers
- `internal/usecase`: business logic
- `internal/repository`: DB access
- `internal/model`: domain models
- `migrations`: schema
- `seed`: all-table seed scripts
- `api/openapi.yaml`: API contract

## Quick Start
1. Start Postgres:
   - `docker compose up -d`
2. Install migration tool (`migrate`) and run:
   - `make migrate-up`
3. Seed all tables:
   - `make db-seed`
4. Run API:
   - `make run`

## Core Endpoints
- `GET /health`
- `GET /api/v1/destinations`
- `GET /api/v1/destinations/{id}/venues`
- `GET /api/v1/destinations/{id}/accommodations`
- `GET /api/v1/destinations/{id}/travel-options`
- `PUT /api/v1/me/preferences`
- `POST /api/v1/me/venues/{venueID}/feedback`
- `POST /api/v1/trips`
- `POST /api/v1/trips/{tripID}/generate-itinerary`
- `GET /api/v1/trips/{tripID}/itinerary`
