<template>
  <div v-if="!data || data.length === 0" class="detail-result-empty">
    <p>No data to display.</p>
  </div>

  <!-- Venues -->
  <div v-else-if="type === 'venues'" class="detail-result-grid">
    <div v-for="venue in data as Venue[]" :key="venue.id" class="result-card venue-card">
      <div class="result-card-header">
        <h4>{{ venue.name }}</h4>
        <span class="venue-type-badge">{{ venue.type }}</span>
      </div>
      <div class="result-card-body">
        <div class="result-field">
          <span class="result-label">Rating</span>
          <span class="result-value">{{ venue.rating.toFixed(1) }} ⭐</span>
        </div>
        <div class="result-field">
          <span class="result-label">Price</span>
          <span class="result-value">${ venue.price_tier }</span>
        </div>
        <div class="result-field">
          <span class="result-label">Ambiance</span>
          <span class="result-value">{{ venue.ambiance }}</span>
        </div>
        <div class="result-field">
          <span class="result-label">Open Late</span>
          <span class="result-value">{{ venue.is_open_late ? "Yes" : "No" }}</span>
        </div>
      </div>
    </div>
  </div>

  <!-- Accommodations -->
  <div v-else-if="type === 'accommodations'" class="detail-result-grid">
    <div v-for="acc in data as Accommodation[]" :key="acc.id" class="result-card accommodation-card">
      <div class="result-card-header">
        <h4>{{ acc.name }}</h4>
      </div>
      <div class="result-card-body">
        <div class="result-field">
          <span class="result-label">Price Tier</span>
          <span class="result-value">${ acc.price_tier }</span>
        </div>
        <div class="result-field">
          <span class="result-label">Vibe</span>
          <span class="result-value">{{ acc.vibe }}</span>
        </div>
      </div>
    </div>
  </div>

  <!-- Travel Options -->
  <div v-else-if="type === 'travel-options'" class="detail-result-grid">
    <div v-for="option in data as TravelOption[]" :key="option.id" class="result-card travel-card">
      <div class="result-card-header">
        <h4>{{ option.method }}</h4>
        <span class="travel-method-badge">{{ option.method }}</span>
      </div>
      <div class="result-card-body">
        <div class="result-field">
          <span class="result-label">From</span>
          <span class="result-value">{{ option.origin_city }}</span>
        </div>
        <div class="result-field">
          <span class="result-label">Duration</span>
          <span class="result-value">{{ option.min_minutes }}-{{ option.max_minutes }} min</span>
        </div>
        <div class="result-field">
          <span class="result-label">Practicality</span>
          <span class="result-value">{{ option.practicality_score.toFixed(2) }}</span>
        </div>
      </div>
    </div>
  </div>

  <!-- Fallback for raw JSON -->
  <div v-else class="json-block">
    <pre>{{ JSON.stringify(data, null, 2) }}</pre>
  </div>
</template>

<script setup lang="ts">
import type { Venue, Accommodation, TravelOption } from "../types";

defineProps<{
  type: "venues" | "accommodations" | "travel-options" | "raw";
  data: Venue[] | Accommodation[] | TravelOption[] | unknown[] | null;
}>();
</script>

<style scoped>
.detail-result-empty {
  padding: 2rem;
  text-align: center;
  color: var(--color-ink-muted);
  font-style: italic;
}

.detail-result-grid {
  display: grid;
  gap: 1rem;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
}

.result-card {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  background: var(--color-surface);
  overflow: hidden;
  transition: box-shadow var(--duration-fast) ease, border-color var(--duration-fast) ease;
}

.result-card:hover {
  border-color: var(--color-accent);
  box-shadow: var(--shadow-md);
}

.result-card-header {
  padding: 1rem;
  border-bottom: 1px solid var(--color-border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.5rem;
  background: linear-gradient(135deg, var(--color-bg-subtle) 0%, rgba(255, 255, 255, 0.01) 100%);
}

.result-card-header h4 {
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
  flex: 1;
  min-width: 0;
  word-break: break-word;
}

.venue-type-badge,
.travel-method-badge {
  display: inline-block;
  padding: 0.25rem 0.6rem;
  background: var(--color-accent-soft);
  color: var(--color-accent);
  border-radius: var(--radius-pill);
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  white-space: nowrap;
}

.travel-method-badge {
  background: rgba(45, 156, 219, 0.1);
  color: var(--color-info);
}

.result-card-body {
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.result-field {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  font-size: 0.9rem;
  flex-wrap: wrap;
}

.result-label {
  color: var(--color-ink-muted);
  font-weight: 500;
  text-transform: uppercase;
  font-size: 0.75rem;
  letter-spacing: 0.04em;
  flex: 0 0 auto;
}

.result-value {
  color: var(--color-ink);
  font-weight: 600;
  text-align: right;
  flex: 1;
  min-width: fit-content;
}

.json-block {
  margin-top: 1rem;
  background: #1d2d31;
  color: #d8e6e2;
  border-radius: var(--radius-md);
  border: 1px solid #274046;
  padding: 0.86rem;
  overflow: auto;
  max-height: 360px;
}

.json-block pre {
  margin: 0;
  font-size: 0.85rem;
  font-family: "Courier New", monospace;
}

@media (max-width: 768px) {
  .detail-result-grid {
    grid-template-columns: 1fr;
  }
}
</style>
