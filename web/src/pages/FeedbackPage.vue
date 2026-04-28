<template>
  <section class="page">
    <UiPageHeader
      title="Venue Feedback"
      subtitle="Capture user sentiment against venues with a simple controlled signal."
    />

    <UiPanel>
      <div class="form-grid">
        <UiField label="Venue ID">
          <input type="number" min="1" v-model.number="venueId" />
        </UiField>
        <UiField label="User ID">
          <input type="number" min="1" v-model.number="userId" />
        </UiField>
        <UiField label="Feedback">
          <select v-model="feedback">
            <option value="up">up</option>
            <option value="down">down</option>
          </select>
        </UiField>
      </div>

      <div class="actions-row">
        <UiButton :loading="loading" text="Submit Feedback" loading-text="Submitting..." @click="submit" />
      </div>

      <UiStatusMessage v-if="success" tone="success">{{ success }}</UiStatusMessage>
      <UiStatusMessage v-if="error" tone="error">{{ error }}</UiStatusMessage>
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

const venueId = ref(2);
const userId = ref(1);
const feedback = ref<"up" | "down">("up");
const loading = ref(false);
const success = ref("");
const error = ref("");

async function submit() {
  loading.value = true;
  success.value = "";
  error.value = "";
  try {
    const result = await api.submitFeedback(venueId.value, {
      user_id: userId.value,
      feedback: feedback.value
    });
    success.value = `Server response: ${result.status}`;
  } catch (err) {
    error.value = err instanceof Error ? err.message : "Feedback submission failed";
  } finally {
    loading.value = false;
  }
}
</script>
