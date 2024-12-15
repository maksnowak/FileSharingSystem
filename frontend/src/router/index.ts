import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/Login.vue'
import CreateAccount from '@/views/CreateAccount.vue'

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/home', component: Home },
    { path: '/login', component: Login },
    { path: '/create-account', component: CreateAccount }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;