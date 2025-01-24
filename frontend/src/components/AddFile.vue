<script setup lang="ts">
import { ref } from 'vue';
import { Form, FormField } from '@primevue/forms';
import { InputText, Button, Divider } from 'primevue';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import { hashPassword } from "@/utils/password";
import { userStore } from "@/user";

const showPopup = ref(false);
const selectedFile = ref<File | null>(null);
const password = ref('');
const tags = ref<string[]>([]);
const newTag = ref('');
const toast = useToast();

const handleFileSelection = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files) {
    selectedFile.value = target.files[0];
  }
};

const addTag = () => {
  if (newTag.value.trim()) {
    tags.value.push(newTag.value.trim());
    newTag.value = '';
  }
};

const removeTag = (index: number) => {
  tags.value.splice(index, 1);
};

const encryptFile = async (file: File, passwordHash: string): Promise<Blob> => {
  const formData = new FormData();
  formData.append('file', file);
  formData.append('password-hash', passwordHash);
  
  const response = await fetch('http://localhost:7780/encrypt', {
    method: 'POST',
    headers: {
      accept: 'application/octet-stream',
    },
    body: formData,
  });
  
  if (!response.ok) {
    toast.add({severity: 'error', summary: 'Error', detail: 'Could not encrypt file.', life: 3000});
    throw new Error('Encryption failed');
  }
  
  return response.blob();
};

const uploadFile = async () => {
  if (!selectedFile.value) {
    toast.add({ severity: 'warn', summary: 'No File Selected', detail: 'Please select a file to upload.', life: 3000 });
    return;
  }

  try {
    let fileData = selectedFile.value;
    const fileName = selectedFile.value.name;
    let passwordSalt: string | null = null;
    let passwordHash: string | null = null;

    if (password.value) {
      passwordSalt = Array.prototype.map.call(
        crypto.getRandomValues(new Uint8Array(16)), 
        x => ('00' + x.toString(16)).slice(-2)
      ).join('');
      passwordHash = await hashPassword(password.value, passwordSalt);
      fileData = await encryptFile(selectedFile.value, passwordHash);
    }

    const fileID = Array.prototype.map.call(
      crypto.getRandomValues(new Uint8Array(16)), 
      x => ('00' + x.toString(16)).slice(-2)
    ).join('');
    
    const user = userStore().getUser;
    const uid = user.id;

    let response = await fetch('http://localhost:8080/files', {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        data: fileData,
        file_name: fileName,
        has_access: [],
        id: fileID,
        tags: tags.value,
        user_id: uid,
        password_salt: passwordSalt,
        password_hash: passwordHash
      })
    });

    if (!response.ok) {
      toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to upload file.', life: 3000 });
      return;
    }

    toast.add({ severity: 'success', summary: 'Success', detail: 'File uploaded successfully!', life: 3000 });

    // Update user's owned files
    const updatedUser = { ...user };
    updatedUser.ownedFiles.push(fileID);
    userStore().setUser(updatedUser);

    response = await fetch(`http://localhost:2024/accounts/${uid}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: user.email,
        passwordHash: user.passwordHash,
        passwordSalt: user.passwordSalt,
        sharedFiles: user.sharedFiles,
        ownedFiles: updatedUser.ownedFiles
      })
    });

    if (!response.ok) {
      toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to update user.', life: 3000 });
      console.log("Update failed: " + await response.text());
      return;
    }

    console.log("User updated successfully");
    resetForm();

  } catch (error) {
    console.error('Error uploading file:', error);
    toast.add({ severity: 'error', summary: 'Error', detail: 'An error occurred while uploading.', life: 3000 });
  }
};

const resetForm = () => {
  selectedFile.value = null;
  password.value = '';
  tags.value = [];
  newTag.value = '';
};

const closePopup = () => {
  showPopup.value = false;
  resetForm();
};
</script>

<template>
  <div>
  <Button
    type="button"
    label="Open Upload Form"
    icon="pi pi-upload"
    class="p-button-primary"
    @click="showPopup = true"
  />
  </div>
  <div v-if="showPopup" class="popup-overlay">
    <div class="popup-content">
      <h3>Upload File</h3>

      <Form class="form">
        <FormField label="Choose a file:" required>
          <input type="file" @change="handleFileSelection" class="file-input" />
        </FormField>

        <FormField label="Password (optional):">
          <InputText v-model="password" type="password" placeholder="Enter password to encrypt file" />
        </FormField>

        <FormField label="Tags:">
          <div class="tags-container">
            <div class="tag" v-for="(tag, index) in tags" :key="index">
              {{ tag }} <button type="button" class="remove-tag" @click="removeTag(index)">&times;</button>
            </div>
          </div>
          <InputText
            v-model="newTag"
            @keyup.enter="addTag"
            placeholder="Press Enter to add a tag"
          />
        </FormField>

        <Divider />
        <div class="form-buttons">
          <Button
            type="button"
            label="Upload"
            icon="pi pi-upload"
            class="p-button-primary"
            @click="uploadFile"
          />
          <Button type="button" label="Close" icon="pi pi-times" class="p-button-secondary" @click="closePopup" />
        </div>
      </Form>
    </div>
    <Toast />
  </div>
</template>

<style scoped>
.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup-content {
  background: #222222;
  color: white;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.file-input {
  width: 100%;
  margin-top: 8px;
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  margin-bottom: 5px;
}

.tag {
  background: #e3f2fd;
  color: #0d47a1;
  padding: 5px 10px;
  border-radius: 15px;
  display: flex;
  align-items: center;
}

.remove-tag {
  background: none;
  border: none;
  margin-left: 5px;
  cursor: pointer;
  color: #f44336;
}

.form-buttons {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
</style>