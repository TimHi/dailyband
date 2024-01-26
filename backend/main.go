package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/timhi/bandcamp-scraper/backend/m/v2/src/model"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/upload",
			Handler: func(c echo.Context) error {
				var requestPayload []model.Album

				rawBody, err := io.ReadAll(c.Request().Body)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "Internal server error",
					})
				}

				if err := json.Unmarshal(rawBody, &requestPayload); err != nil {
					return c.JSON(http.StatusBadRequest, map[string]string{
						"error": "Invalid JSON payload",
					})
				}

				app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
					rawQuery := "DELETE FROM album"
					if _, err := txDao.DB().NewQuery(rawQuery).Execute(); err != nil {
						log.Println(err)
						return c.JSON(http.StatusInternalServerError, map[string]string{
							"error": "Error deleting the old data",
						})
					}

					return nil
				})

				albumCollection, err := app.Dao().FindCollectionByNameOrId("album")
				if err != nil {
					return err
				}

				for _, album := range requestPayload {
					record := models.NewRecord(albumCollection)

					record.Set("title", album.Title)
					record.Set("artist", album.Artist[0])
					record.Set("date", album.Date)
					record.Set("image", album.Image)
					record.Set("link", album.Link)
					record.Set("genres", "")
					record.MarkAsNew()
					if err := app.Dao().SaveRecord(record); err != nil {
						log.Println(err)
						return c.JSON(http.StatusInternalServerError, err)
					}
				}

				return c.JSON(http.StatusOK, "updated scraped albums")

			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				middleware.CORSWithConfig(middleware.CORSConfig{
					AllowOrigins: []string{"*"},
				}),
			},
		})
		return nil
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
