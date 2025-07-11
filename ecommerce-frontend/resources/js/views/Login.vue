<template>
  <div class="max-w-md mx-auto mt-10">
    <h2 class="text-2xl font-bold text-center">Login</h2>
    <form @submit.prevent="handleLogin" class="mt-8 space-y-6">
      <div>
        <label for="email">Email address</label>
        <input id="email" v-model="email" type="email" required class="w-full px-3 py-2 border border-gray-300 rounded-md">
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" required class="w-full px-3 py-2 border border-gray-300 rounded-md">
      </div>
      <div>
        <button type="submit" class="w-full py-2 px-4 bg-blue-600 text-white rounded-md">
          Sign in
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../store/auth.store';

const email = ref('');
const password = ref('');
const authStore = useAuthStore();

const handleLogin = async () => {
  try {
    await authStore.login({ email: email.value, password: password.value });
  } catch (error) {
    alert('Invalid credentials');
  }
};
</script>