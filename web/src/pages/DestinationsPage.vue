<template>
  <section class="page">
    <UiPageHeader
      class="planner-hero"
      kicker="Travel Planner"
      title="Discover Cities Where Food and Nightlife Truly Matter"
    >
      <p class="hero-highlight-strip">Tonight-worthy cities for food-first escapes and nightlife vibes 🍸</p>

      <div class="quick-link-row">
        <RouterLink class="link-pill" to="/preferences">Tune preferences</RouterLink>
        <RouterLink class="link-pill" to="/trips">Create trip</RouterLink>
        <RouterLink class="link-pill" to="/itinerary">View itinerary</RouterLink>
      </div>
    </UiPageHeader>

    <UiPanel v-if="showCuratedSkeleton" class="hero-accent" aria-live="polite" aria-busy="true">
      <div class="kpi-grid kpi-skeleton-grid">
        <div v-for="n in 3" :key="n" class="kpi-skeleton-card">
          <div class="kpi-skeleton-label"></div>
          <div class="kpi-skeleton-value"></div>
        </div>
      </div>
    </UiPanel>

    <UiPanel v-else-if="showCuratedStats" class="hero-accent">
      <div class="kpi-grid">
        <UiKpiCard label="Loaded destinations" :value="destinations.length" />
        <UiKpiCard label="Top combined score" :value="topCombinedScore" />
        <UiKpiCard label="Current ranking" :value="sort" />
      </div>
    </UiPanel>

    <UiPanel v-else-if="showCuratedEmpty" class="hero-accent curated-empty-panel">
      <p class="curated-empty-text">No curated city match this round. Try widening your vibe and ask again.</p>
    </UiPanel>

    <UiPanel>
      <div class="form-grid">
        <UiField label="Sort by" for-id="sort">
          <select id="sort" v-model="sort">
            <option value="combined">combined</option>
            <option value="food">food</option>
            <option value="bar">bar</option>
          </select>
        </UiField>
        <UiField label="Minimum score" for-id="minScore">
          <input id="minScore" type="number" step="0.1" min="0" max="10" v-model.number="minScore" />
        </UiField>
      </div>
      <div class="actions-row hero-cta-row">
        <UiButton
          class="hero-cta-primary"
          variant="primary"
          :loading="loadingDestinations"
          text="Load Curated Picks ✨"
          loading-text="Loading..."
          @click="loadCuratedPicks"
        />
        <UiButton class="hero-cta-secondary" variant="secondary" text="All Destinations" @click="loadDestinations" />
      </div>
      <p class="hero-cta-microcopy">AI-picked cities you'll actually love.</p>
      <UiStatusMessage v-if="error" tone="error">{{ error }}</UiStatusMessage>
    </UiPanel>

    <UiPanel>
      <h3>Destinations</h3>
      <div class="destination-card-grid" v-if="destinations.length">
        <article
          v-for="d in destinations"
          :key="d.id"
          class="destination-card"
          :class="{ 'is-selected': selectedDestinationId === d.id }"
        >
          <div class="destination-media" :style="{ background: destinationBackdrop(d) }">
            <span class="destination-country">{{ d.country_code }}</span>
            <div class="destination-score-stack">
              <span class="score-chip score-chip-combined">C {{ formatScore(d.combined_score) }}</span>
              <span class="score-chip">F {{ formatScore(d.food_score) }}</span>
              <span class="score-chip">B {{ formatScore(d.bar_score) }}</span>
            </div>
          </div>

          <div class="destination-body">
            <h4>{{ d.name }}</h4>
            <p>{{ d.city }}, {{ d.country_code }}</p>

            <div class="actions-row">
              <UiButton
                :variant="selectedDestinationId === d.id ? 'secondary' : 'primary'"
                :text="selectedDestinationId === d.id ? 'Selected' : 'Inspect Destination'"
                @click="selectDestination(d.id)"
              />
            </div>
          </div>
        </article>
      </div>
      <p v-else>The cities are shy. Hit "Curated Picks" to wake them up.</p>
    </UiPanel>

    <UiPanel v-if="selectedDestinationId">
      <h3>Destination {{ selectedDestinationId }} Details</h3>
      <div class="form-grid">
        <UiField label="Venue type">
          <select v-model="venueType">
            <option value="">all</option>
            <option value="restaurant">restaurant</option>
            <option value="bar">bar</option>
          </select>
        </UiField>
        <UiField label="Open late">
          <select v-model="openLate">
            <option value="">all</option>
            <option value="true">true</option>
            <option value="false">false</option>
          </select>
        </UiField>
        <UiField label="Price tier">
          <input type="number" min="1" max="4" v-model.number="priceTier" />
        </UiField>
        <UiField label="Origin city (travel)">
          <input type="text" v-model="originCity" placeholder="optional" />
        </UiField>
      </div>

      <div class="actions-row">
        <UiButton text="Load Venues" @click="loadVenues" />
        <UiButton variant="secondary" text="Load Accommodations" @click="loadAccommodations" />
        <UiButton variant="ghost" text="Load Travel Options" @click="loadTravelOptions" />
      </div>

      <DetailResultDisplay :type="detailType" :data="detailData" />
    </UiPanel>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { api } from "../lib/api";
