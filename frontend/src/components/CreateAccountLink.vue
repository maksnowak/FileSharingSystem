<script setup lang="ts">
import { Button } from 'primevue';
import { useRouter } from 'vue-router';
const router = useRouter();
const goToCreateAccount = async () => {
    await router.push('/create-account');
}
</script>

<template>
    <main class="link-component">
        <p class="link-element vertical-padding">If you don't have an account, click the button below:</p>
        <Button class="link-element" severity="secondary" label="Create account" @click="goToCreateAccount" />
    </main>
</template>

<script lang="ts">
</script>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.link-component {
    text-align: center;
    margin: auto;
    max-width: 400px;
}
.link-element {
    width: 100%;
}
</style>