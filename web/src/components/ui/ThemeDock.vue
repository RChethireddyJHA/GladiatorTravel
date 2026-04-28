<template>
  <div ref="rootEl" class="theme-dock">
    <button
      ref="triggerEl"
      type="button"
      class="theme-dock-trigger"
      :aria-label="ariaLabel"
      :aria-expanded="isOpen"
      aria-haspopup="menu"
      :aria-controls="menuId"
      @click="toggleMenu"
      @keydown="onTriggerKeydown"
    >
      <span class="theme-dock-trigger-icon" aria-hidden="true">🎨</span>
      <span class="theme-dock-trigger-text">Theme: {{ selectedOption?.label ?? fallbackLabel }}</span>
      <span class="theme-dock-trigger-caret" aria-hidden="true">▾</span>
    </button>

    <Teleport to="body" v-if="isOpen">
      <ul
        ref="menuEl"
        :id="menuId"
        class="theme-dock-menu"
        role="menu"
        aria-label="Theme options"
        :style="menuPositionStyle"
      >
        <li v-for="(option, index) in options" :key="option.id" role="none">
          <button
            ref="optionRefs"
            type="button"
            class="theme-dock-option"
            role="menuitemradio"
            :aria-checked="option.id === modelValue"
            :class="{ active: option.id === modelValue }"
            @click="selectOption(option.id)"
            @keydown="onOptionKeydown(index, $event)"
          >
            <span aria-hidden="true">{{ option.icon }}</span>
            <span>{{ option.label }}</span>
          </button>
        </li>
      </ul>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, onBeforeUnmount, onMounted, ref } from "vue";

type ThemeOption = {
  id: string;
  label: string;
  icon: string;
};

const props = withDefaults(
  defineProps<{
    modelValue: string;
    options: ThemeOption[];
    ariaLabel?: string;
  }>(),
  {
    ariaLabel: "Select global theme"
  }
);

const emit = defineEmits<{
  (event: "update:modelValue", value: string): void;
}>();

const menuId = "theme-dock-menu";
const fallbackLabel = "Sunset";
const isOpen = ref(false);
const focusedIndex = ref(0);
const rootEl = ref<HTMLElement | null>(null);
const menuEl = ref<HTMLElement | null>(null);
const triggerEl = ref<HTMLButtonElement | null>(null);
const optionRefs = ref<HTMLButtonElement[]>([]);
const menuPosition = ref({ top: 0, right: 0 });

const selectedIndex = computed(() => {
  const index = props.options.findIndex((item) => item.id === props.modelValue);
  return index >= 0 ? index : 0;
});

const selectedOption = computed(() => props.options[selectedIndex.value]);

const menuPositionStyle = computed(() => ({
  top: `${menuPosition.value.top}px`,
  right: `${menuPosition.value.right}px`
}));

function updateMenuPosition() {
  if (!triggerEl.value) return;
  const rect = triggerEl.value.getBoundingClientRect();
  menuPosition.value = {
    top: rect.bottom + 6,
    right: window.innerWidth - rect.right
  };
}

function closeMenu() {
  isOpen.value = false;
}

function focusOption(index: number) {
  focusedIndex.value = index;
  nextTick(() => {
    optionRefs.value[index]?.focus();
  });
}

function openMenu(focusSelected = true) {
  if (!props.options.length) return;
  isOpen.value = true;
  updateMenuPosition();
  const startIndex = focusSelected ? selectedIndex.value : 0;
  focusOption(startIndex);
}

function toggleMenu() {
  if (isOpen.value) {
    closeMenu();
    return;
  }
  openMenu();
}

function selectOption(id: string) {
  emit("update:modelValue", id);
  closeMenu();
  nextTick(() => triggerEl.value?.focus());
}

function onTriggerKeydown(event: KeyboardEvent) {
  if (event.key === "ArrowDown" || event.key === "Enter" || event.key === " ") {
    event.preventDefault();
    openMenu();
  }
}

