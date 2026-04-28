<template>
  <p class="status-inline" :class="toneClass">
    <slot>{{ message }}</slot>
  </p>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    tone?: "success" | "error";
    message?: string;
  }>(),
  {
    tone: "success",
    message: ""
  }
);

const toneClass = computed(() => (props.tone === "error" ? "error" : "success"));
</script>
