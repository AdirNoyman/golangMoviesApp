package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set CORS headers only on development(http request and not https)
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		// If the request is a OPTIONS http method request -> send all the headers our server supports
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,X-CSRF-Token,Authorization")
			return
		} else {

			// Just send the regular http response
			h.ServeHTTP(w, r)

		}

	})
}
