package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/timhi/bandcamp-scraper/m/v2/src/model"
)

var (
	OPENING_QUOTE = "“"
	CLOSING_QUOTE = "”"
)

func Scrape() []model.Album {
	c := colly.NewCollector()
	albums := []model.Album{}

	c.OnHTML(".aotd", func(e *colly.HTMLElement) {
		album := model.Album{}
		err := scrapeAlbum(e, &album)
		if err != nil {
			log.Println(err)
		} else {
			albums = append(albums, album)
		}
	})

	c.OnHTML(".pagination-link", func(e *colly.HTMLElement) {
		if e.ChildText(".back-text") == "← Older posts" {
			log.Println(e.Attr("href"))
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.Visit("https://daily.bandcamp.com/album-of-the-day")
	return albums
}

func scrapeAlbum(e *colly.HTMLElement, a *model.Album) error {
	a.Image = e.ChildAttr("img", "src")
	a.Link = e.ChildAttr("a", "href")
	rawTitle := e.ChildText(".title")

	titleParts := strings.Split(rawTitle, ",")
	artists, albumTitle, err := processTitle(titleParts)
	if err != nil {
		return fmt.Errorf("error parsing artists and titles: %v", rawTitle)
	}

	a.Artist = artists
	a.Title = albumTitle

	dateText := e.ChildText(".article-info-text")
	parts := strings.Split(dateText, "·")
	if len(parts) < 1 {
		return fmt.Errorf("error parsing date: %v", dateText)
	} else {
		a.Date = strings.TrimSpace(parts[1])
	}
	return nil
}

func processTitle(titleParts []string) ([]string, string, error) {
	if len(titleParts) != 2 {
		return []string{}, "", fmt.Errorf("titleparts are not parsable")
	}
	artist := strings.TrimSpace(titleParts[0])
	title := strings.TrimSpace(strings.Trim(titleParts[1], "\""))
	title = strings.ReplaceAll(title, "”", "")
	title = strings.ReplaceAll(title, "“", "")
	title = strings.TrimSpace(title)
	return []string{artist}, title, nil
}
