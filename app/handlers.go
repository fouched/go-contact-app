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

// ContactsList is the contacts page handler
func ContactsList(w http.ResponseWriter, r *http.Request) {

	//TODO the db stuff should live elsewhere - fine for now
	rows, err := db.Query("SELECT * FROM contacts")
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

	RenderTemplate(w, "contacts.list.tmpl", contacts)
}

// ContactsAdd is the contacts view and edit page handler
func ContactsAdd(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contacts.add.tmpl", nil)
}

// ContactsNew is the new contact page handler
func ContactsNew(w http.ResponseWriter, r *http.Request) {
	fe := r.ParseForm()
	if fe != nil {
		fmt.Println("Cannot parse form", fe)
		return
	}

	//TODO the db stuff should live elsewhere - fine for now
	var id int
	stmt := `INSERT INTO contacts (first, last, phone, email)
    			VALUES($1, $2, $3, $4) returning id`

	err := db.QueryRow(stmt,
		r.Form.Get("first"),
		r.Form.Get("last"),
		r.Form.Get("phone"),
		r.Form.Get("email"),
	).Scan(&id)

	if err == nil {
		fmt.Println(fmt.Sprintf("Inserted contact with id %d", id))
	} else {
		fmt.Println("DB error, cannot insert contact", fe)
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// RenderTemplate TODO extract below to cache template parsing in a production environment
func RenderTemplate(w http.ResponseWriter, tmpl string, td any) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
