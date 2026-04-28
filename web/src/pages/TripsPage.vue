<template>
  <section class="page">
    <UiPageHeader
      title="Trips and Auto-Generation"
      subtitle="Create a trip record, then generate an itinerary plan using the selected destination and day range."
    />

    <UiPanel>
      <h3>Trip Creation</h3>
      <div class="form-grid">
        <UiField label="User ID">
          <input type="number" min="1" v-model.number="tripForm.user_id" />
        </UiField>
        <UiField label="Destination ID">
          <input type="number" min="1" v-model.number="tripForm.destination_id" />
        </UiField>
        <UiField label="Start date">
          <input type="date" v-model="tripForm.start_date" />
        </UiField>
        <UiField label="End date">
          <input type="date" v-model="tripForm.end_date" />
        </UiField>
      </div>

      <div class="actions-row">
        <UiButton :loading="creating" text="Create Trip" loading-text="Creating..." @click="createTrip" />
      </div>

      <UiStatusMessage v-if="createdTripId" tone="success">Trip created with ID {{ createdTripId }}</UiStatusMessage>
      <UiStatusMessage v-if="createError" tone="error">{{ createError }}</UiStatusMessage>
    </UiPanel>

    <UiPanel>
      <h3>Generate Itinerary</h3>
      <div class="form-grid">
        <UiField label="Trip ID">
          <input type="number" min="1" v-model.number="generateForm.trip_id" />
        </UiField>
        <UiField label="Destination ID">
          <input type="number" min="1" v-model.number="generateForm.destination_id" />
        </UiField>
        <UiField label="Days (2-4 in backend)">
          <input type="number" min="1" max="7" v-model.number="generateForm.days" />
        </UiField>
      </div>

      <div class="actions-row">
        <UiButton
          variant="secondary"
          :loading="generating"
          text="Generate Itinerary"
          loading-text="Generating..."
          @click="generate"
        />
      </div>

      <UiStatusMessage v-if="generateSuccess" tone="success">{{ generateSuccess }}</UiStatusMessage>
      <UiStatusMessage v-if="generateError" tone="error">{{ generateError }}</UiStatusMessage>
    </UiPanel>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from "vue";
import { api } from "../lib/api";
import UiButton from "../components/ui/UiButton.vue";
import UiField from "../components/ui/UiField.vue";
import UiPageHeader from "../components/ui/UiPageHeader.vue";
import UiPanel from "../components/ui/UiPanel.vue";
import UiStatusMessage from "../components/ui/UiStatusMessage.vue";

function todayOffset(days: number): string {
  const dt = new Date();
  dt.setDate(dt.getDate() + days);
  return dt.toISOString().slice(0, 10);
}

const tripForm = reactive({
  user_id: 1,
  destination_id: 1,
  start_date: todayOffset(3),
  end_date: todayOffset(6)
});

const generateForm = reactive({
  trip_id: 1,
  destination_id: 1,
  days: 3
});

const creating = ref(false);
const generating = ref(false);
const createdTripId = ref<number | null>(null);
const createError = ref("");
const generateSuccess = ref("");
const generateError = ref("");

watch(
  () => createdTripId.value,
  (id) => {
    if (id) {
      generateForm.trip_id = id;
    }
  }
);

async function createTrip() {
  creating.value = true;
  createError.value = "";
  createdTripId.value = null;
  try {
    const result = await api.createTrip(tripForm);
    createdTripId.value = result.trip_id;
  } catch (err) {
    createError.value = err instanceof Error ? err.message : "Trip creation failed";
  } finally {
    creating.value = false;
  }
}

async function generate() {
  generating.value = true;
  generateSuccess.value = "";
  generateError.value = "";
  try {
    const result = await api.generateItinerary(generateForm.trip_id, {
      destination_id: generateForm.destination_id,
      days: generateForm.days
    });
    generateSuccess.value = `Server response: ${result.status}`;
  } catch (err) {
    generateError.value = err instanceof Error ? err.message : "Generation failed";
  } finally {
    generating.value = false;
  }
}
</script>
