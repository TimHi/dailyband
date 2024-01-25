package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/timhi/bandcamp-scraper/m/v2/src/model"
)

func SendParsedData(albums *[]model.Album) error {
	backend := "http://localhost:9000/api/upload"

	// Convert albums slice to JSON
	albumsJSON, err := json.Marshal(albums)
	if err != nil {
		return fmt.Errorf("error marshaling albums to JSON: %v", err)
	}

	// Create a buffer with the JSON data
	buffer := bytes.NewBuffer(albumsJSON)

	resp, err := http.Post(backend, "application/json", buffer)
	if err != nil {
		return fmt.Errorf("error sending data to backend: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("backend returned non-OK status: %v", resp.Status)
	}

	// Optionally, you can read the response body if needed

	return nil
}
