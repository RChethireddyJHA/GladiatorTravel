CREATE TABLE IF NOT EXISTS destinations (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  country_code TEXT NOT NULL,
  city TEXT NOT NULL,
  lat DOUBLE PRECISION NOT NULL,
  lng DOUBLE PRECISION NOT NULL
);

CREATE TABLE IF NOT EXISTS destination_scores (
  destination_id BIGINT PRIMARY KEY REFERENCES destinations(id) ON DELETE CASCADE,
  food_score NUMERIC(4,2) NOT NULL,
  bar_score NUMERIC(4,2) NOT NULL,
  combined_score NUMERIC(4,2) NOT NULL,
  venue_count INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS venues (
  id BIGSERIAL PRIMARY KEY,
  destination_id BIGINT NOT NULL REFERENCES destinations(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  type TEXT NOT NULL CHECK (type IN ('restaurant', 'bar')),
  rating NUMERIC(3,2) NOT NULL,
  price_tier INT NOT NULL CHECK (price_tier BETWEEN 1 AND 4),
  is_open_late BOOLEAN NOT NULL DEFAULT FALSE,
  lat DOUBLE PRECISION NOT NULL,
  lng DOUBLE PRECISION NOT NULL,
  ambiance TEXT NOT NULL DEFAULT 'balanced'
);

CREATE TABLE IF NOT EXISTS accommodations (
  id BIGSERIAL PRIMARY KEY,
  destination_id BIGINT NOT NULL REFERENCES destinations(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  price_tier INT NOT NULL CHECK (price_tier BETWEEN 1 AND 4),
  lat DOUBLE PRECISION NOT NULL,
  lng DOUBLE PRECISION NOT NULL,
  vibe TEXT NOT NULL DEFAULT 'balanced'
);

CREATE TABLE IF NOT EXISTS travel_options (
  id BIGSERIAL PRIMARY KEY,
  origin_city TEXT NOT NULL,
  destination_id BIGINT NOT NULL REFERENCES destinations(id) ON DELETE CASCADE,
  method TEXT NOT NULL CHECK (method IN ('plane', 'train', 'car', 'boat')),
  min_minutes INT NOT NULL,
  max_minutes INT NOT NULL,
  practicality_score INT NOT NULL CHECK (practicality_score BETWEEN 1 AND 100)
);

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  display_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_preferences (
  user_id BIGINT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
  budget_level INT NOT NULL CHECK (budget_level BETWEEN 1 AND 4),
  prefers_late_night BOOLEAN NOT NULL DEFAULT FALSE,
  avoid_tags TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS user_venue_feedback (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  venue_id BIGINT NOT NULL REFERENCES venues(id) ON DELETE CASCADE,
  feedback TEXT NOT NULL CHECK (feedback IN ('up', 'down')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS trips (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  destination_id BIGINT NOT NULL REFERENCES destinations(id) ON DELETE CASCADE,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS itineraries (
  id BIGSERIAL PRIMARY KEY,
  trip_id BIGINT NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
  day_number INT NOT NULL,
  slot TEXT NOT NULL CHECK (slot IN ('lunch', 'dinner', 'bar')),
  venue_id BIGINT NOT NULL REFERENCES venues(id) ON DELETE CASCADE,
  notes TEXT NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_destination_scores_combined ON destination_scores (combined_score DESC);
CREATE INDEX IF NOT EXISTS idx_venues_dest_type_rating ON venues (destination_id, type, rating DESC);
CREATE INDEX IF NOT EXISTS idx_accommodations_dest ON accommodations (destination_id);
CREATE INDEX IF NOT EXISTS idx_travel_dest_origin ON travel_options (destination_id, origin_city);
CREATE INDEX IF NOT EXISTS idx_itineraries_trip_day_slot ON itineraries (trip_id, day_number, slot);
