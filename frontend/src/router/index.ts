import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/Login.vue'
import CreateAccount from '@/views/CreateAccount.vue'
import AccountView from '@/views/AccountView.vue'
import { userStore } from "@/user";

const routes = [
  { path: '/', redirect: '/home', meta: { requiresRedirect: true } },
  { path: '/home', component: Home, meta: { requiresAuth: true } },
  { path: '/shared', component: Home, meta: { requiresAuth: true } },
  { path: '/login', component: Login },
  { path: '/create-account', component: CreateAccount },
  { path: '/account', component: AccountView, meta: { requiresAuth: true } }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const user = userStore().getUser; // Directly access the getter to get the user

  if (to.meta.requiresAuth && !user) {
    next("/login"); // Redirect to login if not authenticated
  } else if ((to.path === "/login" || to.path === "/create-account") && user) {
    next("/home"); // Redirect to home if already logged in
  } else if (to.path === "/" || to.path === "/default") {
    next(user ? "/home" : "/login"); // Redirect to home if logged in, else to login
  } else {
    next(); // Allow navigation
  }
});

export default router;
