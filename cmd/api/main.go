package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DataSrcName string
	Domain      string
	DB          *sql.DB
}

func main() {

	// set application configuration
	var app application

	// read from command line arguments
	// 'flag' package is used to parse command line arguments. it reads arguments(flags) from the command line
	flag.StringVar(&app.DataSrcName, "dsn", "host=localhost port=5432 dbname=movies user=postgres password=postgres sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database
	connectionPool, err := app.connectToDB()
	if err != nil {

		log.Fatalf("Failed to connect to database: %v, application was kiiled in action 💐😢", err)
	}

	app.DB = connectionPool
	// Close all connections in the pool, after the program exits 💪
	defer app.DB.Close()

	app.Domain = "example.com"

	log.Println("Starting application on port:", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
