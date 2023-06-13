package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux (router)
	mux := chi.NewRouter()

	// middlewares
	// Will throw an http 500 error if our will panics and will log the error with backtrace and brings the application back up
	mux.Use(middleware.Recoverer)

	// routes
	mux.Get("/", app.Home)

	return mux
}
