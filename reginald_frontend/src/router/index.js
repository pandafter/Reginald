import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/auth';

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/Login.vue'),
        meta: { guest: true }
    },
    {
        path: '/register',
        name: 'register',
        component: () => import('../views/Register.vue'),
        meta: { guest: true }
    },
    {
        path: '/',
        redirect: '/dashboard'
    },
    {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../views/Dashboard.vue'),
        meta: { requiresAuth: true }
    },
    {
        path: '/requests',
        name: 'requests',
        component: () => import('../views/Requests.vue'),
        meta: { requiresAuth: true }
    },
    {
        path: '/users',
        name: 'users',
        component: () => import('../views/Users.vue'),
        meta: { requiresAuth: true, requiresAdmin: true }
    },
    {
        path: '/monitoring',
        name: 'monitoring',
        component: () => import('../views/Monitoring.vue'),
        meta: { requiresAuth: true, requiresAdmin: true }
    },
    {
        path: '/:pathMatch(.*)*',
        redirect: '/'
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach((to, from, next) => {
    const auth = useAuthStore();

    if (to.meta.requiresAuth && !auth.isAuthenticated) {
        next('/login');
    } else if (to.meta.guest && auth.isAuthenticated) {
        next('/dashboard');
    } else if (to.meta.requiresAdmin && !auth.isAdmin) {
        next('/dashboard');
    } else {
        next();
    }
});

export default router;
