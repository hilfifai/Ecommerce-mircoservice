<template>
  <div class="bg-white border border-gray-200 rounded-lg shadow-lg overflow-hidden group transform hover:-translate-y-2 transition-all duration-300">
    <router-link :to="`/product/${product.id}`">
      <div class="relative">
        <!-- Placeholder untuk Gambar Produk -->
        <div class="w-full h-48 bg-gray-200 flex items-center justify-center">
            <svg class="w-16 h-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path></svg>
        </div>
        <div class="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300">
          <span class="text-white font-bold">View Details</span>
        </div>
      </div>
    </router-link>
    <div class="p-4">
      <h3 class="text-lg font-bold text-gray-800 truncate">{{ product.name }}</h3>
      <p class="text-sm text-gray-600 mt-1 h-10 line-clamp-2">{{ product.description }}</p>
      <div class="mt-4 flex justify-between items-center">
        <span class="text-xl font-extrabold text-gray-900">${{ product.price ? product.price.toFixed(2) : '0.00' }}</span>
        <button @click="addToCart" class="bg-blue-600 text-white px-4 py-2 rounded-md font-semibold hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 transition-colors">
          Add to Cart
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '../store/cart.store';
import { defineProps } from 'vue';

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
});

const cartStore = useCartStore();

const addToCart = () => {
  cartStore.addToCart(props.product);
  // Mungkin tambahkan notifikasi "Added to cart!" di sini
};
</script>

<style>
/* Untuk line-clamp, jika Tailwind JIT tidak mengambilnya */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;  
  overflow: hidden;
}
</style>