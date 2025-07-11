<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Our Products</h1>

    <!-- Skeleton Loading State -->
    <div v-if="isLoading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <ProductCardSkeleton v-for="n in 8" :key="n" />
    </div>

    <!-- Product Grid -->
    <div v-else-if="productStore.products.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="product in productStore.products"
        :key="product.id"
        :product="product"
      />
    </div>
    
    <!-- Empty State -->
    <div v-else class="text-center py-16 bg-white rounded-lg shadow-md">
      <h2 class="text-2xl font-semibold text-gray-700">No Products Found</h2>
      <p class="text-gray-500 mt-2">Please check back later or contact support.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useProductStore } from '../store/product.store';
import ProductCard from '../components/ProductCard.vue';
import ProductCardSkeleton from '../components/ProductCardSkeleton.vue';

const productStore = useProductStore();
const isLoading = ref(true);

onMounted(async () => {
  isLoading.value = true;
  await productStore.fetchProducts();
  isLoading.value = false;
});
</script>