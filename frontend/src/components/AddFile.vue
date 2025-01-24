<template>
  <div v-if="showPopup" class="popup-overlay">
    <div class="popup-content">
      <h3>Upload File</h3>

      <Form class="form">
        <!-- File Input -->
        <FormField label="Choose a file:" required>
          <input type="file" @change="handleFileSelection" class="file-input" />
        </FormField>

        <!-- Password Input -->
        <FormField label="Password (optional):">
          <InputText v-model="password" type="password" placeholder="Enter password to encrypt file" />
        </FormField>

        <!-- Tags Input -->
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

        <!-- Buttons -->
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

<script>
import { Form, FormField } from '@primevue/forms';
import { InputText, Button, Divider } from 'primevue';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import {hashPassword} from "@/utils/password.ts";
import {userStore} from "@/user";

export default {
  components: { Form, FormField, InputText, Button, Divider, Toast },
  data() {
    return {
      showPopup: false,
      selectedFile: null,
      password: '',
      tags: [],
      newTag: '',
      toast: useToast(),
    };
  },
  methods: {
    handleFileSelection(event) {
      this.selectedFile = event.target.files[0];
    },
    addTag() {
      if (this.newTag.trim()) {
        this.tags.push(this.newTag.trim());
        this.newTag = '';
      }
    },
    removeTag(index) {
      this.tags.splice(index, 1);
    },
    async uploadFile() {
      if (!this.selectedFile) {
        this.toast.add({ severity: 'warn', summary: 'No File Selected', detail: 'Please select a file to upload.', life: 3000 });
        return;
      }

      try {
        let fileData = this.selectedFile;
        let fileName = this.selectedFile.name;
        let passwordSalt = null;
        let passwordHash = null;
        // Encrypt the file if password is provided
        if (this.password) {
          passwordSalt = Array.prototype.map.call(crypto.getRandomValues(new Uint8Array(16)), x=>(('00'+x.toString(16)).slice(-2))).join('');
          passwordHash = hashPassword(this.password, passwordSalt);
          fileData = await this.encryptFile(this.selectedFile, passwordHash);
        }

        // Upload file
        console.log('Uploading file:', fileData);
        console.log('Tags:', this.tags);
        const response = await fetch('http://localhost:8080/files', {
          method: "POST",
          // mode: 'no-cors',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            "data": fileData,
            "file_name": fileName,
            "has_access": [],
            "id": Array.prototype.map.call(crypto.getRandomValues(new Uint8Array(16)), x=>(('00'+x.toString(16)).slice(-2))).join(''),
            "tags": this.tags,
            "user_id": userStore().getUser.id,
            "password_salt": passwordSalt,
            "passwod_hash": passwordHash
          })
        });
        if (!(await response).ok) {
          this.toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to upload file.', life: 3000 });
        } else {
          this.toast.add({ severity: 'success', summary: 'Success', detail: 'File uploaded successfully!', life: 3000 });
        }
        // Reset the form
        this.resetForm();
      } catch (error) {
        console.error('Error uploading file:', error);
        this.toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to upload file.', life: 3000 });
      }
    },
    async encryptFile(file, passwordHash) {
      const formData = new FormData();
      formData.append('file',file);
      formData.append('password-hash',passwordHash);
      let response = await fetch('http://localhost:7780/encrypt', {
        method: 'POST',
        headers: {
          accept: 'application/octet-stream',
        },
        body: formData,
      });
      if (!(await response).ok) {
        this.toast.add({severity: 'error', summary: 'Error', detail: 'Could not encrypt file.', life: 3000})
      }
      return response.blob();
    },
    resetForm() {
      this.selectedFile = null;
      this.password = '';
      this.tags = [];
      this.newTag = '';
    },
    closePopup() {
      this.showPopup = false;
    },
  },
};
</script>

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