import UiButton from "../components/ui/UiButton.vue";
import UiField from "../components/ui/UiField.vue";
import UiPageHeader from "../components/ui/UiPageHeader.vue";
import UiPanel from "../components/ui/UiPanel.vue";
import UiKpiCard from "../components/ui/UiKpiCard.vue";
import UiStatusMessage from "../components/ui/UiStatusMessage.vue";
import DetailResultDisplay from "../components/DetailResultDisplay.vue";
import type { Destination } from "../types";

const sort = ref("combined");
const minScore = ref(0);
const loadingDestinations = ref(false);
const error = ref("");
const destinations = ref<Destination[]>([]);
const curatedRequested = ref(false);
const curatedLoading = ref(false);
const curatedLoadSucceeded = ref(false);

const selectedDestinationId = ref<number | null>(null);
const venueType = ref("");
const openLate = ref("");
const priceTier = ref<number | undefined>();
const originCity = ref("");
const detailType = ref<"venues" | "accommodations" | "travel-options" | "raw">("raw");
const detailData = ref<unknown[] | null>(null);

const topCombinedScore = computed(() => {
  if (!destinations.value.length) {
    return "-";
  }
  const score = Math.max(...destinations.value.map((d) => d.combined_score));
  return score.toFixed(2);
});

const showCuratedSkeleton = computed(() => curatedLoading.value);
const showCuratedStats = computed(() => curatedLoadSucceeded.value && destinations.value.length > 0);
const showCuratedEmpty = computed(
  () => curatedRequested.value && curatedLoadSucceeded.value && !curatedLoading.value && destinations.value.length === 0
);

function formatScore(value: number) {
  return value.toFixed(1);
}

function destinationBackdrop(destination: Destination) {
  const palettes = [
    ["#073b4c", "#118ab2", "#8ecae6"],
    ["#2f5233", "#4f772d", "#90a955"],
    ["#5a189a", "#7b2cbf", "#c77dff"],
    ["#8f2d56", "#d81159", "#ffbc42"],
    ["#1b4965", "#5fa8d3", "#bee9e8"]
  ];

  const palette = palettes[destination.id % palettes.length];
  return `linear-gradient(145deg, ${palette[0]} 0%, ${palette[1]} 58%, ${palette[2]} 100%)`;
}

async function loadDestinations(isCurated = false) {
  if (isCurated) {
    curatedRequested.value = true;
    curatedLoading.value = true;
    curatedLoadSucceeded.value = false;
  }

  loadingDestinations.value = true;
  error.value = "";
  try {
    destinations.value = await api.listDestinations(sort.value, minScore.value);
    if (isCurated) {
      curatedLoadSucceeded.value = true;
    }
  } catch (err) {
    error.value = err instanceof Error ? err.message : "Failed loading destinations";
    if (isCurated) {
      curatedLoadSucceeded.value = false;
    }
  } finally {
    if (isCurated) {
      curatedLoading.value = false;
    }
    loadingDestinations.value = false;
  }
}

function loadCuratedPicks() {
  sort.value = "combined";
  minScore.value = 8.5;
  void loadDestinations(true);
}

function selectDestination(id: number) {
  selectedDestinationId.value = id;
  detailData.value = null;
  detailType.value = "raw";
}

async function loadVenues() {
  if (!selectedDestinationId.value) return;
  try {
    const rows = await api.listVenues(selectedDestinationId.value, {
      type: venueType.value || undefined,
      openLate: openLate.value === "" ? undefined : openLate.value === "true",
      priceTier: priceTier.value
    });
    detailData.value = rows;
    detailType.value = "venues";
  } catch (err) {
    detailData.value = null;
    detailType.value = "raw";
  }
}

async function loadAccommodations() {
  if (!selectedDestinationId.value) return;
  try {
    const rows = await api.listAccommodations(selectedDestinationId.value);
    detailData.value = rows;
    detailType.value = "accommodations";
  } catch (err) {
    detailData.value = null;
    detailType.value = "raw";
  }
}

async function loadTravelOptions() {
  if (!selectedDestinationId.value) return;
  try {
    const rows = await api.listTravelOptions(selectedDestinationId.value, originCity.value || undefined);
    detailData.value = rows;
    detailType.value = "travel-options";
  } catch (err) {
    detailData.value = null;
    detailType.value = "raw";
  }
}
</script>
