package main

import (
	"encoding/json"
	"log"

	"github.com/timhi/bandcamp-scraper/m/v2/src/scraper"
)

func main() {
	scrapedAlbums := scraper.Scrape()
	data, err := json.Marshal(scrapedAlbums)
	if err != nil {
		panic("Failed to marshal parsed albums")
	}
	log.Println(string(data))
}
