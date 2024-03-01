<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum
</script>

<template>
  <main class="center">
    <div class="card">
      <h1>Try this tune</h1>
      <p>Daily Camp's Musical Recommendation</p>
      <h2>{{ dailyAlbum.date }}</h2>
      <div class="images">
        <img class="item" :src="dailyAlbum.image" />
        <div
          class="coverBackground"
          :style="{ backgroundImage: 'url(' + dailyAlbum.image + ')' }"
        ></div>
        <img class="backgroundImage" src="../assets/images/half-vinyl.png" />
      </div>
      <div class="container">
        <h3>
          <b>{{ dailyAlbum.title }}</b>
        </h3>
        <h4>by {{ dailyAlbum.artist[0] }}</h4>
        <a :href="'https://daily.bandcamp.com' + dailyAlbum.link"> Learn more </a>
      </div>
    </div>
  </main>
</template>
<style>
.images {
  display: flex;
  align-items: center;
  min-width: 500px;
  justify-content: center;
}

.coverBackground {
  background-size: cover;
  background-repeat: no-repeat;
  background-color: aqua;
}

.backgroundImage {
  height: 280px;
}

.container {
  padding: 2px 16px;
}

.center {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
.card {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: 0.3s;
  background-color: #2a2a2a;
  border-radius: 2%;
  text-align: center;
  margin: auto;
  min-width: 500px;
  border-color: #444444;
  border-style: solid;
}

.backgroundvinyl {
  background-image: url(dailyAlbum.image);
  background-size: contain;
  background-repeat: no-repeat;
  height: 250px;
  display: flex;
  align-items: center;
  background-position: center;
  justify-content: center;
  min-width: 200px;
}

.item {
  width: auto;
  border-radius: 5%;
  height: 300px;
}

.container {
  padding: 2px 16px;
  max-width: 400px;
  flex-wrap: wrap;
}

h1,
h2,
h3 {
  text-align: center;
}

.border {
  border: 1px solid #383a3f;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 10px;
  text-align: center;
  height: auto;
  width: auto;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
