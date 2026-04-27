<template>
  <article class="kpi-card">
    <p class="kpi-label">{{ label }}</p>
    <p class="kpi-value">{{ value }}</p>
  </article>
</template>

<script setup lang="ts">
defineProps<{
  label: string;
  value: string | number;
}>();
</script>
