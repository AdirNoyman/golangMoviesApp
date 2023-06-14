package models

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	MPAARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	// "-" -> means leave it empty (dont include it in the json response)
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
