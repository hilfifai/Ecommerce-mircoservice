import { defineStore } from 'pinia';
import { ref } from 'vue';
import { directusApi } from '../services/api';

export const useProductStore = defineStore('product', () => {
    const products = ref([]);
    
    async function fetchProducts() {
        try {
            const response = await directusApi.get('/items/products');
            products.value = response.data.data;
        } catch (error) {
            console.error("Failed to fetch products from Directus:", error);
        }
    }
    
    return { products, fetchProducts };
});