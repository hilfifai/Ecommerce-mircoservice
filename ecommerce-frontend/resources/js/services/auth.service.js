import { backendApi } from './api';

class AuthService {
    login(credentials) {
        return backendApi.post('/auth/login', credentials);
    }

    register(userData) {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        
        const requestOptions = {
            method: "POST",
            headers: myHeaders,
            body: JSON.stringify(userData),
            redirect: "follow",
            credentials: "include" // Untuk mengirim cookies jika diperlukan
        };

        return fetch(`${backendApi.defaults.baseURL}/auth/register`, requestOptions)
            .then(async (response) => {
                const data = await response.json();
                if (!response.ok) {
                    throw new Error(data.message || 'Registration failed');
                }
                return data;
            });
    }

    // logout() {
    // }
}

export default new AuthService();