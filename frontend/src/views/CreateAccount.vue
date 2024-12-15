<script setup lang="ts">
import CreateAccountForm from '@/components/CreateAccountForm.vue';
</script>

<template>
    <main class="login-page">
        <CreateAccountForm />
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