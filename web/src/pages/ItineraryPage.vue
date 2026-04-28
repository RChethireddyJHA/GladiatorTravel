<template>
  <section class="page">
    <UiPageHeader
      title="Itinerary Retrieval"
      subtitle="Fetch finalized itinerary timelines by trip ID and inspect them in table or JSON view."
    />

    <UiPanel>
      <div class="form-grid">
        <UiField label="Trip ID">
          <input type="number" min="1" v-model.number="tripId" />
        </UiField>
        <UiField label="Retrieve">
          <UiButton :loading="loading" text="Get Itinerary" loading-text="Loading..." @click="load" />
        </UiField>
      </div>

      <UiStatusMessage v-if="error" tone="error">{{ error }}</UiStatusMessage>

      <div class="table-wrap" v-if="items.length">
        <table class="table">
          <thead>
            <tr>
              <th>Day</th>
              <th>Slot</th>
              <th>Venue</th>
              <th>Notes</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in items" :key="`${item.day_number}-${item.slot}-${index}`">
              <td>{{ item.day_number }}</td>
              <td>{{ item.slot }}</td>
              <td>{{ item.venue_id }}</td>
              <td>{{ item.notes }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <p v-else-if="searched">No itinerary found for trip {{ tripId }}.</p>
      <p v-else-if="!loading">Enter a Trip ID and click Get Itinerary.</p>
    </UiPanel>
  </section>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { api } from "../lib/api";
import UiButton from "../components/ui/UiButton.vue";
import UiField from "../components/ui/UiField.vue";
import UiPageHeader from "../components/ui/UiPageHeader.vue";
import UiPanel from "../components/ui/UiPanel.vue";
import UiStatusMessage from "../components/ui/UiStatusMessage.vue";
import type { ItineraryItem } from "../types";

const tripId = ref(1);
const loading = ref(false);
const error = ref("");
const items = ref<ItineraryItem[]>([]);
const searched = ref(false);

async function load() {
  loading.value = true;
  error.value = "";
  items.value = [];
  searched.value = false;
  try {
    const result = await api.getItinerary(tripId.value);
    items.value = Array.isArray(result) ? result : [];
    searched.value = true;
  } catch (err) {
    error.value = err instanceof Error ? err.message : "Failed to fetch itinerary";
  } finally {
    loading.value = false;
  }
}
</script>
