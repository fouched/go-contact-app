package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// Contacts is the contacts page handler
func Contacts(w http.ResponseWriter, r *http.Request) {
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
