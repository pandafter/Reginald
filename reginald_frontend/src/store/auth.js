import { defineStore } from 'pinia';
import api from '../api/axios';

export const useAuthStore = defineStore('auth', {
    state: () => {
        const savedUser = localStorage.getItem('user');
        let user = null;
        try {
            if (savedUser && savedUser !== 'undefined') {
                user = JSON.parse(savedUser);
            }
        } catch (e) {
            console.error('Error parsing user from localStorage', e);
            localStorage.removeItem('user');
        }

        return {
            user,
            loading: false,
            error: null,
            isVerifying: false,
        };
    },

    getters: {
        isAuthenticated: (state) => !!state.user,
        isAdmin: (state) => state.user?.role === 'admin',
    },

    actions: {
        async login(email, password) {
            this.loading = true;
            this.error = null;
            try {
                const response = await api.post('/auth/login', { email, password });
                this.user = response.data.user;

                // We only store user info, token is in HTTP-only cookie
                localStorage.setItem('user', JSON.stringify(this.user));

                return true;
            } catch (err) {
                this.error = err.response?.data?.message || 'Error al iniciar sesión';
                return false;
            } finally {
                this.loading = false;
            }
        },

        async logout() {
            try {
                await api.post('/auth/logout');
            } catch (e) {
                console.error('Error during logout api call', e);
            } finally {
                this.user = null;
                localStorage.removeItem('user');
                window.location.href = '/login';
            }
        },

        async checkSession() {
            if (this.isVerifying) return;
            this.isVerifying = true;
            try {
                const response = await api.get('/auth/verify');
                // Session is valid, update user info (important if name or role changed)
                this.user = response.data;
                localStorage.setItem('user', JSON.stringify(this.user));
                return true;
            } catch (e) {
                console.warn('Session verification failed, logging out...');
                this.logout();
                return false;
            } finally {
                this.isVerifying = false;
            }
        }
    },
});
