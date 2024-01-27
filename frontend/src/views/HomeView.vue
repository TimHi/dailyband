<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum
</script>

<template>
  <main>
    <h1>You might like</h1>
    <a :href="'https://daily.bandcamp.com/' + dailyAlbum.link">
      <img :src="dailyAlbum.image" />
      <h2>{{ dailyAlbum.artist }} - {{ dailyAlbum.title }}</h2>
      <h3>Daily Album from {{ dailyAlbum.date }}</h3>
    </a>
  </main>
</template>
