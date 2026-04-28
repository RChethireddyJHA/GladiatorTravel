<script setup lang="ts">
import { cn } from "../../lib/utils"

defineProps<{ modelValue?: string | number; class?: string; type?: string; placeholder?: string }>()
defineEmits<{ "update:modelValue": [value: string] }>()
</script>

<template>
  <input
    :value="modelValue"
    :type="type || 'text'"
    :placeholder="placeholder"
    :class="
      cn(
        'h-11 w-full rounded-xl border border-slate-300 bg-white px-3 text-sm outline-none transition focus:border-sky-500 focus:ring-2 focus:ring-sky-200',
        $props.class,
      )
    "
    @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
  />
</template>
