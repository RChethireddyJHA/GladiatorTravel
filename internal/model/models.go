package model

type Destination struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	CountryCode   string  `json:"country_code"`
	City          string  `json:"city"`
	FoodScore     float64 `json:"food_score"`
	BarScore      float64 `json:"bar_score"`
	CombinedScore float64 `json:"combined_score"`
}

type Venue struct {
	ID          int64   `json:"id"`
	Destination int64   `json:"destination_id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Rating      float64 `json:"rating"`
	PriceTier   int     `json:"price_tier"`
	OpenLate    bool    `json:"is_open_late"`
	Ambiance    string  `json:"ambiance"`
}

type Accommodation struct {
	ID          int64  `json:"id"`
	Destination int64  `json:"destination_id"`
	Name        string `json:"name"`
	PriceTier   int    `json:"price_tier"`
	Vibe        string `json:"vibe"`
}

type TravelOption struct {
	ID                int64  `json:"id"`
	OriginCity        string `json:"origin_city"`
	Destination       int64  `json:"destination_id"`
	Method            string `json:"method"`
	MinMinutes        int    `json:"min_minutes"`
	MaxMinutes        int    `json:"max_minutes"`
	PracticalityScore int    `json:"practicality_score"`
}

type UserPreferences struct {
	UserID           int64  `json:"user_id"`
	BudgetLevel      int    `json:"budget_level"`
	PrefersLateNight bool   `json:"prefers_late_night"`
	AvoidTags        string `json:"avoid_tags"`
}

type Trip struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Destination int64  `json:"destination_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type ItineraryItem struct {
	DayNumber int    `json:"day_number"`
	Slot      string `json:"slot"`
	VenueID   int64  `json:"venue_id"`
	Notes     string `json:"notes"`
}
