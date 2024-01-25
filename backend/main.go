package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()
	database := app.Dao()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
