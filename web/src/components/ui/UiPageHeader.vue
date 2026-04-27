<template>
  <header class="page-header panel-elevated">
    <p v-if="kicker" class="page-kicker">{{ kicker }}</p>
    <h2 class="page-title">{{ title }}</h2>
    <p v-if="subtitle" class="page-subtitle">{{ subtitle }}</p>
    <slot />
  </header>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  kicker?: string;
  subtitle?: string;
}>();
</script>
