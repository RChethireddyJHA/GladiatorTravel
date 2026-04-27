<template>
  <button :type="type" :disabled="disabled || loading" :class="className">
    {{ loading ? loadingText : text }}
  </button>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    text: string;
    loading?: boolean;
    loadingText?: string;
    disabled?: boolean;
    variant?: "primary" | "secondary" | "ghost";
    type?: "button" | "submit" | "reset";
  }>(),
  {
    loading: false,
    loadingText: "Working...",
    disabled: false,
    variant: "primary",
    type: "button"
  }
);

const className = computed(() => {
  if (props.variant === "secondary") return "secondary";
  if (props.variant === "ghost") return "ghost";
  return "";
});
</script>
