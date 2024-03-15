package scraper

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/timhi/bandcamp-scraper/m/v2/src/model"
)

var (
	OPENING_QUOTE = "“"
	CLOSING_QUOTE = "”"
	BASE_URL      = "https://daily.bandcamp.com/album-of-the-day?page="
)

func Scrape() []model.Album {
	albums := []model.Album{}
	totalPages := getPageCount()
	log.Println(totalPages)
	var wg sync.WaitGroup
	for i := 1; i <= totalPages; i++ {
		tPage := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			scrapedAlbums := scrapePage(tPage)
			albums = append(albums, scrapedAlbums...)
		}()
	}
	wg.Wait()
	return albums
}

func scrapePage(page int) []model.Album {
	albums := []model.Album{}

	c := colly.NewCollector()
	c.OnHTML(".aotd", func(e *colly.HTMLElement) {
		album := model.Album{}
		err := scrapeAlbum(e, &album)
		if err != nil {
			log.Println(err)
		} else {
			scrapeDetails(&album)
			albums = append(albums, album)
		}
	})

	albumPage := BASE_URL + fmt.Sprint(page)
	c.Visit(albumPage)
	return albums
}

func getPageCount() int {
	c := colly.NewCollector()

	pageCount := 0
	c.OnHTML(".album-of-the-day", func(e *colly.HTMLElement) {
		pageCount++
		newUrl := BASE_URL + fmt.Sprint(pageCount)
		c.Visit(newUrl)
	})
	c.Visit(BASE_URL + fmt.Sprint(0))
	return pageCount
}

func scrapeDetails(a *model.Album) {
	g := colly.NewCollector()
	g.OnHTML("div.genre a", func(e *colly.HTMLElement) {
		linkText := e.Text
		a.Genres = append(a.Genres, linkText)
	})

	g.OnHTML(".aotd", func(e *colly.HTMLElement) {
		rawText := e.ChildText("p")
		if rawText != "" {
			log.Println("Found Description")
		}
		a.Descriptions = append(a.Descriptions, rawText)
	})

	g.Visit("https://daily.bandcamp.com" + a.Link)
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

	a.Artist = artists[0]
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
