<script setup lang="ts">
import { onMounted } from "vue"
import { useRoute } from "vue-router"
import PackageCard from "../components/PackageCard.vue"
import Input from "../components/ui/Input.vue"
import Button from "../components/ui/Button.vue"
import Skeleton from "../components/ui/Skeleton.vue"
import { usePackagesStore } from "../stores/packages"
import { useWishlistStore } from "../stores/wishlist"

const route = useRoute()
const store = usePackagesStore()
const wishlist = useWishlistStore()

onMounted(async () => {
  if (route.query.destination) {
    store.updateFilters({ destination: String(route.query.destination) })
  }
  if (route.query.city) {
    store.updateFilters({ city: String(route.query.city) })
  }
  if (route.query.startDate) {
    store.updateFilters({ startDate: String(route.query.startDate) })
  }
  await store.load()
  wishlist.syncFromPackages(store.items)
})
</script>

<template>
  <div class="grid gap-6 lg:grid-cols-[280px_1fr]">
    <aside class="space-y-3 rounded-2xl border border-slate-200 bg-white p-4">
      <h2 class="font-semibold">Filters</h2>
      <Input :model-value="store.filters.destination" placeholder="Destination" @update:model-value="store.updateFilters({ destination: $event })" />
      <Input :model-value="store.filters.city" placeholder="City" @update:model-value="store.updateFilters({ city: $event })" />
      <Input type="date" :model-value="store.filters.startDate" @update:model-value="store.updateFilters({ startDate: $event })" />
      <label class="text-sm">Min rating: {{ store.filters.minRating }}</label>
      <input
        class="w-full"
        type="range"
        min="0"
        max="5"
        step="0.5"
        :value="store.filters.minRating"
        @input="store.updateFilters({ minRating: Number(($event.target as HTMLInputElement).value) })"
      />
      <label class="text-sm">Max price: ${{ store.filters.maxPrice }}</label>
      <input
        class="w-full"
        type="range"
        min="300"
        max="5000"
        step="100"
        :value="store.filters.maxPrice"
        @input="store.updateFilters({ maxPrice: Number(($event.target as HTMLInputElement).value) })"
      />
      <select
        class="h-11 w-full rounded-xl border border-slate-300 px-3"
        :value="store.filters.sortBy"
        @change="store.updateFilters({ sortBy: ($event.target as HTMLSelectElement).value as 'price' | 'popularity' })"
      >
        <option value="popularity">Sort by popularity</option>
        <option value="price">Sort by price</option>
      </select>
      <Button class="w-full" @click="async () => { await store.load(); wishlist.syncFromPackages(store.items) }">Apply filters</Button>
    </aside>

    <section>
      <p class="mb-4 text-sm text-slate-600">Dynamic results with real-time pricing and availability</p>
      <div v-if="store.loading" class="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <Skeleton v-for="n in 6" :key="n" class="h-80 w-full" />
      </div>
      <div v-else-if="store.error" class="rounded-xl border border-red-200 bg-red-50 p-4 text-red-700">
        {{ store.error }}
      </div>
      <div v-else-if="store.hasEmptyResult" class="rounded-xl border border-slate-200 bg-slate-50 p-6 text-slate-600">
        No packages found. Try a different destination or date range.
      </div>
      <div v-else class="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <PackageCard v-for="item in store.pagedItems" :key="item.id" :item="item" />
      </div>
      <div v-if="!store.loading && !store.error && !store.hasEmptyResult" class="mt-5 flex items-center justify-center gap-3">
        <Button variant="outline" :disabled="store.page <= 1" @click="store.setPage(store.page - 1)">Prev</Button>
        <span class="text-sm text-slate-600">Page {{ store.page }} / {{ store.totalPages }}</span>
        <Button variant="outline" :disabled="store.page >= store.totalPages" @click="store.setPage(store.page + 1)">Next</Button>
      </div>
    </section>
  </div>
</template>
