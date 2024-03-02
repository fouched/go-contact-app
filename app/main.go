package main

import (
	"database/sql"
	"fmt"
	"github.com/fouched/go-contact-app/app/data"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
)

const port = ":8000"
const dbString = "host=localhost port=5432 dbname=contact_app user=fouche password=javac"

func main() {
	dbPool, err := run()
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
	dbPool, err := data.CreateDbPool(dbString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	return dbPool, nil
}
