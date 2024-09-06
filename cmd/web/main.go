package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-contact-app/internal/config"
	"github.com/fouched/go-contact-app/internal/handlers"
	"github.com/fouched/go-contact-app/internal/helpers"
	"github.com/fouched/go-contact-app/internal/render"
	"github.com/fouched/go-contact-app/internal/repo"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jaswdr/faker/v2"
	"log"
	"net/http"
	"time"
)

const port = ":9080"
const dbString = "host=localhost port=5432 dbname=contact_app user=fouche password=javac"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	dbPool, err := run()
	if err != nil {
		log.Fatalln(err)
	}

	//seed(dbPool)

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

	// register complex type for session
	gob.Register(helpers.Archive{})
	// create the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.InProduction = false

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)

	return dbPool, nil
}

func seed(db *sql.DB) {
	fmt.Println("Seeding database")
	fake := faker.New()

	person := fake.Person()
	phone := fake.Phone()
	for i := 0; i < 100000; i++ {
		firstName := person.FirstName()
		lastName := person.LastName()

		stmt := `INSERT INTO contacts (first, last, phone, email, created_at, updated_at)
    			VALUES($1, $2, $3, $4, $5, $6)`

		db.Exec(stmt, firstName, lastName, phone.E164Number(),
			firstName+"@"+lastName+".com", time.Now(), time.Now())

	}
}
