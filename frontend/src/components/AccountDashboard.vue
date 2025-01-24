<script setup lang="ts">
import { Form, FormField, type FormSubmitEvent } from '@primevue/forms';
import { InputText, Message, Button, Divider } from 'primevue';
import { useToast } from 'primevue/usetoast';
import Toast from 'primevue/toast';
import { hashPassword } from "@/utils/password";
import { useRouter } from "vue-router";
import { userStore } from "@/user";
import { computed } from "vue";
const router = useRouter();
const toast = useToast();
const store = userStore();
const user = computed(() => store.getUser);

const updateEmail = async (event: FormSubmitEvent) => {
  let email = event.states.email.value;
  if (!email) {
    console.log("Empty email");
    toast.add({severity: 'error', summary: 'Credentials cannot be empty', life: 3000});
    return;
  }
  if (email == user._value.email) {
    console.log("Email unchanged");
    toast.add({severity: 'error', summary: 'New email cannot be the same as the old one.', life: 3000});
    return;
  }
  const result = await updateAccount(email, user._value.passwordHash);
  if (!result) {
    console.log("Email updated");
    const updatedUser = { ...user.value, email };  // Make a copy and update the email
    store.setUser(updatedUser);
    toast.add({severity: 'success', summary: 'Email updated successfully', life: 3000});
  } else {
    console.log('Account creation failed');
    toast.add({severity: 'error', summary: 'Email update failed: ' + result, life: 3000});
  }
}

const changePassword = async (event: FormSubmitEvent) => {
  let old_password = event.states.oldpassword.value;
  let new_password = event.states.newpassword.value;
  if (!old_password || !new_password) {
    console.log("Password fields missing");
    toast.add({severity: 'error', summary: 'Password fields must not be empty', life: 3000});
    return;
  }
  // CHECK OLD PASSWORD
  const oldHash = await hashPassword(old_password, user._value.passwordSalt);
  if (oldHash !== user._value.passwordHash) {
    console.log("Wrong password");
    toast.add({severity: 'error', summary: 'Old password is incorrect', life: 3000});
    return;
  }
  const passwordHash = await hashPassword(new_password, user._value.passwordSalt);
  if (oldHash == passwordHash) {
    console.log("Passwords the same");
    toast.add({severity: 'error', summary: 'New password cannot be the same as old password', life: 3000});
    return;
  }
  const result = await updateAccount(user._value.email, passwordHash);
  if (!result) {
    console.log("Password updated");
    const updatedUser = { ...user.value, passwordHash };  // Make a copy and update the email
    store.setUser(updatedUser);
    toast.add({severity: 'success', summary: 'Password updated successfully', life: 3000});
  } else {
    console.log('Account creation failed');
    toast.add({severity: 'error', summary: 'Password update failed: ' + result, life: 3000});
  }
}

const updateAccount = async (email: string, hashedPassword: string) => {
  console.log(user._value.id);
  let response = fetch(`http://localhost:2024/accounts/`+user._value.id, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: email,
      passwordHash: hashedPassword,
      passwordSalt: user._value.passwordSalt,
      sharedFiles: user._value.sharedFiles,
      ownedFiles: user._value.ownedFiles
    })
  })
  if ((await response).status === 200) {
    return null;
  } else {
    return await (await response).text();
  }
}

const onDelete = async (event: FormSubmitEvent) => {
  const password = event.states.password.value;
  if (!password) {
    console.log("Password field missing");
    toast.add({severity: 'error', summary: 'You must type in your password before deleting your account', life: 3000});
    return;
  }
  const passwordHash = await hashPassword(password, user._value.passwordSalt);
  if (passwordHash !== user._value.passwordHash) {
    console.log("Wrong password");
    toast.add({severity: 'error', summary: 'Password incorrect', life: 3000});
    return;
  }
  console.log("Starting deletion procedure...");
  // TODO: Remove this when file-transfer starts working
  if (1 == 1) {
    console.log("Not yet implemented");
    toast.add({severity: 'error', summary: 'Operation not yet implemented', life: 3000});
    return;
  }
  // Delete user files first
  let filesToDelete = user._value.ownedFiles;
  for (const file of filesToDelete) {
    console.log("Deleting "+file);
    let response = await fetch ("http://localhost:8080/"+file, {
      method: "DELETE"
    });
    if (!response.ok) {
      this.toast.add({severity: 'error', summary: 'Could not delete files.', life: 3000});
      return;
    }
  }

  let response = fetch(`http://localhost:2024/accounts/`+user._value.id, {method: 'DELETE'})
  if ((await response).status === 200) {
    await store.clearUser();
    console.log("Account deleted successfully");
    toast.add({severity: 'success', summary: 'Account deleted successfully. Redirecting to login page in 3 seconds...', life: 3000});
    await new Promise(f => setTimeout(f, 3000));
    await router.push("/login");
  } else {
    console.log('Account deletion failed');
    toast.add({severity: 'error', summary: 'Account deletion failed: ' + await (await response).text(), life: 3000});
  }
}

</script>

<template>
  <main class="register-component">
    <Toast />
    <h1 class="vertical-padding">Your account dashboard</h1>
    <div class="left">
      <p><span class="bold">Username:</span> {{ user.username }}</p>
      <p><span class="bold">Email:</span> {{ user.email }}</p>
      <p><span class="bold">Date of creation:</span> {{ user.createdAt }}</p>
    </div>
    <Divider type="solid" />
    <h2>Change your email</h2>
    <Form @submit="updateEmail">
      <FormField v-slot="$email" name="email" class="vertical-padding">
        <InputText type="email" placeholder="Email" v-bind="$email.props" class="register-element"/>
        <Message v-if="$email.invalid" severity="error" size="small" variant="simple">{{ $email.error?.message }}</Message>
      </FormField>
      <Button type="submit" label="Update email" severity="secondary" class="register-element" />
    </Form>
    <Divider type="solid" />
    <h2>Change your password</h2>
    <Form @submit="changePassword">
      <FormField v-slot="$oldpassword" name="oldpassword" class="vertical-padding">
        <InputText type="password" placeholder="Old password" v-bind="$oldpassword.props" class="register-element"/>
        <Message v-if="$oldpassword.invalid" severity="error" size="small" variant="simple">{{ $oldpassword.error?.message }}</Message>
      </FormField>
      <FormField v-slot="$newpassword" name="newpassword" class="vertical-padding">
        <InputText type="password" placeholder="New password" v-bind="$newpassword.props" class="register-element"/>
        <Message v-if="$newpassword.invalid" severity="error" size="small" variant="simple">{{ $newpassword.error?.message }}</Message>
      </FormField>
      <Button type="submit" severity="secondary" label="Change your password" class="register-element"/>
    </Form>
    <Divider type="solid" />
    <h2>Delete your account</h2>
    <Form @submit="onDelete">
      <FormField v-slot="$password" name="password" class="vertical-padding">
        <InputText type="password" placeholder="Password" v-bind="$password.props" class="register-element"/>
        <Message v-if="$password.invalid" severity="error" size="small" variant="simple">{{ $password.error?.message }}</Message>
      </FormField>
      <Button type="submit" label="Delete account" class="register-element red"/>
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

.left {
  text-align: left;
}

.bold {
  font-weight: bolder;
}

.red {
  background: #690000;
  border: #690000;
  color: white;
}
</style>
