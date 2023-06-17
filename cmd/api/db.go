package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// openDB opens a connection to the database specified by the DSN. *sql.DB = connection pool
func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping() // check if the connection is alive
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectToDB() (*sql.DB, error) {

	connection, err := openDB(app.DataSrcName)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgress database 😎🤘")

	return connection, nil
}
