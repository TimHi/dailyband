import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { useAlbumStore } from './stores/album'

const app = createApp(App)

app.use(createPinia())
app.use(router)
const albumStore = useAlbumStore()
await albumStore.fetchDailys()
app.mount('#app')
