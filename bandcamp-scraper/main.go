package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/timhi/bandcamp-scraper/m/v2/src/network"
	"github.com/timhi/bandcamp-scraper/m/v2/src/scraper"
)

func main() {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	psqlPw := os.Getenv("PSQL_PW")
	scrapedAlbums := scraper.Scrape()
	err := network.SendParsedDataToPocketbase(&scrapedAlbums)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Succesfully sent scraped data to pocketbase backend")
	}
	err_sup := network.SendParsedDataToSubabase(&scrapedAlbums, psqlPw)
	if err_sup != nil {
		log.Println(err_sup)
	} else {
		log.Println("Succesfully sent scraped data to supabase backend")
	}
}
