import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { useAlbumStore } from './stores/album'
import PrimeVue from 'primevue/config'
import 'primevue/resources/themes/lara-dark-cyan/theme.css'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faMusic, faArrowRight } from '@fortawesome/free-solid-svg-icons'

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faMusic)
library.add(faArrowRight)
const app = createApp(App)
app.use(PrimeVue, { ripple: true })
app.use(createPinia())
app.use(router)
app.component('font-awesome-icon', FontAwesomeIcon)
const albumStore = useAlbumStore()
await albumStore.fetchDailys()
app.mount('#app')
