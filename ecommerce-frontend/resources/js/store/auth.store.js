import { defineStore } from 'pinia';
import { ref } from 'vue';
import { backendApi } from '../services/api';
import router from '../router';

export const useAuthStore = defineStore('auth', () => {
    const user = ref(null);
    const token = ref(localStorage.getItem('token') || null);

    async function login(credentials) {
        try {
            const response = await backendApi.post('/auth/login', credentials);
            token.value = response.data.token;
            localStorage.setItem('token', token.value);
            // Fetch user profile if needed
            router.push('/');
        } catch (error) {
            console.error("Login failed:", error);
            throw error;
        }
    }

    function logout() {
        user.value = null;
        token.value = null;
        localStorage.removeItem('token');
        router.push('/login');
    }

    return { user, token, login, logout };
});