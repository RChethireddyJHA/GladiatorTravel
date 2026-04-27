<template>
  <div class="field">
    <label :for="forId">{{ label }}</label>
    <slot />
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string;
  forId?: string;
}>();
</script>
