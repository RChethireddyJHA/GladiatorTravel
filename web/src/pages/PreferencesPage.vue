<template>
  <section class="page">
    <UiPageHeader
      title="Traveler Preferences"
      subtitle="Standardized profile controls for budget, late-night preference, and tag exclusions."
    />

    <UiPanel>
      <div class="form-grid">
        <UiField label="User ID">
          <input type="number" min="1" v-model.number="form.user_id" />
        </UiField>
        <UiField label="Budget level (1-4)">
          <input type="number" min="1" max="4" v-model.number="form.budget_level" />
        </UiField>
        <UiField label="Prefers late night">
          <select v-model="form.prefers_late_night">
            <option :value="true">true</option>
            <option :value="false">false</option>
          </select>
        </UiField>
        <UiField label="Avoid tags">
          <input type="text" v-model="form.avoid_tags" placeholder="quiet_only" />
        </UiField>
      </div>

      <div class="actions-row">
        <UiButton :loading="loading" text="Save Preferences" loading-text="Saving..." @click="submit" />
      </div>

      <UiStatusMessage v-if="success" tone="success">{{ success }}</UiStatusMessage>
      <UiStatusMessage v-if="error" tone="error">{{ error }}</UiStatusMessage>
    </UiPanel>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { api } from "../lib/api";
import UiButton from "../components/ui/UiButton.vue";
import UiField from "../components/ui/UiField.vue";
import UiPageHeader from "../components/ui/UiPageHeader.vue";
import UiPanel from "../components/ui/UiPanel.vue";
import UiStatusMessage from "../components/ui/UiStatusMessage.vue";

const loading = ref(false);
const success = ref("");
const error = ref("");

const form = reactive({
  user_id: 1,
  budget_level: 3,
  prefers_late_night: true,
  avoid_tags: "quiet_only"
});

async function submit() {
  loading.value = true;
  success.value = "";
  error.value = "";

  try {
    const result = await api.savePreferences(form);
    success.value = `Server response: ${result.status}`;
  } catch (err) {
    error.value = err instanceof Error ? err.message : "Preference save failed";
  } finally {
    loading.value = false;
  }
}
</script>
