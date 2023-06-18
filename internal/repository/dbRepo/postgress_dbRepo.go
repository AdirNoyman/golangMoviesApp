package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgressDBRepo struct {
	DB *sql.DB
}

// dbTimeOut is the maximum amount of time a database query can take before timing out
const dbTimeOut = time.Second * 3

func (m *PostgressDBRepo) GetAllMovies() ([]*models.Movie, error) {
	// The session will timeout after the 3 seconds in case there is no interaction with the database 🤨
	sessionContext, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	// Build our Query for the database
	query := `
		SELECT id, title, description,release_date, duration, mpaa_rating, created_at, updated_at, coalesce(image, '')
		FROM movies
		ORDER BY title
		`

	// Execute the query and assign the results set to the rows variable
	rows, err := m.DB.QueryContext(sessionContext, query)
	// Return nil for the data and an error in case the query failed
	if err != nil {

		return nil, err
	}

	// Close the rows after the query ended
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {

		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Duration, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt, &movie.Image)

		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}
