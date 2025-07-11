import axios from 'axios';
import { useAuthStore } from '../store/auth.store';

const backendApi = axios.create({
    baseURL: 'http://localhost:8080/api/v1', // URL Backend Golang
});

const directusApi = axios.create({
    baseURL: 'http://localhost:8055', // URL Directus
});

backendApi.interceptors.request.use(config => {
    const authStore = useAuthStore();
    const token = authStore.token;
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export { backendApi, directusApi };