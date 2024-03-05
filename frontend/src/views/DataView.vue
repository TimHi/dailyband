<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

console.log('aal')
const albumStore = useAlbumStore()
const albums = albumStore.albums
</script>

<template>
  <DataTable :value="albums" stripedRows tableStyle="min-width: 50rem">
    <Column field="title" header="Title"></Column>
    <Column field="artist" header="Artist"></Column>
    <Column field="date" header="Date"></Column>
    <Column field="image" header="Image"></Column>
    <Column field="link" header="Link"></Column>
  </DataTable>
</template>
