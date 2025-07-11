import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { backendApi } from '../services/api';
import { useAuthStore } from './auth.store';

export const useCartStore = defineStore('cart', () => {
    const items = ref([]);

    const cartTotal = computed(() => {
        return items.value.reduce((total, item) => total + item.price * item.quantity, 0);
    });

    function addToCart(product) {
        const existingItem = items.value.find(item => item.id === product.id);
        if (existingItem) {
            existingItem.quantity++;
        } else {
            items.value.push({ ...product, quantity: 1 });
        }
    }

    function clearCart() {
        items.value = [];
    }

    async function checkout() {
        const authStore = useAuthStore();
        if (!authStore.token) {
            alert('Please login to checkout.');
            return;
        }

        const orderPayload = {
            // User ID would ideally come from a decoded JWT or a /me endpoint
            user_id: 1, // Placeholder
            items: items.value.map(item => ({
                product_id: item.id,
                quantity: item.quantity,
            }))
        };
        
        try {
            await backendApi.post('/orders', orderPayload);
            alert('Order placed successfully! It is being processed.');
            clearCart();
        } catch(error) {
            console.error("Checkout failed:", error);
            alert('Failed to place order.');
        }
    }

    return { items, addToCart, cartTotal, checkout };
});