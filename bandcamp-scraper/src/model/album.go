package model

type Album struct {
	Title  string   `json:"title"`
	Artist []string `json:"artist"`
	Date   string   `json:"date"`
	Genres []string `json:"genre"`
	Image  string   `json:"image"`
	Link   string   `json:"link"`
}
