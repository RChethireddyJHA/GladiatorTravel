<script setup lang="ts">
import { cn } from "../../lib/utils"

defineProps<{ class?: string }>()
</script>

<template>
  <article :class="cn('rounded-2xl border border-slate-200 bg-white p-4 shadow-sm', $props.class)">
    <slot />
  </article>
</template>
