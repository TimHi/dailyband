<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum
</script>

<template>
  <main class="center">
    <div class="border">
      <img class="item" :src="dailyAlbum.image" />
      <div>
        <h1 class="item">{{ dailyAlbum.artist }} - {{ dailyAlbum.title }}</h1>
        <h2 class="item">
          Daily Album from {{ dailyAlbum.date }} -
          <a :href="'https://daily.bandcamp.com/' + dailyAlbum.link">Link 🔗</a>
        </h2>
      </div>
    </div>
  </main>
</template>

<style>
.center {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
.border {
  border: 1px solid #383a3f;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 10px; /* Add some padding for better visual appearance */
  text-align: center; /* Center text horizontally */
  height: auto;
  width: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
