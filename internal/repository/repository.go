package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"gladiatortravel/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(databaseURL string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Close() error { return r.db.Close() }

func (r *PostgresRepository) Health(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

func (r *PostgresRepository) ListDestinations(ctx context.Context, sortBy string, minScore float64) ([]model.Destination, error) {
	orderColumn := "combined_score"
	if sortBy == "food" {
		orderColumn = "food_score"
	}
	if sortBy == "bar" {
		orderColumn = "bar_score"
	}

	query := fmt.Sprintf(`
		SELECT d.id, d.name, d.country_code, d.city, s.food_score, s.bar_score, s.combined_score
		FROM destinations d
		JOIN destination_scores s ON s.destination_id = d.id
		WHERE s.combined_score >= $1
		ORDER BY %s DESC`, orderColumn)

	rows, err := r.db.QueryContext(ctx, query, minScore)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.Destination
	for rows.Next() {
		var d model.Destination
		if err := rows.Scan(&d.ID, &d.Name, &d.CountryCode, &d.City, &d.FoodScore, &d.BarScore, &d.CombinedScore); err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) ListVenues(ctx context.Context, destinationID int64, venueType string, openLate *bool, priceTier *int) ([]model.Venue, error) {
	query := `
		SELECT id, destination_id, name, type, rating, price_tier, is_open_late, ambiance
		FROM venues
		WHERE destination_id = $1
		  AND ($2 = '' OR type = $2)
		  AND ($3::boolean IS NULL OR is_open_late = $3)
		  AND ($4::integer IS NULL OR price_tier = $4)
		ORDER BY rating DESC`
	rows, err := r.db.QueryContext(ctx, query, destinationID, venueType, openLate, priceTier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.Venue
	for rows.Next() {
		var v model.Venue
		if err := rows.Scan(&v.ID, &v.Destination, &v.Name, &v.Type, &v.Rating, &v.PriceTier, &v.OpenLate, &v.Ambiance); err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) ListAccommodations(ctx context.Context, destinationID int64) ([]model.Accommodation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, destination_id, name, price_tier, vibe
		FROM accommodations
		WHERE destination_id = $1`, destinationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.Accommodation
	for rows.Next() {
		var a model.Accommodation
		if err := rows.Scan(&a.ID, &a.Destination, &a.Name, &a.PriceTier, &a.Vibe); err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) ListTravelOptions(ctx context.Context, destinationID int64, originCity string) ([]model.TravelOption, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, origin_city, destination_id, method, min_minutes, max_minutes, practicality_score
		FROM travel_options
		WHERE destination_id = $1 AND ($2 = '' OR origin_city = $2)
		ORDER BY practicality_score DESC`, destinationID, originCity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []model.TravelOption
	for rows.Next() {
		var t model.TravelOption
		if err := rows.Scan(&t.ID, &t.OriginCity, &t.Destination, &t.Method, &t.MinMinutes, &t.MaxMinutes, &t.PracticalityScore); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) UpsertPreferences(ctx context.Context, p model.UserPreferences) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO user_preferences (user_id, budget_level, prefers_late_night, avoid_tags)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id)
		DO UPDATE SET budget_level = EXCLUDED.budget_level,
					  prefers_late_night = EXCLUDED.prefers_late_night,
					  avoid_tags = EXCLUDED.avoid_tags`, p.UserID, p.BudgetLevel, p.PrefersLateNight, p.AvoidTags)
	return err
}

func (r *PostgresRepository) AddVenueFeedback(ctx context.Context, userID, venueID int64, feedback string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO user_venue_feedback (user_id, venue_id, feedback)
		VALUES ($1, $2, $3)`, userID, venueID, feedback)
	return err
}

func (r *PostgresRepository) CreateTrip(ctx context.Context, t model.Trip) (int64, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO trips (user_id, destination_id, start_date, end_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, t.UserID, t.Destination, t.StartDate, t.EndDate).Scan(&id)
	return id, err
}

func (r *PostgresRepository) InsertItineraryItems(ctx context.Context, tripID int64, items []model.ItineraryItem) error {
	for _, item := range items {
		if _, err := r.db.ExecContext(ctx, `
			INSERT INTO itineraries (trip_id, day_number, slot, venue_id, notes)
			VALUES ($1, $2, $3, $4, $5)`, tripID, item.DayNumber, item.Slot, item.VenueID, item.Notes); err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) GetItinerary(ctx context.Context, tripID int64) ([]model.ItineraryItem, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT day_number, slot, venue_id, notes
		FROM itineraries
		WHERE trip_id = $1
		ORDER BY day_number, slot`, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []model.ItineraryItem{}
	for rows.Next() {
		var it model.ItineraryItem
		if err := rows.Scan(&it.DayNumber, &it.Slot, &it.VenueID, &it.Notes); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) TopVenuesForDestination(ctx context.Context, destinationID int64) ([]model.Venue, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, destination_id, name, type, rating, price_tier, is_open_late, ambiance
		FROM venues
		WHERE destination_id = $1
		ORDER BY rating DESC
		LIMIT 20`, destinationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []model.Venue
	for rows.Next() {
		var v model.Venue
		if err := rows.Scan(&v.ID, &v.Destination, &v.Name, &v.Type, &v.Rating, &v.PriceTier, &v.OpenLate, &v.Ambiance); err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) CreateUser(ctx context.Context, email, displayName string) (model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO users (email, display_name)
		VALUES ($1, $2)
		RETURNING id, email, display_name`, email, displayName).Scan(&u.ID, &u.Email, &u.DisplayName)
	return u, err
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, display_name
		FROM users
		WHERE email = $1`, email).Scan(&u.ID, &u.Email, &u.DisplayName)
	return u, err
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, display_name
		FROM users
		WHERE id = $1`, id).Scan(&u.ID, &u.Email, &u.DisplayName)
	return u, err
}