function onOptionKeydown(index: number, event: KeyboardEvent) {
  if (!props.options.length) return;

  if (event.key === "ArrowDown") {
    event.preventDefault();
    focusOption((index + 1) % props.options.length);
    return;
  }

  if (event.key === "ArrowUp") {
    event.preventDefault();
    focusOption((index - 1 + props.options.length) % props.options.length);
    return;
  }

  if (event.key === "Home") {
    event.preventDefault();
    focusOption(0);
    return;
  }

  if (event.key === "End") {
    event.preventDefault();
    focusOption(props.options.length - 1);
    return;
  }

  if (event.key === "Escape") {
    event.preventDefault();
    closeMenu();
    nextTick(() => triggerEl.value?.focus());
    return;
  }

  if (event.key === "Tab") {
    closeMenu();
    return;
  }

  if (event.key === "Enter" || event.key === " ") {
    event.preventDefault();
    selectOption(props.options[index].id);
  }
}

function onPointerDown(event: PointerEvent) {
  if (!isOpen.value) return;
  const target = event.target as Node | null;
  if (!target) return;
  const clickedTriggerArea = rootEl.value?.contains(target) ?? false;
  const clickedMenu = menuEl.value?.contains(target) ?? false;
  if (!clickedTriggerArea && !clickedMenu) {
    closeMenu();
  }
}

function onWindowResize() {
  if (isOpen.value) {
    updateMenuPosition();
  }
}

onMounted(() => {
  window.addEventListener("pointerdown", onPointerDown);
  window.addEventListener("resize", onWindowResize);
});

onBeforeUnmount(() => {
  window.removeEventListener("pointerdown", onPointerDown);
  window.removeEventListener("resize", onWindowResize);
});
</script>

<style scoped>
.theme-dock {
  position: relative;
}

.theme-dock-trigger {
  width: auto;
  border-radius: var(--radius-pill);
  border: 1px solid var(--color-border);
  background: rgba(255, 255, 255, 0.68);
  color: var(--color-ink);
  padding: 0.44rem 0.72rem;
  display: inline-flex;
  align-items: center;
  gap: 0.45rem;
  font-size: 0.82rem;
  font-weight: 700;
  transition:
    background-color var(--duration-theme) ease,
    border-color var(--duration-theme) ease,
    color var(--duration-theme) ease,
    box-shadow var(--duration-fast) ease;
}

.theme-dock-trigger:hover {
  transform: none;
  filter: none;
  border-color: var(--color-border-strong);
}

.theme-dock-trigger:focus-visible {
  outline: none;
  box-shadow: 0 0 0 3px rgba(29, 143, 125, 0.18);
}

.theme-dock-trigger-icon,
.theme-dock-trigger-caret {
  line-height: 1;
}

.theme-dock-trigger-text {
  white-space: nowrap;
}

:global(.theme-dock-menu) {
  position: fixed;
  width: 170px;
  margin: 0;
  padding: 0.34rem;
  list-style: none;
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border);
  background: rgba(255, 255, 255, 0.93);
  box-shadow: var(--shadow-sm);
  backdrop-filter: blur(8px);
  z-index: 9999;
  animation: menu-pop var(--duration-fast) ease;
}

.theme-dock-option {
  width: 100%;
  border: 1px solid transparent;
  border-radius: var(--radius-sm);
  background: transparent;
  color: var(--color-ink);
  display: flex;
  align-items: center;
  gap: 0.48rem;
  padding: 0.45rem 0.52rem;
  font-size: 0.82rem;
  font-weight: 700;
  text-align: left;
}

.theme-dock-option:hover,
.theme-dock-option:focus-visible {
  transform: none;
  filter: none;
  border-color: var(--color-border-strong);
  background: var(--color-accent-soft);
  outline: none;
}

.theme-dock-option.active {
  border-color: var(--color-accent);
  background: var(--nav-active-gradient);
}

@keyframes menu-pop {
  from {
    opacity: 0;
    transform: translateY(-4px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 620px) {
  .theme-dock-trigger {
    font-size: 0.78rem;
    padding: 0.4rem 0.64rem;
  }

  :global(.theme-dock-menu) {
    left: auto;
  }
}
</style>
