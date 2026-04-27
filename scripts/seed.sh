#!/bin/sh
set -eu

DB_URL="${DATABASE_URL:-postgres://postgres:postgres@db:5432/gladiatortravel?sslmode=disable}"

echo "Resetting seeded tables..."
psql "$DB_URL" -v ON_ERROR_STOP=1 -c "
TRUNCATE TABLE
  itineraries,
  trips,
  user_venue_feedback,
  user_preferences,
  users,
  travel_options,
  accommodations,
  venues,
  destination_scores,
  destinations
RESTART IDENTITY CASCADE;"

echo "Applying seed files..."
psql "$DB_URL" -v ON_ERROR_STOP=1 -f /seed/001_base.sql
psql "$DB_URL" -v ON_ERROR_STOP=1 -f /seed/002_scores_and_core_data.sql
psql "$DB_URL" -v ON_ERROR_STOP=1 -f /seed/003_itinerary_samples.sql

echo "Seeding complete."
