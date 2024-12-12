<script setup lang="ts">
import LoginForm from '../components/LoginForm.vue'
</script>

<template>
    <main class="login-page">
        <LoginForm />
    </main>
</template>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  text-align: center;
  align-items: center;
  height: 100vh;
}
</style>