//TODO: Add basic input validation
//FIXME: Not entering anything in the input fields throws errors in the console


<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
import { hashPassword } from "@/utils/password";
const toast = useToast();
const onFormSubmit = (event: FormSubmitEvent) => {
    let username = event.states.username;
    let email = event.states.email;
    let password = event.states.password;
    let invalid = false;
    if (!username.value) {
        username.invalid = true;
        username.valid = false;
        username.error = { message: 'Username is required' };
        invalid = true;
        toast.add({severity: 'error', summary: 'Username is required', life: 3000});
        return;
    }
    if (!email.value) {
        email.invalid = true;
        email.valid = false;
        email.error = { message: 'Email is required' };
        invalid = true;
        toast.add({severity: 'error', summary: 'Email is required', life: 3000});
        return;
    }
    if (!password.value) {
        password.invalid = true;
        password.valid = false;
        password.error = { message: 'Password is required' };
        invalid = true;
        toast.add({severity: 'error', summary: 'Password is required', life: 3000});
        return;
    }
    if (invalid) {
        return;
    }
    createAccount(username.value, email.value, password.value, toast);
};
const createAccount = async (username: string, email: string, password: string, toast: any) => {
    let passwordSalt = Array.prototype.map.call(crypto.getRandomValues(new Uint8Array(16)), x=>(('00'+x.toString(16)).slice(-2))).join('');
    let hashedPassword = hashPassword(password, passwordSalt);
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
        toast.add({severity: 'success', summary: 'Account created', life: 3000});
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
        <Form v-slot="$form" @submit="onFormSubmit" class="register-element">
            <div class="register-element vertical-padding">
                <InputText name="username" type="text" placeholder="Username" class="register-input"/>
                <Message v-if="$form.states?.username.valid" severity="error" size="small" variant="simple">{{ $form.states.username.error?.message }}</Message>
            </div>
            <div class="register-element vertical-padding">
                <InputText name="email" type="email" placeholder="Email" class="register-input"/>
                <Message v-if="$form.states?.email.invalid" severity="error" size="small" variant="simple">{{ $form.states.email.error?.message }}</Message>
            </div>
            <div class="register-element vertical-padding">
                <InputText name="password" type="password" placeholder="Password" class="register-input"/>
                <Message v-if="$form.states?.password.invalid" severity="error" size="small" variant="simple">{{ $form.states.password.error?.message }}</Message>
            </div>
            <Button type="submit" severity="secondary" label="Create account" class="register-input"/>
        </Form>
    </main>
</template>

<style scoped>
.vertical-padding {
    padding-top: 0.25rem;
    padding-bottom: 0.25rem;
}
.register-input {
    width: 100%;
}
.register-element {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}
.register-component {
    text-align: center;
    margin: auto;
    width: 60vw;
    max-width: 400px;
}
</style>
