<script setup lang="ts">
import Message from 'primevue/message';
import AccountDashboard from "@/components/AccountDashboard.vue";
</script>

<template>
    <main>
        <Message severity="success">Logged in successfully</Message>
        <AccountDashboard />
    </main>
</template>
