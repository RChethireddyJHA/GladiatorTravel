import { computed, ref } from "vue"
import { defineStore } from "pinia"
import type { TravelPackage } from "../types/travel"

const KEY = "gt_wishlist"

export const useWishlistStore = defineStore("wishlist", () => {
  const ids = ref<string[]>(JSON.parse(localStorage.getItem(KEY) || "[]"))
  const items = ref<TravelPackage[]>([])

  const count = computed(() => ids.value.length)

  function persist() {
    localStorage.setItem(KEY, JSON.stringify(ids.value))
  }

  function isSaved(id: string) {
    return ids.value.includes(id)
  }

  function toggle(item: TravelPackage) {
    if (isSaved(item.id)) {
      ids.value = ids.value.filter((x) => x !== item.id)
      items.value = items.value.filter((x) => x.id !== item.id)
    } else {
      ids.value = [...ids.value, item.id]
      if (!items.value.some((x) => x.id === item.id)) {
        items.value = [item, ...items.value]
      }
    }
    persist()
  }

  function syncFromPackages(packages: TravelPackage[]) {
    const byId = new Map(packages.map((x) => [x.id, x]))
    const next = ids.value.map((id) => byId.get(id)).filter(Boolean) as TravelPackage[]
    if (next.length > 0) {
      items.value = next
    }
  }

  return { ids, items, count, isSaved, toggle, syncFromPackages }
})
