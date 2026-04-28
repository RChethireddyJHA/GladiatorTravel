import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "home", component: () => import("../views/HomeView.vue") },
    { path: "/search", name: "search", component: () => import("../views/SearchView.vue") },
    { path: "/packages/:id", name: "package-details", component: () => import("../views/PackageDetailsView.vue") },
    { path: "/booking/:id", name: "booking", component: () => import("../views/BookingFlowView.vue") },
    { path: "/dashboard", name: "dashboard", component: () => import("../views/DashboardView.vue") },
    { path: "/auth", name: "auth", component: () => import("../views/AuthView.vue") },
  ],
})

router.beforeEach((to) => {
  const token = localStorage.getItem("gt_token")
  if (to.name === "dashboard" && !token) {
    return { name: "auth", query: { redirect: "/dashboard" } }
  }
  return true
})

export default router
