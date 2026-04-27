import type {
  Accommodation,
  ApiError,
  Destination,
  ItineraryItem,
  TravelOption,
  Venue
} from "../types";

const API_BASE = import.meta.env.VITE_API_BASE ?? "";

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers ?? {})
    },
    ...init
  });

  if (!response.ok) {
    let message = `Request failed with status ${response.status}`;
    try {
      const body = (await response.json()) as ApiError;
      if (body?.error) {
        message = body.error;
      }
    } catch {
      // Keep default fallback message if response is not JSON.
    }
    throw new Error(message);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return (await response.json()) as T;
}

export const api = {
  health: () => request<{ status: string }>("/health"),
  listDestinations: (sort?: string, minScore?: number) => {
    const params = new URLSearchParams();
    if (sort) params.set("sort", sort);
    if (typeof minScore === "number") params.set("min_score", String(minScore));
    const query = params.toString();
    return request<Destination[]>(`/api/v1/destinations${query ? `?${query}` : ""}`);
  },
  listVenues: (destinationId: number, filters: { type?: string; openLate?: boolean; priceTier?: number }) => {
    const params = new URLSearchParams();
    if (filters.type) params.set("type", filters.type);
    if (typeof filters.openLate === "boolean") params.set("open_late", String(filters.openLate));
    if (typeof filters.priceTier === "number") params.set("price_tier", String(filters.priceTier));
    const query = params.toString();
    return request<Venue[]>(`/api/v1/destinations/${destinationId}/venues${query ? `?${query}` : ""}`);
  },
  listAccommodations: (destinationId: number, mode?: string) => {
    const query = mode ? `?mode=${encodeURIComponent(mode)}` : "";
    return request<Accommodation[]>(`/api/v1/destinations/${destinationId}/accommodations${query}`);
  },
  listTravelOptions: (destinationId: number, originCity?: string) => {
    const query = originCity ? `?origin_city=${encodeURIComponent(originCity)}` : "";
    return request<TravelOption[]>(`/api/v1/destinations/${destinationId}/travel-options${query}`);
  },
  savePreferences: (payload: {
    user_id: number;
    budget_level: number;
    prefers_late_night: boolean;
    avoid_tags: string;
  }) => request<{ status: string }>("/api/v1/me/preferences", { method: "PUT", body: JSON.stringify(payload) }),
  submitFeedback: (venueId: number, payload: { user_id: number; feedback: "up" | "down" }) =>
    request<{ status: string }>(`/api/v1/me/venues/${venueId}/feedback`, {
      method: "POST",
      body: JSON.stringify(payload)
    }),
  createTrip: (payload: { user_id: number; destination_id: number; start_date: string; end_date: string }) =>
    request<{ trip_id: number }>("/api/v1/trips", { method: "POST", body: JSON.stringify(payload) }),
  generateItinerary: (tripId: number, payload: { destination_id: number; days: number }) =>
    request<{ status: string }>(`/api/v1/trips/${tripId}/generate-itinerary`, {
      method: "POST",
      body: JSON.stringify(payload)
    }),
  getItinerary: (tripId: number) => request<ItineraryItem[]>(`/api/v1/trips/${tripId}/itinerary`)
};
