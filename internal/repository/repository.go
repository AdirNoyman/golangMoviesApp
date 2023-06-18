package repository

import "backend/internal/models"

type DatabaseRepo interface {

	// GetAllMovies returns all movies from the DB, as a slice(array) and a possible error
	GetAllMovies() ([]*models.Movie, error)
}
