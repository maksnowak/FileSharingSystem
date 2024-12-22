<script setup lang="ts">
import {Button} from "primevue";
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
  <div>
    <Button
      v-if="user"
      label="Log Out"
      icon="pi pi-sign-out"
      class="logout-button"
      @click="onLogout"
    />
    <RouterView />
  </div>
</template>

<style scoped>
.logout-button {
  position: absolute;
  top: 5px;
  right: 10px;
  height: 30px;
  text-align: center;
  margin: auto;
}
</style>
