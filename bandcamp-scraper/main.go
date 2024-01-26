package main

import (
	"log"

	"github.com/timhi/bandcamp-scraper/m/v2/src/network"
	"github.com/timhi/bandcamp-scraper/m/v2/src/scraper"
)

func main() {
	scrapedAlbums := scraper.Scrape()
	err := network.SendParsedData(&scrapedAlbums)
	if err != nil {
		log.Println(err)
	}
	log.Println("Succesfully sent scraped data to backend")
}
