package main

import (
	"fmt"
	"github.com/fouched/go-contact-app/app/data"
	"html/template"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func ContactsView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contacts.view.tmpl", nil)
}

// ContactsList displays contacts
func ContactsList(w http.ResponseWriter, r *http.Request) {
	err, contacts := data.SelectContacts()
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	RenderTemplate(w, "contacts.list.tmpl", contacts)
}

// ContactsAdd persists a contact and redirects to the list page
func ContactsAdd(w http.ResponseWriter, r *http.Request) {
	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	err, _ := data.AddContact(
		r.Form.Get("first"),
		r.Form.Get("last"),
		r.Form.Get("phone"),
		r.Form.Get("email"))

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
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
