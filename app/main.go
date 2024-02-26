package main

import (
	"database/sql"
	"fmt"
	"html/template"
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
	// close db conn pool after app stop
	defer dbPool.Close()

	http.HandleFunc("/", Home)
	http.HandleFunc("/contacts", Contacts)

	fmt.Println(fmt.Sprintf("Starting application on %s", port))
	log.Fatalln(http.ListenAndServe(port, nil))
}

func run() (*sql.DB, error) {
	conn, err := DbPool(dbString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	return conn, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact.page.tmpl")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

func DbPool(dsn string) (*sql.DB, error) {
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
