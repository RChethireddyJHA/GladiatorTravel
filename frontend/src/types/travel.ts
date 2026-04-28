export interface TravelPackage {
  id: string
  title: string
  city: string
  destination: string
  country: string
  description: string
  durationDays: number
  rating: number
  image: string
  availableSpots: number
  currentPrice: number
  oldPrice?: number
  startDate: string
  tags: string[]
}

export interface PackageFilters {
  destination: string
  city: string
  minPrice: number
  maxPrice: number
  startDate: string
  endDate: string
  minRating: number
  sortBy: "price" | "popularity"
}

export interface UserProfile {
  id: string
  fullName: string
  email: string
}
