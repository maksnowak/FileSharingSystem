<script setup lang="ts">
import { Button } from "primevue";
import { userStore } from "@/user"; // Import your userStore
import { useRouter } from "vue-router";
import { computed } from "vue";
const router = useRouter();
const store = userStore(); // Get the user store

// Check if the user is logged in
const user = computed(() => store.getUser);

// Logout action
const onLogout = () => {
  store.clearUser(); // Clear the user session
  router.push("/login"); // Redirect to the login page
};
</script>

<template>
  <nav>
    <ul>
      <li>
        <Button
          class="p-button-text nav-button"
          @click="router.push('/home')"
        >
          <span>My files</span>
        </Button>
      </li>
      <li>
        <Button
          class="p-button-text nav-button"
          @click="router.push('/shared')"
        >
          <span>Shared files</span>
        </Button>
      </li>
      <li>
        <Button
          class="p-button-text nav-button"
          @click="router.push('/account')"
        >
          <span>Account</span>
        </Button>
      </li>
    </ul>
    <div class="user-controls">
      <span class="username">Hello, {{ user ? user.username : "Guest" }}! 👋</span>
      <Button
        v-if="user"
        label="Log Out"
        class="logout-button p-button-text"
        @click="onLogout"
      />
    </div>
  </nav>
  <RouterView />
</template>

<style scoped>
nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background: linear-gradient(to right, #0F172A, #334155);
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

ul {
  list-style: none;
  display: flex;
  gap: 1rem;
  margin: 0;
  padding: 0;
}

li {
  padding: 0.5rem 0;
}

.logout-button {
  color: white !important;
  border: 1px solid rgba(255,255,255,0.3) !important;
  display: flex !important;
  align-items: center !important;
  gap: 0.5rem !important;
}

.logout-button:hover {
  background: rgba(255,255,255,0.1) !important;
}

.user-controls {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.username {
  color: white;
  font-size: 1.1rem;
  opacity: 0.9;
  white-space: nowrap; /* Prevent username from wrapping */
  margin: 0; /* Remove margin-right since we're using gap */
}

@media (max-width: 768px) {
  nav {
    padding: 1rem;
  }
  
  ul {
    gap: 0.5rem;
  }
}
</style>
