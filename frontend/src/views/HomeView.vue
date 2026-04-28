<script setup lang="ts">
import { onMounted, reactive } from "vue"
import { useRouter } from "vue-router"
import PackageCard from "../components/PackageCard.vue"
import Button from "../components/ui/Button.vue"
import Input from "../components/ui/Input.vue"
import { usePackagesStore } from "../stores/packages"

const router = useRouter()
const store = usePackagesStore()
const search = reactive({ destination: "", city: "", startDate: "" })

onMounted(() => {
  store.load()
})

function goToSearch() {
  router.push({
    path: "/search",
    query: { destination: search.destination, city: search.city, startDate: search.startDate },
  })
}
</script>

<template>
  <section class="overflow-hidden rounded-3xl bg-slate-900 p-6 text-white md:p-10">
    <div class="grid gap-6 md:grid-cols-2 md:items-end">
      <div class="space-y-4">
        <p class="text-sm text-sky-200">Explore the world with confidence</p>
        <h1 class="text-3xl font-semibold md:text-5xl">Find your next unforgettable trip.</h1>
        <p class="max-w-lg text-slate-200">Discover curated packages with real-time prices and instant booking.</p>
      </div>
      <div class="rounded-2xl bg-white/95 p-4 text-slate-900">
        <div class="grid gap-3 md:grid-cols-2">
          <Input v-model="search.destination" placeholder="Destination" aria-label="Destination" />
          <Input v-model="search.city" placeholder="City" aria-label="City" />
          <Input v-model="search.startDate" type="date" aria-label="Start date" class="md:col-span-2" />
        </div>
        <Button class="mt-3 w-full" @click="goToSearch">Search packages</Button>
      </div>
    </div>
  </section>

  <section class="mt-8">
    <div class="mb-4 flex items-center justify-between">
      <h2 class="text-2xl font-semibold">Featured destinations</h2>
      <RouterLink to="/search" class="text-sm text-sky-700">View all</RouterLink>
    </div>
    <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
      <PackageCard v-for="item in store.items.slice(0, 6)" :key="item.id" :item="item" />
    </div>
  </section>

  <section class="mt-8 grid gap-4 md:grid-cols-3">
    <article class="rounded-2xl bg-gradient-to-r from-sky-600 to-cyan-500 p-5 text-white shadow-sm">Summer promo: up to 20% off Mediterranean packages.</article>
    <article class="rounded-2xl bg-gradient-to-r from-amber-500 to-orange-500 p-5 text-white shadow-sm">Early booking bonus for autumn city breaks.</article>
    <article class="rounded-2xl bg-gradient-to-r from-emerald-500 to-teal-500 p-5 text-white shadow-sm">Last-minute escapes with flexible cancellation.</article>
  </section>
</template>
