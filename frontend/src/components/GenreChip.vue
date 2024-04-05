<script setup lang="ts">
import Chip from 'primevue/chip';
const props = defineProps({
    genreTitle: String
})
const genreLink = "https://bandcamp.com/discover/" + props.genreTitle?.toLowerCase();
</script>

<template>
    <a :href="genreLink">
        <Chip :label="props.genreTitle" />
    </a>
</template>
