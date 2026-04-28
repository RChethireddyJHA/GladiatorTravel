<script setup lang="ts">
import { onMounted, ref } from "vue"
import Card from "../components/ui/Card.vue"
import Button from "../components/ui/Button.vue"
import { fetchProfile } from "../api/travel"
import type { UserProfile } from "../types/travel"
import { useRoute } from "vue-router"
import { useWishlistStore } from "../stores/wishlist"

const tab = ref<"bookings" | "saved" | "settings">("bookings")
const profile = ref<UserProfile | null>(null)
const route = useRoute()
const wishlist = useWishlistStore()

onMounted(async () => {
  if (route.query.tab === "saved") {
    tab.value = "saved"
  }
  profile.value = await fetchProfile().catch(() => null)
})
</script>

<template>
  <section class="space-y-4">
    <h1 class="text-2xl font-semibold">User dashboard</h1>
    <div class="flex flex-wrap gap-2">
      <Button :variant="tab === 'bookings' ? 'default' : 'outline'" @click="tab = 'bookings'">My bookings</Button>
      <Button :variant="tab === 'saved' ? 'default' : 'outline'" @click="tab = 'saved'">Saved trips</Button>
      <Button :variant="tab === 'settings' ? 'default' : 'outline'" @click="tab = 'settings'">Profile settings</Button>
    </div>

    <Card v-if="tab === 'bookings'">
      <h2 class="font-semibold">Upcoming bookings</h2>
      <p class="mt-2 text-slate-600">No bookings yet. Once you book a trip, it will appear here.</p>
    </Card>

    <Card v-else-if="tab === 'saved'">
      <h2 class="font-semibold">Wishlist</h2>
      <p v-if="wishlist.items.length === 0" class="mt-2 text-slate-600">Нямаш запазени дестинации.</p>
      <div v-else class="mt-3 space-y-2">
        <div v-for="item in wishlist.items" :key="item.id" class="flex items-center justify-between rounded-xl border border-slate-200 p-3">
          <div>
            <p class="font-medium">{{ item.title }}</p>
            <p class="text-sm text-slate-600">{{ item.city }}, {{ item.country }}</p>
          </div>
          <RouterLink :to="`/packages/${item.id}`" class="text-sm text-sky-700">Open</RouterLink>
        </div>
      </div>
    </Card>

    <Card v-else>
      <h2 class="font-semibold">Profile settings</h2>
      <p class="mt-2 text-slate-600">Name: {{ profile?.fullName || "Not loaded" }}</p>
      <p class="text-slate-600">Email: {{ profile?.email || "Not loaded" }}</p>
    </Card>
  </section>
</template>
