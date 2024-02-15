<script setup lang="ts">
import { useAlbumStore } from '@/stores/album'

const albumStore = useAlbumStore()
const { getDailyAlbum } = albumStore
const dailyAlbum = getDailyAlbum

</script>

<template>
    <main class="center">
        <div class="card">
            <img class="item" :src="dailyAlbum.image" />
            <div class="container">
                <h4><b>John Doe</b></h4>
                <p>Architect & Engineer</p>
            </div>
        </div>
    </main>
</template>

<style>
img {
    border-radius: 5%;
}

.card {
    /* Add shadows to create the "card" effect */
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
    transition: 0.3s;
    background-color: aqua;
    border-radius: 5%;
}

/* On mouse-over, add a deeper shadow */
.card:hover {
    box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
}

/* Add some padding inside the card container */
.container {
    padding: 2px 16px;
}

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
