//TODO: Add basic input validation
//FIXME: Not entering anything in the input fields throws errors in the console


<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
const toast = useToast();
const onFormSubmit = (event: FormSubmitEvent) => {
    let username = event.states.username.value;
    let email = event.states.email.value;
    let password = event.states.password.value;
    createAccount(username, email, password, toast);
};
const createAccount = async (username: string, email: string, password: string, toast: any) => {
    let passwordSalt = Array.prototype.map.call(crypto.getRandomValues(new Uint8Array(16)), x=>(('00'+x.toString(16)).slice(-2))).join('');
    let saltedPassword = password + passwordSalt;
    let hashBuffer = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(saltedPassword));
    let hashedPassword = Array.prototype.map.call(new Uint8Array(hashBuffer), x=>(('00'+x.toString(16)).slice(-2))).join('');
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
        toast.add({severity: 'error', summary: 'Account creation failed', life: 3000});
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