<script setup lang="ts">
import { computed } from "vue"
import { cn } from "../../lib/utils"

const props = withDefaults(
  defineProps<{
    variant?: "default" | "secondary" | "outline" | "ghost"
    class?: string
    type?: "button" | "submit" | "reset"
  }>(),
  {
    variant: "default",
    type: "button",
  },
)

const classes = computed(() =>
  cn(
    "inline-flex items-center justify-center rounded-xl px-4 py-2.5 text-sm font-medium transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 disabled:pointer-events-none disabled:opacity-50",
    {
      "bg-sky-600 text-white hover:bg-sky-700": props.variant === "default",
      "bg-amber-100 text-slate-900 hover:bg-amber-200": props.variant === "secondary",
      "border border-slate-300 bg-white text-slate-900 hover:bg-slate-50": props.variant === "outline",
      "text-slate-700 hover:bg-slate-100": props.variant === "ghost",
    },
    props.class,
  ),
)
</script>

<template>
  <button :type="type" :class="classes">
    <slot />
  </button>
</template>
