<script setup lang="ts">
import { onMounted, ref } from "vue"
import { useRoute } from "vue-router"
import Button from "../components/ui/Button.vue"
import Card from "../components/ui/Card.vue"
import { fetchPackageById } from "../api/travel"
import type { TravelPackage } from "../types/travel"

const route = useRoute()
const item = ref<TravelPackage | null>(null)
const loading = ref(false)
const error = ref("")

onMounted(async () => {
  loading.value = true
  try {
    item.value = await fetchPackageById(String(route.params.id))
  } catch {
    error.value = "Could not load package details."
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div v-if="loading" class="h-64 animate-pulse rounded-2xl bg-slate-200" />
  <div v-else-if="error" class="rounded-xl border border-red-200 bg-red-50 p-4 text-red-700">{{ error }}</div>
  <div v-else-if="item" class="grid gap-6 lg:grid-cols-[1.3fr_0.7fr]">
    <section class="space-y-4">
      <img :src="item.image" :alt="item.title" class="h-96 w-full rounded-2xl object-cover" />
      <h1 class="text-3xl font-semibold">{{ item.title }}</h1>
      <p class="text-slate-600">{{ item.description }}</p>
      <Card>
        <h2 class="font-semibold">Itinerary highlights</h2>
        <ul class="mt-2 list-disc space-y-1 pl-5 text-sm text-slate-600">
          <li>Arrival and welcome dinner</li>
          <li>Guided city and culture experience</li>
          <li>Free exploration and optional activities</li>
        </ul>
      </Card>
    </section>

    <aside class="space-y-4">
      <Card>
        <p class="text-sm text-slate-500">Starting from</p>
        <p class="text-3xl font-semibold text-sky-700">${{ item.currentPrice }}</p>
        <p class="text-sm text-slate-500">{{ item.availableSpots }} spots left</p>
        <RouterLink :to="`/booking/${item.id}`">
          <Button class="mt-4 w-full">Book now</Button>
        </RouterLink>
      </Card>
      <Card>
        <h3 class="font-semibold">Availability calendar</h3>
        <input type="date" class="mt-3 h-11 w-full rounded-xl border border-slate-300 px-3" />
      </Card>
      <Card>
        <h3 class="font-semibold">Pricing tiers</h3>
        <p class="mt-2 text-sm text-slate-600">Standard · Premium · Luxury options available by date.</p>
      </Card>
    </aside>
  </div>
</template>
