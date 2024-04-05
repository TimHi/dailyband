<script setup lang="ts">
import { GetColors } from '@/services/data'
import { useAlbumStore } from '@/stores/album'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum
const colors = await GetColors(dailyAlbum.image)
if (dailyAlbum.descriptions === undefined || dailyAlbum.descriptions === null)
  dailyAlbum.descriptions = []
const dateColor = `color: ${colors[0]}`
</script>

<template>
  <main class="center">
    <div class="card">
      <h1>Try this tune</h1>
      <p>Daily Camp's Musical Recommendation</p>
      <div class="line-container">
        <div class="line"></div>
        <font-awesome-icon icon="music" :style="{ color: colors[0] }" />
        <div class="line"></div>
      </div>

      <h2 class="shadow" :style="dateColor">{{ dailyAlbum.date }}</h2>
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
        <h4>by {{ dailyAlbum.artist }}</h4>
        <div>
          <p v-if="dailyAlbum.descriptions !== undefined && dailyAlbum.descriptions.length > 0" class="line-clamp">
            {{ dailyAlbum.descriptions[0] }}
          </p>
          <div class="link-container">
            <a class="link-style" :href="'https://daily.bandcamp.com' + dailyAlbum.link">
              <p class="shadow" v-if="dailyAlbum.descriptions !== undefined && dailyAlbum.descriptions.length > 0">More
              </p>
              <p class="shadow" v-else>Details</p>
              <font-awesome-icon class="shadow" icon="arrow-right" :style="{
          color: colors[0],
          display: 'flex',
          alignSelf: 'center',
          marginLeft: '4px'
        }" />
            </a>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<style>
.shadow {
  text-shadow: 1px 1px #3b3b3b;
}

.link-style {
  display: flex;
}

.link-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.line-container {
  display: flex;
  align-items: center;
  position: relative;
  width: 100%;
}

.line {
  border: 1px solid #383a3f;
  z-index: 1;
  width: 50%;
}

.character {
  padding: 0 10px;
  /* Adjust as needed */
}

.line-clamp {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  overflow: hidden;
  -webkit-line-clamp: 3;
  max-height: calc(1.2em * 4);
  text-align: left;
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
