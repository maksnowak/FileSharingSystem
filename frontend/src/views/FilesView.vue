<script setup lang="ts">
import { ref, onMounted } from 'vue';
import FileComponent from '@/components/FileComponent.vue';
import { ProgressSpinner } from 'primevue';
import AddFile from '@/components/AddFile.vue';

interface FileData {
  id: string;
  file_name: string;
  user_id: string;
  tags: string[];
  has_access: string[];
}

const files = ref<FileData[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const fetchFiles = async () => {
  try {
    const response = await fetch('http://localhost:8080/files');
    if (!response.ok) throw new Error('Failed to fetch files');
    files.value = await response.json();
  } catch (e) {
    error.value = e instanceof Error ? e.message : 'Failed to load files';
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchFiles();
});
</script>

<template>
  <div class="files-container">
    <h2>My Files</h2>
    
    <div v-if="loading" class="loading-container">
      <ProgressSpinner />
    </div>
    
    <div v-else-if="error" class="error-message">
      {{ error }}
    </div>

    <div v-else-if="files.length === 0" class="no-files-message">
      No files found. Why not add one?
      <AddFile />
    </div>
    
    <div v-else class="files-grid">
      <FileComponent
        v-for="file in files"
        :key="file.id"
        :file="file"
      />
    </div>
  </div>
</template>

<style scoped>
.files-container {
  padding: 2rem;
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
</style>