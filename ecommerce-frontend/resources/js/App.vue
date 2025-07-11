<template>
  <div class="flex flex-col min-h-screen">
    <Navbar />
    <main class="flex-grow container mx-auto p-4">
      <router-view></router-view>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';
</script>