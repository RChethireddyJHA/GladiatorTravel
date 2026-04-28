<script setup lang="ts">
import { useAuthStore } from "./stores/auth"
import { useWishlistStore } from "./stores/wishlist"

const auth = useAuthStore()
const wishlist = useWishlistStore()
</script>

<template>
  <div class="min-h-screen">
    <header class="sticky top-0 z-20 border-b border-slate-200 bg-white/90 backdrop-blur">
      <div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-3">
        <RouterLink to="/" class="text-xl font-semibold text-sky-700">Gladiator Travel</RouterLink>
        <nav class="flex items-center gap-5 text-sm">
          <RouterLink to="/search" class="hover:text-sky-600">Packages</RouterLink>
          <RouterLink :to="{ name: 'dashboard', query: { tab: 'saved' } }" class="hover:text-sky-600">Saved ({{ wishlist.count }})</RouterLink>
          <RouterLink v-if="auth.isAuthenticated" to="/dashboard" class="hover:text-sky-600">Dashboard</RouterLink>
          <button
            v-if="auth.isAuthenticated"
            class="rounded-xl bg-slate-200 px-4 py-2 hover:bg-slate-300"
            @click="auth.logout"
          >
            Logout
          </button>
          <RouterLink v-else to="/auth" class="rounded-xl bg-sky-600 px-4 py-2 text-white hover:bg-sky-700">Login</RouterLink>
        </nav>
      </div>
    </header>
    <main class="mx-auto max-w-7xl px-4 py-6">
      <RouterView />
    </main>
  </div>
</template>
