# GladiatorTravel Vue UI

Modern Vue 3 dashboard for all GladiatorTravel API features.

## Features covered

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

## Start

1. Install Node.js 20+ and npm.
2. From `web/`, run `npm install`.
3. Run `npm run dev`.
4. Open `http://localhost:5173`.

The dev server proxies `/api` and `/health` to `http://localhost:8081`.
