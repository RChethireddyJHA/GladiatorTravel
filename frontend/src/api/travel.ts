import { api } from "./client"
import type { PackageFilters, TravelPackage, UserProfile } from "../types/travel"

const pseudoPrice = (score: number, food: number, bar: number) =>
  Math.round(450 + score * 120 + food * 35 + bar * 25)

const pseudoAvailability = (score: number) => Math.max(2, 18 - Math.round(score*2))

const pseudoDate = (id: number) => {
  const d = new Date()
  d.setDate(d.getDate() + (id * 4) % 26)
  return d.toISOString().slice(0, 10)
}

const mapPackage = (p: any): TravelPackage => ({
  id: String(p.id),
  title: p.title || p.name || `${p.city || "City"} Escape`,
  city: p.city || p.destination || "Unknown",
  destination: p.destination || p.city || "Unknown",
  country: p.country || p.country_code || "",
  description: p.description || "Handpicked itinerary for food, culture and unforgettable views.",
  durationDays: Number(p.duration_days || 5),
  rating: Number(p.rating || p.combined_score || 4.5),
  image: p.image || "https://images.unsplash.com/photo-1469474968028-56623f02e42e?auto=format&fit=crop&w=1200&q=80",
  availableSpots: Number(p.available_spots || 12),
  currentPrice: Number(p.current_price || 999),
  oldPrice: p.old_price ? Number(p.old_price) : undefined,
  startDate: p.start_date || new Date().toISOString().slice(0, 10),
  tags: p.tags || ["city", "food", "culture"],
})

export async function fetchPackages(filters: PackageFilters): Promise<TravelPackage[]> {
  const sort = filters.sortBy === "price" ? "food" : "combined"
  const { data } = await api.get("/v1/destinations", { params: { sort } })
  const list = Array.isArray(data) ? data : []

  let items = list.map((d: any) =>
    mapPackage({
      id: d.id,
      name: d.name || `${d.city} Discovery`,
      city: d.city,
      country_code: d.country_code,
      combined_score: Number(d.combined_score || 0),
      food_score: Number(d.food_score || 0),
      bar_score: Number(d.bar_score || 0),
      current_price: pseudoPrice(Number(d.combined_score || 0), Number(d.food_score || 0), Number(d.bar_score || 0)),
      available_spots: pseudoAvailability(Number(d.combined_score || 0)),
      start_date: pseudoDate(Number(d.id || 1)),
      tags: ["food", "bar", "city"],
    }),
  )

  if (filters.destination.trim()) {
    const needle = filters.destination.toLowerCase()
    items = items.filter((x) => `${x.destination} ${x.title} ${x.country}`.toLowerCase().includes(needle))
  }
  if (filters.city.trim()) {
    const cityNeedle = filters.city.toLowerCase()
    items = items.filter((x) => x.city.toLowerCase().includes(cityNeedle))
  }
  items = items.filter((x) => x.rating >= filters.minRating)
  items = items.filter((x) => x.currentPrice >= filters.minPrice && x.currentPrice <= filters.maxPrice)
  if (filters.startDate) {
    items = items.filter((x) => x.startDate >= filters.startDate)
  }
  if (filters.endDate) {
    items = items.filter((x) => x.startDate <= filters.endDate)
  }
  items.sort((a, b) => (filters.sortBy === "price" ? a.currentPrice - b.currentPrice : b.rating - a.rating))
  return items
}

export async function fetchPackageById(id: string): Promise<TravelPackage> {
  const { data } = await api.get("/v1/destinations")
  const match = (Array.isArray(data) ? data : []).find((d: any) => String(d.id) === id) || {}
  return mapPackage({
    ...match,
    current_price: pseudoPrice(Number(match.combined_score || 0), Number(match.food_score || 0), Number(match.bar_score || 0)),
    available_spots: pseudoAvailability(Number(match.combined_score || 0)),
    start_date: pseudoDate(Number(match.id || 1)),
  })
}

export async function createBooking(payload: Record<string, unknown>) {
  const { data } = await api.post("/v1/trips", payload)
  return data
}

export async function fetchProfile(): Promise<UserProfile> {
  const { data } = await api.get("/v1/me")
  return {
    id: String(data.id || "0"),
    fullName: data.display_name || data.fullName || data.name || "Traveler",
    email: data.email || "traveler@example.com",
  }
}
