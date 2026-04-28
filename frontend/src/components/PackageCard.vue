<script setup lang="ts">
import Card from "./ui/Card.vue"
import Button from "./ui/Button.vue"
import type { TravelPackage } from "../types/travel"
import { useWishlistStore } from "../stores/wishlist"

const props = defineProps<{ item: TravelPackage }>()
const wishlist = useWishlistStore()
</script>

<template>
  <Card class="overflow-hidden p-0">
    <img :src="item.image" :alt="item.title" class="h-48 w-full object-cover" />
    <div class="space-y-3 p-4">
      <div class="flex items-center justify-between">
        <h3 class="font-semibold">{{ item.title }}</h3>
        <span class="rounded-full bg-amber-100 px-2 py-1 text-xs">⭐ {{ item.rating.toFixed(1) }}</span>
      </div>
      <p class="text-sm text-slate-600">{{ item.city }}, {{ item.country }}</p>
      <p class="text-sm text-slate-500">{{ item.durationDays }} days · {{ item.availableSpots }} spots left</p>
      <div class="flex items-center justify-between gap-2">
        <p class="text-lg font-semibold text-sky-700">${{ item.currentPrice }}</p>
        <Button
          variant="outline"
          class="px-3"
          @click="wishlist.toggle(props.item)"
        >
          {{ wishlist.isSaved(item.id) ? "Saved" : "Save" }}
        </Button>
        <RouterLink :to="`/packages/${item.id}`" class="ml-auto">
          <Button>View details</Button>
        </RouterLink>
      </div>
    </div>
  </Card>
</template>
