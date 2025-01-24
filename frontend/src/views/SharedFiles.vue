<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import Dialog from 'primevue/dialog';
import Password from 'primevue/password';
import Button from 'primevue/button';
import SharedFileComponent from '@/components/SharedFileComponent.vue';
import { ProgressSpinner } from 'primevue';
import { userStore } from "@/user";
import { useToast } from 'primevue';

interface SharedFileData {
  id: string;
  file_name: string;
  user_id: string;
  tags: string[];
  has_access: string[];
  passwordHash: string;
  iv: Uint8Array;
  encryptedContent: Uint8Array;
}

const files = ref<SharedFileData[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);
const selectedFile = ref<SharedFileData | null>(null);
const showPasswordDialog = ref(false);
const showFileInfoDialog = ref(false);
const password = ref('');
const searchQuery = ref('');
const toast = useToast();

const store = userStore();
const userId = computed(() => store.getUser?.id);

const filteredFiles = computed(() => {
  if (!searchQuery.value) return files.value;
  return files.value.filter((file: SharedFileData) =>
    file.file_name.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});

const fetchSharedFiles = async () => {
  try {
    loading.value = true;
    const response = await fetch('http://localhost:8080/files');
    if (!response.ok) throw new Error('Failed to fetch files');

    const allFiles = await response.json();
    console.log('API Response:', allFiles); // Debug log

    // Filter files where current user has access
    files.value = allFiles.filter((file: SharedFileData) =>
      file.has_access.includes(userId.value)
    );

  } catch (e) {
    console.error('Fetch error:', e); // Debug log
    error.value = e instanceof Error ? e.message : 'Failed to load files';
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchSharedFiles();
});

const handleFileClick = (file: SharedFileData) => {
  selectedFile.value = file;
  showPasswordDialog.value = true;
  password.value = '';
  error.value = null;
};

const hashFilePassword = async (password: string) => {
  let hashBuffer = await crypto.subtle.digest('SHA-256', new TextEncoder().encode(password));
  return Array.prototype.map.call(new Uint8Array(hashBuffer), x=>(('00'+x.toString(16)).slice(-2))).join('');
};

const handlePasswordSubmit = async () => {
  if (!selectedFile.value) return;

  loading.value = true;
  error.value = '';

  try {
    const hashedPassword = await hashFilePassword(password.value);
    if (hashedPassword !== selectedFile.value.passwordHash) {
      throw new Error('Invalid password');
    }

    showPasswordDialog.value = false;
    showFileInfoDialog.value = true;
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Invalid password';
  } finally {
    loading.value = false;
    password.value = '';
  }
};

const handleDownload = async () => {
  if (!selectedFile.value?.id) return;

  try {
    const response = await fetch(`http://localhost:8080/files/${selectedFile.value.id}`);
    if (response.ok) {
      toast.add({severity: 'success', summary: 'File downloaded successfully', life: 3000});
      showFileInfoDialog.value = false;
    } else {
      toast.add({severity: 'error', summary: 'Failed to download file', life: 3000});
    }
  } catch (error) {
    toast.add({severity: 'error', summary: 'Error downloading file', life: 3000});
  }
};
</script>

<template>
  <div class="files-container">
    <h2>Files Shared With Me</h2>

    <input
      v-model="searchQuery"
      type="text"
      placeholder="Search files..."
      class="search-input"
    />

    <div v-if="loading" class="loading-container">
      <ProgressSpinner />
    </div>

    <div v-else-if="error" class="error-message">
      {{ error }}
    </div>

    <div v-else class="files-grid">
      <SharedFileComponent
        v-for="file in filteredFiles"
        :key="file.id"
        :file="file"
        @fileSelect="handleFileClick"
      />
    </div>

    <Dialog
      v-model:visible="showPasswordDialog"
      modal
      :header="selectedFile?.file_name"
      :style="{ width: '450px' }"
    >
      <div class="password-dialog">
        <Password
          v-model="password"
          placeholder="Enter file password"
          :feedback="false"
          @keyup.enter="handlePasswordSubmit"
        />
        <small class="error-text">{{ error }}</small>
      </div>
      <template #footer>
        <Button
          label="Cancel"
          class="p-button-text"
          @click="showPasswordDialog = false"
        />
        <Button
          label="Access File"
          :loading="loading"
          @click="handlePasswordSubmit"
        />
      </template>
    </Dialog>

    <Dialog
      v-model:visible="showFileInfoDialog"
      modal
      :header="selectedFile?.file_name"
      :style="{ width: '500px' }"
    >
      <div class="file-info" v-if="selectedFile">
        <div class="info-row">
          <span class="label">Name:</span>
          <span>{{ selectedFile.file_name }}</span>
        </div>
        <div class="info-row">
          <span class="label">Owner:</span>
          <span>{{ selectedFile.user_id }}</span>
        </div>
        <div class="info-row">
          <span class="label">Tags:</span>
          <span>{{ selectedFile.tags.join(', ') }}</span>
        </div>
        <div class="info-row">
          <span class="label">Access:</span>
          <span>{{ selectedFile.has_access.join(', ') }}</span>
        </div>
        <div class="info-row">
          <span class="label">Status:</span>
          <span class="status">Encrypted</span>
        </div>
      </div>
      <template #footer>
        <Button
          label="Download Encrypted"
          icon="pi pi-download"
          @click="handleDownload"
          class="p-button-primary"
        />
        <Button
          label="Close"
          class="p-button-text"
          @click="showFileInfoDialog = false"
        />
      </template>
    </Dialog>
  </div>
</template>

<style scoped>
.files-container {
  padding: 2rem;
}

.search-input {
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1.5rem;
  background: #1E293B;
  border: 1px solid #334155;
  border-radius: 6px;
  color: #F8FAFC;
}

.loading-container {
  display: flex;
  justify-content: center;
  padding: 2rem;
}

.error-message {
  color: red;
  text-align: center;
  padding: 2rem;
}

.no-files-message {
  text-align: center;
  padding: 2rem;
}

.files-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
  padding: 1rem 0;
}

.password-dialog {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1rem 0;
}

.error-text {
  color: #EF4444;
  height: 20px;
}

.file-info {
  padding: 1rem;
  background: #1E293B;
  border-radius: 6px;
}

.info-row {
  display: flex;
  padding: 0.5rem 0;
  border-bottom: 1px solid #334155;
}

.info-row:last-child {
  border-bottom: none;
}

.label {
  font-weight: 600;
  width: 100px;
  color: #94A3B8;
}

.info-row .status {
  color: #10B981;
  font-weight: 500;
}
</style>