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
	"github.com/timhi/bandcamp-scraper/backend/m/v2/src/model"
)

func main() {
	app := pocketbase.New()
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/upload",
			Handler: func(c echo.Context) error {
				log.Println("Received request to /api/upload")

				// Print the raw request body for debugging
				rawBody, err := io.ReadAll(c.Request().Body)
				if err != nil {
					log.Println("Error reading request body:", err)
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "Internal server error",
					})
				}
				//log.Println("Raw Request Body:", string(rawBody))

				// Decode JSON request
				var requestPayload []model.Album
				if err := json.Unmarshal(rawBody, &requestPayload); err != nil {
					log.Println("Error decoding JSON:", err)
					return c.JSON(http.StatusBadRequest, map[string]string{
						"error": "Invalid JSON payload",
					})
				}

				log.Println("Decoded JSON payload:", requestPayload)

				// Process the JSON payload as needed
				// You can access data like requestPayload["key"]

				return c.JSON(http.StatusOK, map[string]string{
					"message": "Data received successfully",
				})

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
