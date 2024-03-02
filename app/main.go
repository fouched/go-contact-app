package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-contact-app/app/repo"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"time"
)

const port = ":8000"
const dbString = "host=localhost port=5432 dbname=contact_app user=fouche password=javac"

var app AppConfig
var session *scs.SessionManager

func main() {
	dbPool, err := run()
	if err != nil {
		log.Fatalln(err)
	}

	// we have database connectivity, close it after app stops
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
	dbPool, err := repo.CreateDbPool(dbString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	return dbPool, nil
}
