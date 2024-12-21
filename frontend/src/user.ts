import { defineStore } from 'pinia';

export const userStore = defineStore('user', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null, // Retrieve from localStorage if it exists
  }),
  actions: {
    setUser(user) {
      this.user = user;
      // Store user in localStorage for persistence
      localStorage.setItem('user', JSON.stringify(user));
    },
    clearUser() {
      this.user = null;
      localStorage.removeItem('user'); // Remove user from localStorage on logout
    },
  },
  getters: {
    getUser: (state) => state.user,
  },
});
