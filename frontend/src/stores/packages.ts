import { computed, ref } from "vue"
import { defineStore } from "pinia"
import { fetchPackages } from "../api/travel"
import type { PackageFilters, TravelPackage } from "../types/travel"

const defaultFilters: PackageFilters = {
  destination: "",
  city: "",
  minPrice: 0,
  maxPrice: 5000,
  startDate: "",
  endDate: "",
  minRating: 0,
  sortBy: "popularity",
}

export const usePackagesStore = defineStore("packages", () => {
  const items = ref<TravelPackage[]>([])
  const page = ref(1)
  const pageSize = 9
  const loading = ref(false)
  const error = ref("")
  const filters = ref<PackageFilters>({ ...defaultFilters })

  const hasEmptyResult = computed(() => !loading.value && !error.value && items.value.length === 0)
  const totalPages = computed(() => Math.max(1, Math.ceil(items.value.length / pageSize)))
  const pagedItems = computed(() => {
    const start = (page.value - 1) * pageSize
    return items.value.slice(start, start + pageSize)
  })

  async function load() {
    loading.value = true
    error.value = ""
    try {
      items.value = await fetchPackages(filters.value)
      page.value = 1
    } catch {
      error.value = "Could not load travel packages. Please try again."
      items.value = []
    } finally {
      loading.value = false
    }
  }

  function updateFilters(patch: Partial<PackageFilters>) {
    filters.value = { ...filters.value, ...patch }
  }

  function setPage(next: number) {
    if (next < 1) {
      page.value = 1
      return
    }
    if (next > totalPages.value) {
      page.value = totalPages.value
      return
    }
    page.value = next
  }

  return { items, loading, error, filters, hasEmptyResult, pagedItems, page, totalPages, load, updateFilters, setPage }
})
