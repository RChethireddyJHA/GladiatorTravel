<script setup lang="ts">
import { computed, reactive, ref } from "vue"
import Button from "../components/ui/Button.vue"
import Input from "../components/ui/Input.vue"
import { createBooking, fetchProfile } from "../api/travel"
import { useRoute } from "vue-router"

const route = useRoute()
const step = ref(1)
const loading = ref(false)
const error = ref("")
const success = ref("")

const form = reactive({
  fullName: "",
  email: "",
  travelers: 2,
  cardNumber: "",
  cardHolder: "",
  expiryMonth: "",
  expiryYear: "",
  cvv: "",
  billingZip: "",
})

const progress = computed(() => `${Math.round((step.value / 3) * 100)}%`)

async function submitBooking() {
  if (!/^\d{16}$/.test(form.cardNumber.replace(/\s+/g, ""))) {
    error.value = "Card number must be 16 digits."
    return
  }
  if (!/^\d{2}$/.test(form.expiryMonth) || Number(form.expiryMonth) < 1 || Number(form.expiryMonth) > 12) {
    error.value = "Invalid expiry month."
    return
  }
  if (!/^\d{2}$/.test(form.expiryYear)) {
    error.value = "Invalid expiry year."
    return
  }
  if (!/^\d{3,4}$/.test(form.cvv)) {
    error.value = "Invalid CVV."
    return
  }
  loading.value = true
  error.value = ""
  success.value = ""
  try {
    const profile = await fetchProfile().catch(() => ({ id: "1" }))
    const today = new Date()
    const start = today.toISOString().slice(0, 10)
    const end = new Date(today.getTime() + 4 * 24 * 60 * 60 * 1000).toISOString().slice(0, 10)

    await createBooking({
      user_id: Number(profile.id || 1),
      destination_id: Number(route.params.id),
      start_date: start,
      end_date: end,
      traveler_details: {
        full_name: form.fullName,
        email: form.email,
        travelers: Number(form.travelers),
      },
      payment: {
        card_holder: form.cardHolder,
        card_number: form.cardNumber.replace(/\s+/g, ""),
        expiry_month: form.expiryMonth,
        expiry_year: form.expiryYear,
        cvv: form.cvv,
        billing_zip: form.billingZip,
      },
    })
    success.value = "Booking confirmed. Check your dashboard for details."
  } catch {
    error.value = "Booking failed due to an API error. Please retry."
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <section class="mx-auto max-w-3xl rounded-2xl border border-slate-200 bg-white p-6">
    <h1 class="text-2xl font-semibold">Checkout</h1>
    <div class="mt-4 h-2 w-full rounded-full bg-slate-200">
      <div class="h-2 rounded-full bg-sky-600 transition-all" :style="{ width: progress }" />
    </div>
    <p class="mt-2 text-sm text-slate-600">Step {{ step }} of 3</p>

    <div v-if="step === 1" class="mt-6 space-y-3">
      <h2 class="font-semibold">Traveler details</h2>
      <Input v-model="form.fullName" placeholder="Full name" />
      <Input v-model="form.email" placeholder="Email" />
      <Input v-model="form.travelers" type="number" placeholder="Travelers" />
    </div>

    <div v-else-if="step === 2" class="mt-6 space-y-2">
      <h2 class="font-semibold">Booking summary</h2>
      <p class="text-slate-600">Package ID: {{ route.params.id }}</p>
      <p class="text-slate-600">Traveler: {{ form.fullName || "Not provided" }}</p>
      <p class="text-slate-600">Passengers: {{ form.travelers }}</p>
    </div>

    <div v-else class="mt-6 space-y-3">
      <h2 class="font-semibold">Payment</h2>
      <Input v-model="form.cardHolder" placeholder="Card holder name" />
      <Input v-model="form.cardNumber" placeholder="Card number (16 digits)" />
      <div class="grid gap-3 sm:grid-cols-3">
        <Input v-model="form.expiryMonth" placeholder="MM" />
        <Input v-model="form.expiryYear" placeholder="YY" />
        <Input v-model="form.cvv" placeholder="CVV" />
      </div>
      <Input v-model="form.billingZip" placeholder="Billing ZIP / Postal code" />
      <p class="text-sm text-slate-500">All card fields are collected and validated before creating booking.</p>
    </div>

    <p v-if="error" class="mt-4 rounded-xl bg-red-50 p-3 text-sm text-red-700">{{ error }}</p>
    <p v-if="success" class="mt-4 rounded-xl bg-emerald-50 p-3 text-sm text-emerald-700">{{ success }}</p>

    <div class="mt-6 flex justify-between">
      <Button variant="outline" :disabled="step === 1 || loading" @click="step--">Back</Button>
      <Button v-if="step < 3" :disabled="loading" @click="step++">Next</Button>
      <Button v-else :disabled="loading" @click="submitBooking">{{ loading ? "Processing..." : "Confirm booking" }}</Button>
    </div>
  </section>
</template>
