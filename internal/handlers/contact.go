package handlers

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"github.com/fouched/go-contact-app/internal/repository"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (m *HandlerConfig) ContactsNewGet(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["contact"] = models.Contact{}

	render.Template(w, r, "/contacts.new.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ContactsNew persists a contact and redirects to the list page
func (m *HandlerConfig) ContactsNew(w http.ResponseWriter, r *http.Request) {
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

	_, err := repository.InsertContact(contact)

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
	}

	m.App.Session.Put(r.Context(), "success", "Contact created")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// ContactsList displays contacts
func (m *HandlerConfig) ContactsList(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	render.Template(w, r, "/contacts.list.tmpl", &models.TemplateData{
		//Data: data,
	})
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
	render.Template(w, r, "/contacts.view.tmpl", &models.TemplateData{
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
	render.Template(w, r, "/contacts.edit.tmpl", &models.TemplateData{
		Data: data,
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

	contact := models.Contact{
		ID:    contactId,
		First: r.Form.Get("first"),
		Last:  r.Form.Get("last"),
		Phone: r.Form.Get("phone"),
		Email: r.Form.Get("email"),
	}

	err = repository.UpdateContactById(contact)
	if err != nil {
		fmt.Println("Could not update contact", err)
		return
	}

	m.App.Session.Put(r.Context(), "success", "Contact updated")
	http.Redirect(w, r, "/contacts/", http.StatusSeeOther)
}

func (m *HandlerConfig) ContactsDeletePost(w http.ResponseWriter, r *http.Request) {
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
