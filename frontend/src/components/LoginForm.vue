//TODO: Add basic input validation
//FIXME: Not entering anything in the input fields throws errors in the console


<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
const router = useRouter();
const toast = useToast();
const onFormSubmit = async (event: FormSubmitEvent) => {
    let username = event.states.username.value;
    let password = event.states.password.value;
    await login(username, password, router, toast);
};
const getSalt = async (username: string) => {
    let response = await fetch(`http://localhost:2024/login/${username}`);
    let data = await response.json();
    return data.passwordSalt;
};
const hashPassword = async (password: string, salt: string) => {
    let saltedPassword = password + salt;
    let hashBuffer = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(saltedPassword));
    let hashedPassword = Array.prototype.map.call(new Uint8Array(hashBuffer), x=>(('00'+x.toString(16)).slice(-2))).join('');
    return hashedPassword;
}
const login = async (username: string, password: string, router: any, toast: any) => {
    let salt = await getSalt(username);
    let hashedPassword = await hashPassword(password, salt);
    let response = await fetch(`http://localhost:2024/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            passwordHash: hashedPassword,
            username: username,
        })
    });
    if ((await response).status === 200) {
        console.log('Logged in');
        router.push('/home');
    } else {
        console.log('Login failed');
        toast.add({severity: 'error', summary: 'Login failed', life: 3000});
    }
};
</script>

<template>
    <main class="login-component">
        <Toast />
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
            <Button type="submit" label="Log in" class="login-element"/>
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
    text-align: center;
    margin: auto;
    width: 60vw;
    max-width: 400px;
}
</style>