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

const colorService = 'http://localhost:1323/colors/'

export interface ImageColor {
  Y: number
  Cb: number
  Cr: number
}

export interface HexColor {
  color: string
}

export async function GetColors(url: string): Promise<HexColor[]> {
  const response = await fetch(colorService + url, {
    method: 'GET',
    headers: {
      Accept: 'application/json'
    }
  })
  const result = (await response.json()) as HexColor[]
  return result
}
