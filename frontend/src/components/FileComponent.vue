<script setup lang="ts">
import { Button } from 'primevue';
import { Card } from 'primevue';
import { Tag } from 'primevue';
import { useToast } from 'primevue';
import Toast from 'primevue/toast';

interface FileProps {
  id?: string;
  file_name: string;
  user_id: string;
  tags: string[];
  has_access: string[];
}

const props = defineProps<{
  file: FileProps
}>();

const toast = useToast();

const onDownload = async () => {
  try {
    const response = await fetch(`http://localhost:8080/files/${props.file.id}`);
    if (response.ok) {
      toast.add({severity: 'success', summary: 'File downloaded successfully', life: 3000});
    } else {
      toast.add({severity: 'error', summary: 'Failed to download file', life: 3000});
    }
  } catch (error) {
    toast.add({severity: 'error', summary: 'Error downloading file', life: 3000});
  }
};

const onDelete = async () => {
  try {
    const response = await fetch(`http://localhost:8080/files/${props.file.id}`, {
      method: 'DELETE'
    });
    if (response.ok) {
      toast.add({severity: 'success', summary: 'File deleted successfully', life: 3000});
    } else {
      toast.add({severity: 'error', summary: 'Failed to delete file', life: 3000});
    }
  } catch (error) {
    toast.add({severity: 'error', summary: 'Error deleting file', life: 3000});
  }
};
</script>

<template>
  <Card class="file-card">
    <template #title>
      <div class="file-header">
        <span class="file-name">{{ file.file_name }}</span>
      </div>
    </template>
    <template #content>
      <div class="tags-container">
        <Tag v-for="tag in file.tags" :key="tag" :value="tag" severity="info" class="tag" />
      </div>
      <div class="actions">
        <Button label="Download" severity="secondary" @click="onDownload" />
        <Button label="Delete" severity="danger" @click="onDelete" />
      </div>
    </template>
  </Card>
  <Toast />
</template>

<style scoped>
.file-card {
  margin: 0.5rem;
  max-width: 300px;
}

.file-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.file-name {
  font-weight: bold;
  font-size: 1.1rem;
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.tag {
  font-size: 0.8rem;
}

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
}
</style>