package main

import (
	"database/sql"
	"fmt"
	"github.com/fouched/go-contact-app/app/models"
	"html/template"
	"net/http"
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

	//TODO the db stuff should live elsewhere - fine for now
	// and NO there will absolutely not usage of ORM - IMHO it is EVIL!
	rows, err := db.Query("SELECT * FROM contact")
	if err != nil {
		fmt.Println("DB query error, 0 rows will be returned")
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		err := rows.Scan(&c.ID, &c.First, &c.Last, &c.Phone, &c.Email)
		if err != nil {
			fmt.Println("DB query error, ignoring row")
		}
		contacts = append(contacts, c)
	}

	RenderTemplate(w, "contact.page.tmpl", contacts)
}

/*
TODO extract below to cache template parsing in a production environment
*/
func RenderTemplate(w http.ResponseWriter, tmpl string, td any) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
