<template>
  <header class="bg-white shadow-md sticky top-0 z-10">
    <nav class="container mx-auto px-4 md:px-6 py-3 flex justify-between items-center">
      <router-link to="/" class="text-2xl font-bold text-gray-800 hover:text-blue-600 transition-colors">
        Hilfi-Store
      </router-link>
      <div class="flex items-center space-x-4 md:space-x-6 text-gray-600">
        <router-link to="/cart" class="relative hover:text-blue-600 transition-colors">
          <!-- SVG Icon for Cart -->
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          <span v-if="cartStore.items.length > 0" class="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full h-5 w-5 flex items-center justify-center">
            {{ cartStore.items.length }}
          </span>
        </router-link>
        <div class="hidden md:flex items-center space-x-4">
          <template v-if="!authStore.token">
            <router-link to="/login" class="hover:text-blue-600 transition-colors">Login</router-link>
            <router-link to="/register" class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
              Register
            </router-link>
          </template>
          <template v-else>
            <button @click="authStore.logout()" class="hover:text-blue-600 transition-colors">Logout</button>
          </template>
        </div>
        <!-- Tambahkan menu mobile jika diperlukan -->
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useAuthStore } from '../store/auth.store';
import { useCartStore } from '../store/cart.store';

const authStore = useAuthStore();
const cartStore = useCartStore();
</script>