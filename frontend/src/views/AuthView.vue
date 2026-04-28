<script setup lang="ts">
import { ref } from "vue"
import { useRoute, useRouter } from "vue-router"
import Input from "../components/ui/Input.vue"
import Button from "../components/ui/Button.vue"
import { useAuthStore } from "../stores/auth"

const auth = useAuthStore()
const router = useRouter()
const route = useRoute()
const isLogin = ref(true)
const fullName = ref("")
const email = ref("")
const password = ref("")

async function submit() {
  if (isLogin.value) {
    await auth.login(email.value, password.value)
  } else {
    await auth.register(fullName.value, email.value, password.value)
  }
  if (auth.token) {
    router.push(String(route.query.redirect || "/dashboard"))
  }
}
</script>

<template>
  <section class="mx-auto max-w-md rounded-2xl border border-slate-200 bg-white p-6">
    <h1 class="text-2xl font-semibold">{{ isLogin ? "Login" : "Register" }}</h1>
    <p class="mt-1 text-sm text-slate-600">JWT/session-ready auth flow with backend integration.</p>
    <div class="mt-4 space-y-3">
      <Input v-if="!isLogin" v-model="fullName" placeholder="Full name" />
      <Input v-model="email" placeholder="Email" />
      <Input v-model="password" type="password" placeholder="Password" />
      <p v-if="auth.error" class="rounded-xl bg-red-50 p-3 text-sm text-red-700">{{ auth.error }}</p>
      <Button class="w-full" :disabled="auth.loading" @click="submit">{{ auth.loading ? "Please wait..." : isLogin ? "Login" : "Create account" }}</Button>
      <Button variant="ghost" class="w-full" @click="isLogin = !isLogin">
        {{ isLogin ? "Need an account? Register" : "Already have an account? Login" }}
      </Button>
    </div>
  </section>
</template>
