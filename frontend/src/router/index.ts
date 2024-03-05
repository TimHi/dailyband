import { createRouter, createWebHistory } from 'vue-router'

import CardView from '@/components/CardView.vue'
import DataView from '@/views/DataView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: CardView
    },
    {
      path: '/data',
      name: 'data',
      component: DataView
    }
  ]
})

export default router
