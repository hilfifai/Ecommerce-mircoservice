<template>
  <div>
    <h1 class="text-3xl font-bold mb-6">Checkout</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <!-- Order Summary -->
      <div>
        <h2 class="text-xl font-semibold mb-4 border-b pb-2">Order Summary</h2>
        <div v-if="cartStore.items.length > 0">
          <div v-for="item in cartStore.items" :key="item.id" class="flex justify-between mb-2">
            <span>{{ item.name }} x {{ item.quantity }}</span>
            <span class="font-medium">${{ (item.price * item.quantity).toFixed(2) }}</span>
          </div>
          <div class="border-t mt-4 pt-4 flex justify-between font-bold text-lg">
            <span>Total</span>
            <span>${{ cartStore.cartTotal.toFixed(2) }}</span>
          </div>
        </div>
      </div>

      <!-- User & Payment Info -->
      <div>
        <h2 class="text-xl font-semibold mb-4 border-b pb-2">Customer Information</h2>
        <div v-if="authStore.user">
          <p><strong>Name:</strong> {{ authStore.user.name }}</p>
          <p><strong>Email:</strong> {{ authStore.user.email }}</p>
        </div>
        <div v-else>
          <p class="text-red-500">Please <router-link to="/login" class="underline">login</router-link> to proceed.</p>
        </div>

        <button @click="handleCheckout" :disabled="isProcessing || !authStore.token"
          class="mt-8 w-full bg-blue-600 text-white py-3 rounded-lg font-bold hover:bg-blue-700 disabled:bg-gray-400">
          {{ isProcessing ? 'Placing Order...' : 'Place Order' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useCartStore } from '../store/cart.store';
import { useAuthStore } from '../store/auth.store';
import { useRouter } from 'vue-router';
import order_service from '../services/order.service';

const cartStore = useCartStore();
const authStore = useAuthStore();
const router = useRouter();
const isProcessing = ref(false);

const handleCheckout = async () => {
  isProcessing.value = true;
  try {
    await cartStore.checkout(); // Logika checkout ada di dalam store
    alert('Order placed successfully! It is being processed.');
    router.push('/'); // Arahkan ke halaman utama setelah berhasil
  } catch (error) {
    alert('Failed to place order. Please try again.');
    console.error("Checkout failed:", error);
  } finally {
    isProcessing.value = false;
  }
};
</script>