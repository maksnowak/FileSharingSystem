<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
</script>

<template>
    <main class="login-component">
        <h1 class="vertical-padding">Log in</h1>
        <Form @submit="onFormSubmit">
            <FormField v-slot="$username" name="username" class="vertical-padding">
                <InputText type="text" placeholder="Username" v-bind="$username.props" class="login-element"/>
                <Message v-if="$username.invalid" severity="error" size="small" variant="simple">{{ $username.error?.message }}</Message>
            </FormField> 
            <FormField v-slot="$password" name="password" class="vertical-padding">
                <InputText type="password" placeholder="Password" v-bind="$password.props" class="login-element"/>
                <Message v-if="$password.invalid" severity="error" size="small" variant="simple">{{ $password.error?.message }}</Message>
            </FormField>
            <Button type="submit" severity="secondary" label="Log in" class="login-element"/>
        </Form>
    </main>
</template>

<script lang="ts">
const onFormSubmit = (event: FormSubmitEvent) => {
    let username = event.states.username.value;
    let password = event.states.password.value;
    console.log(`Username: ${username}, Password: ${password}`);
};
const getSalt = async (username: string) => {
    let response = await fetch(`http://localhost:2024/api/login/${username}`);
    let data = await response.json();
    return data.passwordSalt;
};
const hashPassword = (password: string, salt: string) => {
    /* placeholder */
}
const login = async (username: string, password: string) => {
    let salt = await getSalt(username);
    let hashedPassword = hashPassword(password, salt);
    let response = await fetch(`http://localhost:2024/api/login/${username}/${hashedPassword}`);
    let data = await response.json();
    return data;
};
</script>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.login-element {
    width: 100%;
}
.login-component {
    text-align: center;
    margin: auto;
    width: 60vw;
    max-width: 400px;
}
</style>