<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum
const colors = await Get
</script>

<template>
  <main class="center">
    <div class="card">
      <h1>Try this tune</h1>
      <p>Daily Camp's Musical Recommendation</p>
      <h2 class="dateTitle">{{ dailyAlbum.date }}</h2>
      <div class="images">
        <img class="item" :src="dailyAlbum.image" />
        <div>
          <img class="backgroundImage" src="../assets/images/half-vinyl.png" />
          <img class="coverBackground" :src="dailyAlbum.image" />
        </div>
      </div>
      <div class="container">
        <h3>
          <b>{{ dailyAlbum.title }}</b>
        </h3>
        <h4>by {{ dailyAlbum.artist[0] }}</h4>
      </div>
      <a :href="'https://daily.bandcamp.com' + dailyAlbum.link">Visit BandCamp</a>
      <a href="/data">Browse full data</a>
    </div>
  </main>
</template>
<style>
.linkContainer {
}

.dateTitle {
  color: #b12323;
}

.images {
  position: relative;
  display: flex;
  align-items: center;
  min-width: 500px;
  justify-content: center;
}

.coverBackground {
  position: absolute;
  left: 55%;
  top: 35%;
  height: 100px;
  transform: rotate(48deg);
  z-index: -1;
}

.backgroundImage {
  height: 280px;
  position: relative;
  z-index: 0;
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
  z-index: 1;
  position: relative;
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
