<template>
  <div class="app-shell">
    <aside class="side-nav panel-elevated">
      <div class="brand-plate" role="banner" aria-label="Gladiator Travel brand">
        <div class="brand-plate-inner">
          <img class="brand-plate-logo" :src="brandLogo" alt="Gladiator helmet logo" />
          <h1 class="brand-plate-title">GLADIATOR</h1>
          <p class="brand-plate-subtitle">TRAVEL</p>
          <p class="brand-plate-tagline">Food • Bars • Victory</p>
        </div>
      </div>

      <nav class="nav-list" aria-label="Primary">
        <RouterLink to="/">Destinations</RouterLink>
        <RouterLink to="/preferences">Preferences</RouterLink>
        <RouterLink to="/feedback">Venue Feedback</RouterLink>
        <RouterLink to="/trips">Trips</RouterLink>
        <RouterLink to="/itinerary">Itinerary</RouterLink>
      </nav>
    </aside>

    <section class="main-column">
      <header class="top-bar panel-elevated">
        <div class="top-actions" aria-label="Global workspace controls">
          <ThemeDock
            :model-value="currentTheme"
            :options="themes"
            aria-label="Global theme selector"
            @update:model-value="onThemeSelect"
          />
          <p class="status-pill"><span class="status-dot"></span> Live planning workspace</p>
        </div>
      </header>

      <main>
        <RouterView />
      </main>
    </section>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import brandLogo from "./assets/gladiator-logo.png";
import ThemeDock from "./components/ui/ThemeDock.vue";

type ThemeId = "sunset" | "ocean" | "aurora";

const themes: Array<{ id: ThemeId; label: string; icon: string }> = [
  { id: "sunset", label: "Sunset", icon: "☀️" },
  { id: "ocean", label: "Ocean", icon: "🌊" },
  { id: "aurora", label: "Aurora", icon: "✨" }
];

const currentTheme = ref<ThemeId>("sunset");

function setTheme(theme: ThemeId) {
  currentTheme.value = theme;
  document.documentElement.setAttribute("data-theme", theme);
  localStorage.setItem("gt-theme", theme);
}

function onThemeSelect(themeId: string) {
  if (!themes.some((theme) => theme.id === themeId)) {
    return;
  }
  setTheme(themeId as ThemeId);
}

onMounted(() => {
  const saved = localStorage.getItem("gt-theme") as ThemeId | null;
  if (saved && themes.some((theme) => theme.id === saved)) {
    setTheme(saved);
    return;
  }
  setTheme("sunset");
});
</script>
