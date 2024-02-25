package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port = ":8000"

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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contacts", Contacts)

	fmt.Println(fmt.Sprintf("Starting application on %s", port))
	log.Fatalln(http.ListenAndServe(port, nil))
}
