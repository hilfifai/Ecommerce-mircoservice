<template>
  <div>
    <h1 class="text-3xl font-bold mb-6">Shopping Cart</h1>
    <div v-if="cartStore.items.length > 0">
      <div class="bg-white shadow-md rounded-lg p-6">
        <!-- Cart Items -->
        <div v-for="item in cartStore.items" :key="item.id" class="flex items-center justify-between border-b py-4">
          <div>
            <h2 class="font-semibold text-lg">{{ item.name }}</h2>
            <p class="text-gray-600">Quantity: {{ item.quantity }}</p>
          </div>
          <div class="text-right">
            <p class="font-semibold">${{ (item.price * item.quantity).toFixed(2) }}</p>
            <button @click="cartStore.removeFromCart(item.id)" class="text-red-500 hover:text-red-700 text-sm mt-1">
              Remove
            </button>
          </div>
        </div>

        <!-- Cart Summary -->
        <div class="mt-6 text-right">
          <p class="text-xl font-bold">Total: ${{ cartStore.cartTotal.toFixed(2) }}</p>
          <button @click="goToCheckout" class="mt-4 bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-6 rounded-lg transition duration-300">
            Proceed to Checkout
          </button>
        </div>
      </div>
    </div>
    <div v-else class="text-center py-10">
      <p class="text-xl text-gray-500">Your cart is empty.</p>
      <router-link to="/" class="mt-4 inline-block bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700">
        Continue Shopping
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '../store/cart.store';
import { useRouter } from 'vue-router';

const cartStore = useCartStore();
const router = useRouter();

const goToCheckout = () => {
  router.push('/checkout');
};


// export const useCartStore = defineStore('cart', () => {
//   ...
//   function removeFromCart(productId) {
//     items.value = items.value.filter(item => item.id !== productId);
//   }
//   return { ..., removeFromCart };
// });
</script>