package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// db is the db connection pool used by handlers
var db *sql.DB

// SetHandlerDb sets the db connection pool for handlers
func SetHandlerDb(d *sql.DB) {
	db = d
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// Contacts is the contacts page handler
func Contacts(w http.ResponseWriter, r *http.Request) {
	//db.ExecContext()
	var greeting string
	err := db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)
	RenderTemplate(w, "contact.page.tmpl")
}

/*
TODO extract below to cache template parsing in a production environment
*/
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
