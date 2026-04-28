import { createRouter, createWebHistory } from "vue-router";

import DestinationsPage from "./pages/DestinationsPage.vue";
import PreferencesPage from "./pages/PreferencesPage.vue";
import FeedbackPage from "./pages/FeedbackPage.vue";
import TripsPage from "./pages/TripsPage.vue";
import ItineraryPage from "./pages/ItineraryPage.vue";

const routes = [
  { path: "/", name: "destinations", component: DestinationsPage },
  { path: "/destinations", redirect: "/" },
  { path: "/preferences", name: "preferences", component: PreferencesPage },
  { path: "/feedback", name: "feedback", component: FeedbackPage },
  { path: "/trips", name: "trips", component: TripsPage },
  { path: "/itinerary", name: "itinerary", component: ItineraryPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
