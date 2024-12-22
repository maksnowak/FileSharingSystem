<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
const router = useRouter();
const toast = useToast();
import { hashPassword } from "@/utils/password";
const onFormSubmit = async (event: FormSubmitEvent) => {
    let username = event.states.username;
    let password = event.states.password;
    let invalid = false;
    if (!username.value) {
        username.invalid = true;
        username.valid = false;
        username.error = { message: 'Username is required' };
        toast.add({severity: 'error', summary: 'Username is required', life: 3000});
        invalid = true;
    }
    if (!password.value) {
        password.invalid = true;
        password.valid = false;
        password.error = { message: 'Password is required' };
        toast.add({severity: 'error', summary: 'Password is required', life: 3000});
        invalid = true;
    }
    console.log(event);
    if (invalid) {
        return;
    }
    await login(username.value, password.value, router, toast);
};
const getSalt = async (username: string) => {
    let response = await fetch(`http://localhost:2024/login/${username}`);
    if (response.status !== 200) { // User does not exist
      console.log("Could not get salt");
      toast.add({severity: 'error', summary: 'Login failed: ' + await (await response).text(), life: 3000});
      return null;
    }
    let data = await response.json();
    return data.passwordSalt;
};
const login = async (username: string, password: string, router: any, toast: any) => {
    let salt = await getSalt(username);
    if (!salt) {
      return; // Wrong username, stop the process
    }
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
      toast.add({severity: 'error', summary: 'Login failed: ' + await (await response).text(), life: 3000});
    }
};
</script>

<template>
    <main class="login-component">
        <Toast />
        <h1 class="vertical-padding">Log in</h1>
        <Form v-slot="$form" @submit="onFormSubmit" class="login-element">
            <div class="login-element vertical-padding">
                <InputText name="username" type="text" placeholder="Username" class="login-input"/>
                <Message v-if="$form.states?.username.invalid" severity="error" size="small" variant="simple">{{ $form.states.username.error?.message }}</Message>
            </div>
            <div class="login-element vertical-padding">
                <InputText name="password" type="password" placeholder="Password" class="login-input"/>
                <Message v-if="$form.states?.password.invalid" severity="error" size="small" variant="simple">{{ $form.states.password.error?.message }}</Message>
            </div>
            <Button type="submit" label="Log in" class="login-input"/>
        </Form>
    </main>
</template>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.login-input {
    width: 100%;
}
.login-element {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}
.login-component {
    text-align: center;
    margin: auto;
    width: 60vw;
    max-width: 400px;
}
</style>
