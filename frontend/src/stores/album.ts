import type { Album } from '@/model/album'
import { GetDaily } from '@/services/data'
import { defineStore } from 'pinia'

export const useAlbumStore = defineStore('albumStore', {
  state: () => ({
    albums: [] as Album[]
  }),
  actions: {
    async fetchDailys() {
      const albums = await GetDaily()
      console.log(albums)
      this.albums = (await GetDaily()) as Album[]
    }
  },
  getters: {
    getDailyAlbum(): Album {
      return this.albums[Math.floor(Math.random() * this.albums.length)]
    }
  }
})
