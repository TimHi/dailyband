import type { Album } from '@/model/album'
import PocketBase from 'pocketbase'
const pb = new PocketBase('http://localhost:9000/')

export async function GetDaily(): Promise<Album[]> {
  return await pb.collection('album').getFullList({
    sort: '-created'
  })
}
