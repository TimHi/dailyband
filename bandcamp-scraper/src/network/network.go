package network

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/timhi/bandcamp-scraper/m/v2/src/model"
)

func SendParsedDataToPocketbase(albums *[]model.Album) error {
	backend := "http://localhost:9000/api/upload"

	albumsJSON, err := json.Marshal(albums)
	if err != nil {
		return fmt.Errorf("error marshaling albums to JSON: %v", err)
	}

	buffer := bytes.NewBuffer(albumsJSON)

	resp, err := http.Post(backend, "application/json", buffer)
	if err != nil {
		return fmt.Errorf("error sending data to backend: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("backend returned non-OK status: %v", resp.Status)
	}

	return nil
}
func SendParsedDataToSubabase(albums *[]model.Album, psqlPw string) error {
	log.Println(*albums)
	const (
		host   = "aws-0-eu-central-1.pooler.supabase.com"
		port   = 5432
		user   = "postgres.jepgwqafueosittvbdmt"
		dbname = "postgres"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, psqlPw, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM public.\"Album\"")
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range *albums {
		albumJSON, err := json.Marshal(album)
		if err != nil {
			log.Println(err)
		}

		_, err = db.Exec("INSERT INTO public.\"Album\" (created_at, album) VALUES ($1, $2)", time.Now(), albumJSON)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer db.Close()
	log.Println("Done Db")
	return nil
}
