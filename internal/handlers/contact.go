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

func (m *HandlerConfig) ContactsView(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["contact"] = models.Contact{}

	render.Template(w, r, "/contacts.view.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ContactsList displays contacts
func (m *HandlerConfig) ContactsList(w http.ResponseWriter, r *http.Request) {
	err, contacts := repository.SelectContacts()
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contacts"] = contacts

	render.Template(w, r, "/contacts.list.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ContactsAdd persists a contact and redirects to the list page
func (m *HandlerConfig) ContactsAdd(w http.ResponseWriter, r *http.Request) {
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

	err, _ := repository.AddContact(contact)

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
	}

	m.App.Session.Put(r.Context(), "success", "Contact created")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (m *HandlerConfig) ContactsById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	data := make(map[string]int)
	data["id"] = contactId
	render.Template(w, r, "/contacts.test.tmpl", &models.TemplateData{
		IntMap: data,
	})
}
