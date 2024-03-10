import type { Album } from '@/model/album'
import { createClient } from '@supabase/supabase-js'

export const supabase = createClient(
  'https://jepgwqafueosittvbdmt.supabase.co',
  import.meta.env.VITE_DB_ANON ?? ''
)

export async function GetDaily(): Promise<Album[]> {
  const { data, error } = await supabase.from('Album').select('*')
  if (error !== null) console.error(error)
  const albumList: Album[] = []
  data?.forEach((a) => albumList.push(a.album))
  return albumList
}

const colorService = 'localhost:1323/colors/'

export async function GetColors() {
  const response = await fetch('https://reqres.in/api/users', {
    method: 'GET',
    headers: {
      Accept: 'application/json'
    }
  })
  const result = await response.json()
  console.log(result)
}
