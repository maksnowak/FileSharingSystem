<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
const router = useRouter();
const toast = useToast();
import { hashPassword } from "@/utils/password";
import {userStore} from "@/user";
const onFormSubmit = async (event: FormSubmitEvent) => {
    let username = event.states.username.value;
    let password = event.states.password.value;
    if (!username || !password) {
      console.log("Empty credentials");
      toast.add({severity: 'error', summary: 'Credentials cannot be empty', life: 3000});
      return;
    }
    await login(username, password, router, toast);
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
        const user = await response.json();
        await userStore().setUser(user);
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
