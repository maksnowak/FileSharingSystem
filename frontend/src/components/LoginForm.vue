<script setup lang="ts">
import { Form, FormField } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
</script>

<template>
    <main>
        <h1>Login</h1>
        <Form class="flex flex-col gap-4 w-full">
            <FormField v-slot="$field" name="username" class="flex flex-col gap-1 w-full">
                <InputText type="text" placeholder="Username" v-bind="$field.props" />
                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}</Message>
            </FormField> 
            <FormField v-slot="$field" name="password" class="flex flex-col gap-1 w-full">
                <InputText type="password" placeholder="Password" v-bind="$field.props" />
                <Message v-if="$field.invalid" severity="error" size="small" variant="simple">{{ $field.error?.message }}</Message>
            </FormField>
            <Button type="submit" severity="secondary" label="Login" />
        </Form>
    </main>
</template>