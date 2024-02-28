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
      <div class="backgroundvinyl">
        <img class="item" :src="dailyAlbum.image" />
      </div>
      <div class="container">
        <h3>
          <b>{{ dailyAlbum.title }}</b>
        </h3>
        <h4>by {{ dailyAlbum.artist[0] }}</h4>
        <p>
          Lorem ipsum der aal Aalt ich brauch hier ein wenig Content am besten so viel das es
          Overflowt also schreib ich hier einfach mal weiter. Sollte passen jetzt. Oder ein zwei
          Sätze mehr sollten auch nicht wehtun.
        </p>
      </div>
    </div>
  </main>
</template>
<style>
img {
  height: 300px;
  border-radius: 5%;
}

.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
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
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  transition: 0.3s;
  background-color: rgb(89, 99, 99);
  border-radius: 2%;
  text-align: center;
  margin: auto;
}

.backgroundvinyl {
  background-image: url(../assets/images/vinyl.png);
  background-size: cover;
  background-repeat: no-repeat;
  background-position: 20px;
  height: 250px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.item {
  height: 200px;
  width: auto;
  border-radius: 5%;
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
