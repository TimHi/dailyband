<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
const albumStore = useAlbumStore()

const albums = albumStore.albums
</script>

<template>
  <h1>Raw Data</h1>
  <DataTable :value="albums" paginator :rows="15" :rowsPerPageOptions="[5, 10, 20, 50]">
    <Column field="title" header="Title"></Column>
    <Column field="artist" header="Artist"></Column>
    <Column field="date" header="Date"></Column>
    <Column field="image" header="Image"></Column>
    <Column field="link" header="Link"></Column>
  </DataTable>
</template>
