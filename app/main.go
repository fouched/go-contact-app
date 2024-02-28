package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const port = ":8000"
const dbString = "host=localhost port=5432 dbname=contact_app user=fouche password=javac"

func main() {
	dbPool, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// close db conn pool after app stops
	defer dbPool.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() (*sql.DB, error) {
	conn, err := DatabasePool(dbString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	// set the db connection for all handlers
	SetHandlerDb(conn)

	return conn, nil
}

func DatabasePool(dsn string) (*sql.DB, error) {
	// no error thrown even host or db does not exist
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5 * time.Minute)

	// do a real test to see if we have a db conn
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
