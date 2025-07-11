<template>
  <div class="max-w-md mx-auto mt-10">
    <h2 class="text-2xl font-bold text-center">Create an Account</h2>
    <form @submit.prevent="handleRegister" class="mt-8 space-y-6">
      <div v-if="errorMessage" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
        <span class="block sm:inline">{{ errorMessage }}</span>
      </div>
      <div>
        <label for="name">Full Name</label>
        <input id="name" v-model="form.name" type="text" required class="w-full px-3 py-2 border border-gray-300 rounded-md">
      </div>
      <div>
        <label for="email">Email address</label>
        <input id="email" v-model="form.email" type="email" required class="w-full px-3 py-2 border border-gray-300 rounded-md">
      </div>
      <div>
        <label for="password">Password</label>
        <input id="password" v-model="form.password" type="password" required class="w-full px-3 py-2 border border-gray-300 rounded-md">
      </div>
      <div>
        <button type="submit" :disabled="isSubmitting" class="w-full py-2 px-4 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:bg-gray-400">
          {{ isSubmitting ? 'Registering...' : 'Register' }}
        </button>
      </div>
       <div class="text-center">
        <p>Already have an account? <router-link to="/login" class="text-blue-600 hover:underline">Login here</router-link></p>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import AuthService from '../services/auth.service';

const router = useRouter();
const form = reactive({
  name: '',
  email: '',
  password: '',
});
const isSubmitting = ref(false);
const errorMessage = ref('');

const handleRegister = async () => {
  isSubmitting.value = true;
  errorMessage.value = '';
  try {
    await AuthService.register(form);
    alert('Registration successful! Please login.');
    router.push('/login');
  } catch (error) {
    errorMessage.value = error.response?.data?.error || 'Registration failed. Please try again.';
    console.error(error);
  } finally {
    isSubmitting.value = false;
  }
};
</script>