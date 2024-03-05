import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { useAlbumStore } from './stores/album'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/lara-dark-cyan/theme.css'
const app = createApp(App)
app.use(PrimeVue, { ripple: true })
app.use(createPinia())
app.use(router)
const albumStore = useAlbumStore()
await albumStore.fetchDailys()
app.mount('#app')
