export type Destination = {
  id: number;
  name: string;
  country_code: string;
  city: string;
  food_score: number;
  bar_score: number;
  combined_score: number;
};

export type Venue = {
  id: number;
  destination_id: number;
  name: string;
  type: "restaurant" | "bar";
  rating: number;
  price_tier: number;
  is_open_late: boolean;
  ambiance: string;
};

export type Accommodation = {
  id: number;
  destination_id: number;
  name: string;
  price_tier: number;
  vibe: string;
};

export type TravelOption = {
  id: number;
  origin_city: string;
  destination_id: number;
  method: "plane" | "train" | "car" | "boat";
  min_minutes: number;
  max_minutes: number;
  practicality_score: number;
};

export type ItineraryItem = {
  day_number: number;
  slot: "lunch" | "dinner" | "bar";
  venue_id: number;
  notes: string;
};

export type ApiError = {
  error: string;
};
