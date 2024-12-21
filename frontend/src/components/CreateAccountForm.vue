<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
import { hashPassword } from "@/utils/password";
import {useRouter} from "vue-router";
const router = useRouter();
const toast = useToast();
const onFormSubmit = (event: FormSubmitEvent) => {
    let username = event.states.username.value;
    let email = event.states.email.value;
    let password = event.states.password.value;
    if (!username || !email || !password) {
      console.log("Empty fields");
      toast.add({severity: 'error', summary: 'Credentials cannot be empty', life: 3000});
      return;
    }
    createAccount(username, email, password, toast);
};
const createAccount = async (username: string, email: string, password: string, toast: any) => {
    let passwordSalt = Array.prototype.map.call(crypto.getRandomValues(new Uint8Array(16)), x=>(('00'+x.toString(16)).slice(-2))).join('');
    let hashedPassword = await hashPassword(password, passwordSalt);
    let response = fetch(`http://localhost:2024/accounts`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email: email,
            passwordHash: hashedPassword,
            passwordSalt: passwordSalt,
            role: "user",
            username: username,
        })
    });
    if ((await response).status === 200) {
        console.log('Account created');
        toast.add({severity: 'success', summary: 'Account created successfully. Redirecting to login page in 3 seconds...', life: 3000});
        await new Promise(f => setTimeout(f, 3000));
        await router.push("/login");
    } else {
        console.log('Account creation failed');
        toast.add({severity: 'error', summary: 'Account creation failed: ' + await (await response).text(), life: 3000});
    }
};
</script>

<template>
    <main class="register-component">
        <Toast />
        <h1 class="vertical-padding">Create an account</h1>
        <Form @submit="onFormSubmit">
            <FormField v-slot="$username" name="username" class="vertical-padding">
                <InputText type="text" placeholder="Username" v-bind="$username.props" class="register-element"/>
                <Message v-if="$username.invalid" severity="error" size="small" variant="simple">{{ $username.error?.message }}</Message>
            </FormField>
            <FormField v-slot="$email" name="email" class="vertical-padding">
                <InputText type="email" placeholder="Email" v-bind="$email.props" class="register-element"/>
                <Message v-if="$email.invalid" severity="error" size="small" variant="simple">{{ $email.error?.message }}</Message>
            </FormField>
            <FormField v-slot="$password" name="password" class="vertical-padding">
                <InputText type="password" placeholder="Password" v-bind="$password.props" class="register-element"/>
                <Message v-if="$password.invalid" severity="error" size="small" variant="simple">{{ $password.error?.message }}</Message>
            </FormField>
            <Button type="submit" severity="secondary" label="Create account" class="register-element"/>
        </Form>
    </main>
</template>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.register-element {
    width: 100%;
}
.register-component {
    text-align: center;
    margin: auto;
    width: 60vw;
    max-width: 400px;
}
</style>
