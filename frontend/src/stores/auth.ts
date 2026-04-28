import { computed, ref } from "vue"
import { defineStore } from "pinia"
import { api } from "../api/client"

export const useAuthStore = defineStore("auth", () => {
  const token = ref<string | null>(localStorage.getItem("gt_token"))
  const isAuthenticated = computed(() => Boolean(token.value))
  const loading = ref(false)
  const error = ref("")

  async function login(email: string, password: string) {
    loading.value = true
    error.value = ""
    try {
      const { data } = await api.post("/v1/auth/login", { email, password })
      token.value = data.token
      localStorage.setItem("gt_token", data.token)
    } catch (e: any) {
      error.value = e?.response?.data?.error || "Login failed. Check credentials and try again."
    } finally {
      loading.value = false
    }
  }

  async function register(fullName: string, email: string, password: string) {
    loading.value = true
    error.value = ""
    try {
      const { data } = await api.post("/v1/auth/register", {
        display_name: fullName,
        email,
        password,
      })
      token.value = data.token
      localStorage.setItem("gt_token", data.token)
    } catch (e: any) {
      error.value = e?.response?.data?.error || "Registration failed. Please try again."
    } finally {
      loading.value = false
    }
  }

  function logout() {
    token.value = null
    localStorage.removeItem("gt_token")
  }

  return { token, isAuthenticated, loading, error, login, register, logout }
})
