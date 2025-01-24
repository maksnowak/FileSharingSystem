<script setup lang="ts">
import { Card } from 'primevue';
import { Tag } from 'primevue';

interface SharedFileProps {
  id?: string;
  file_name: string;
  user_id: string;
  tags: string[];
  has_access: string[];
}

const props = defineProps<{
  file: SharedFileProps
}>();

const emit = defineEmits(['fileSelect']);
</script>

<template>
  <Card class="file-card" @click="emit('fileSelect', file)">
    <template #title>
      <div class="file-header">
        <span class="file-name">{{ file.file_name }}</span>
      </div>
    </template>
    <template #content>
      <div class="tags-container">
        <Tag v-for="tag in file.tags" :key="tag" :value="tag" severity="info" class="tag" />
      </div>
      <div class="owner">Shared by: {{ file.user_id }}</div>
    </template>
  </Card>
</template>

<style scoped>
.file-card {
  margin: 0.5rem;
  max-width: 300px;
  cursor: pointer;
  transition: transform 0.2s;
}

.file-card:hover {
  transform: translateY(-2px);
}

.file-header {
  display: flex;
  align-items: center;
}

.file-name {
  font-weight: bold;
  font-size: 1.1rem;
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin: 1rem 0;
}

.owner {
  color: #94A3B8;
  font-size: 0.9rem;
}
</style>