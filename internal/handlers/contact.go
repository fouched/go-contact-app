package handlers

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"github.com/fouched/go-contact-app/internal/repository"
	"github.com/fouched/go-contact-app/internal/validation"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (m *HandlerConfig) ContactsNewGet(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["contact"] = models.Contact{}
	t := make(map[string]string)
	t["title"] = "Add"
	t["action"] = "/contacts/new"

	render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: t,
	})
}

// ContactsNewPost persists a contact and redirects to the list page
func (m *HandlerConfig) ContactsNewPost(w http.ResponseWriter, r *http.Request) {
	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	contact := parseForm(r)
	form := isValidContact(r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact
		t := make(map[string]string)
		t["title"] = "Add"
		t["action"] = "/contacts/new"
		render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
			Form:      &form,
			Data:      data,
			StringMap: t,
		})
		return
	}

	_, err := repository.InsertContact(contact)

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
	}

	m.App.Session.Put(r.Context(), "success", "Contact created")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func parseForm(r *http.Request) models.Contact {
	contact := models.Contact{
		First: r.Form.Get("first"),
		Last:  r.Form.Get("last"),
		Phone: r.Form.Get("phone"),
		Email: r.Form.Get("email"),
	}
	return contact
}

// isValidContact validates the form
func isValidContact(r *http.Request) validation.Form {

	// populate a new form with the post data
	form := validation.New(r.PostForm)
	// perform validation
	form.Required("first", "last", "phone", "email")
	form.IsEmail("email")
	form.MinLength("first", 2)
	form.MinLength("first", 2)

	return *form
}

// ContactsList displays contacts
func (m *HandlerConfig) ContactsList(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	render.Template(w, r, "/contacts.list.gohtml", &models.TemplateData{})
}

func (m *HandlerConfig) ContactsView(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	contact, err := repository.SelectContactById(contactId)
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contact"] = contact
	render.Template(w, r, "/contacts.view.gohtml", &models.TemplateData{
		Data: data,
	})
}

func (m *HandlerConfig) ContactsEditGet(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	contact, err := repository.SelectContactById(contactId)
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contact"] = contact
	t := make(map[string]string)
	t["title"] = "Edit"
	t["action"] = "/contacts/" + id + "/edit"
	render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: t,
	})
}

func (m *HandlerConfig) ContactsEditPost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	contact := parseForm(r)
	form := isValidContact(r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact
		t := make(map[string]string)
		t["title"] = "Edit"
		t["action"] = "/contacts/" + id + "/edit"
		render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
			Form:      &form,
			Data:      data,
			StringMap: t,
		})
		return
	}

	contact.ID = contactId
	err = repository.UpdateContactById(contact)
	if err != nil {
		fmt.Println("Could not update contact", err)
		return
	}

	m.App.Session.Put(r.Context(), "success", "Contact updated")
	http.Redirect(w, r, "/contacts/", http.StatusSeeOther)
}

func (m *HandlerConfig) ContactsDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	err = repository.DeleteContactById(contactId)
	if err != nil {
		fmt.Println("Could not delete contact", err)
		return
	}

	m.App.Session.Put(r.Context(), "success", "Contact deleted")
	http.Redirect(w, r, "/contacts/", http.StatusSeeOther)
}
