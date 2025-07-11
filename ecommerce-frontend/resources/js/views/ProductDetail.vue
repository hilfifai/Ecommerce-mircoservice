<template>
  <div v-if="product" class="grid grid-cols-1 md:grid-cols-2 gap-10">
    <div>
      <!-- Gambar produk bisa ditambahkan di sini -->
      <div class="bg-gray-200 w-full h-96 rounded-lg flex items-center justify-center">
        <span class="text-gray-500">Product Image</span>
      </div>
    </div>
    <div>
      <h1 class="text-4xl font-bold mb-4">{{ product.name }}</h1>
      <p class="text-gray-700 mb-6 text-lg">{{ product.description }}</p>
      <p class="text-3xl font-semibold text-green-600 mb-6">${{ product.price ? product.price.toFixed(2) : '0.00' }}</p>
      <div class="flex items-center gap-4">
        <button @click="cartStore.addToCart(product)"
          class="bg-blue-600 text-white font-bold py-3 px-8 rounded-lg hover:bg-blue-700 transition duration-300">
          Add to Cart
        </button>
        <p class="text-sm text-gray-500" v-if="product.stock > 0">{{ product.stock }} in stock</p>
        <p class="text-sm text-red-500" v-else>Out of stock</p>
      </div>
    </div>
  </div>
  <div v-else-if="loading" class="text-center py-20">
    <p>Loading product details...</p>
  </div>
  <div v-else class="text-center py-20">
    <h2 class="text-2xl">Product not found.</h2>
    <router-link to="/" class="mt-4 inline-block text-blue-600 hover:underline">Go back to shop</router-link>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import ProductService from '../services/product.service';
import { useCartStore } from '../store/cart.store';

const route = useRoute();
const cartStore = useCartStore();
const product = ref(null);
const loading = ref(true);

onMounted(async () => {
  try {
    const productId = route.params.id;
    const response = await ProductService.getProductById(productId);
    product.value = response.data.data;
  } catch (error) {
    console.error("Failed to fetch product details:", error);
  } finally {
    loading.value = false;
  }
});
</script>