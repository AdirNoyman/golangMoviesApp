package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running...",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	var movies []models.Movie

	releaseDate1, _ := time.Parse("2006-01-02", "1986-03-07")
	releaseDate2, _ := time.Parse("2006-01-02", "1990-04-11")
	releaseDate3, _ := time.Parse("2006-01-02", "2000-05-08")

	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: releaseDate1,
		MPAARating:  "R",
		Duration:    116,
		Description: "There can be only one",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	raders := models.Movie{
		ID:          2,
		Title:       "Raders Of The Last Arc",
		ReleaseDate: releaseDate2,
		MPAARating:  "R",
		Duration:    126,
		Description: "The movie is great",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, raders)

	loveAllAround := models.Movie{
		ID:          3,
		Title:       "Love All Around",
		ReleaseDate: releaseDate3,
		MPAARating:  "PG-13",
		Duration:    96,
		Description: "All you need is love...",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, loveAllAround)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
