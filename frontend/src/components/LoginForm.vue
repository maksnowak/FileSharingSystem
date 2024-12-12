<script setup lang="ts">
import { Form, FormField } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
</script>

<template>
    <main class="login-component">
        <h1>Login</h1>
        <Form>
            <FormField v-slot="$field" name="username" class="vertical-padding">
                <InputText type="text" placeholder="Username" v-bind="$field.props" class="login-element"/>
                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}</Message>
            </FormField> 
            <FormField v-slot="$field" name="password" class="vertical-padding">
                <InputText type="password" placeholder="Password" v-bind="$field.props" class="login-element"/>
                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}</Message>
            </FormField>
            <Button type="submit" severity="secondary" label="Login" class="login-element"/>
        </Form>
    </main>
</template>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.login-element {
    width: 100%;
}
.login-component {
    justify-content: center;
    align-items: center;
    max-width: 400px;
}
</style>