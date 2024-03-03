package main

import (
	"fmt"
	"github.com/fouched/go-contact-app/app/models"
	"github.com/fouched/go-contact-app/app/repo"
	"html/template"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func ContactsView(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["contact"] = models.Contact{}

	RenderTemplate(w, r, "contacts.view.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ContactsList displays contacts
func ContactsList(w http.ResponseWriter, r *http.Request) {
	err, contacts := repo.SelectContacts()
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contacts"] = contacts

	RenderTemplate(w, r, "contacts.list.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ContactsAdd persists a contact and redirects to the list page
func ContactsAdd(w http.ResponseWriter, r *http.Request) {
	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	contact := models.Contact{
		First: r.Form.Get("first"),
		Last:  r.Form.Get("last"),
		Phone: r.Form.Get("phone"),
		Email: r.Form.Get("email"),
	}

	err, _ := repo.AddContact(contact)

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
	}

	session.Put(r.Context(), "success", "Contact created")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// RenderTemplate TODO extract below to cache template parsing in a production environment
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Success = session.PopString(r.Context(), "success")
	td.Warning = session.PopString(r.Context(), "warning")
	td.Error = session.PopString(r.Context(), "error")

	return td
}
