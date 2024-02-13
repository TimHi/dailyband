<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum

</script>

<template>
  <img src="../assets/images/eel.png" />
  <main class="center">
    <div class="container">
      <h1>Try This Tune: Daily Camp's Musical Recommendation</h1>
      <img class="item" :src="dailyAlbum.image" />
      <div>
        <a :href="'https://daily.bandcamp.com/' + dailyAlbum.link">
          <h2 class="item">{{ dailyAlbum.artist[0] }} - {{ dailyAlbum.title }}</h2>
          <p class="item">
            Album of the day: {{ dailyAlbum.date }}
          </p>
        </a>
      </div>
      <div>
        <h3>Keep the music fresh and exciting! Explore the eclectic world of Bandcamp's <a
            href="https://daily.bandcamp.com/album-of-the-day">'Album of the Day'</a> with daily
          randomized picks.</h3>
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

.container {
  max-width: 400px;

  flex-wrap: wrap;
}

.border {
  border: 1px solid #383a3f;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 10px;
  /* Add some padding for better visual appearance */
  text-align: center;
  /* Center text horizontally */
  height: auto;
  width: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
