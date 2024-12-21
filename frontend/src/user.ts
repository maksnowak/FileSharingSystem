import { defineStore } from 'pinia';

export const userStore = defineStore('user', {
  state: () => ({
    user: null, // Global user instance
  }),
  actions: {
    setUser(user) {
      this.user = user;
    },
  },
  getters: {
    getUser: (state) => state.user,
  },
});
