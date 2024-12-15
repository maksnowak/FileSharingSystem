<script setup lang="ts">
import LoginForm from '../components/LoginForm.vue'
import CreateAccountLink from '@/components/CreateAccountLink.vue';
</script>

<template>
    <main class="login-page">
        <LoginForm />
        <CreateAccountLink />
    </main>
</template>

<style scoped>
.login-page {
  justify-content: center;
  margin: auto;
  max-width: 1200px;
  height: 100vh;
}
</style>