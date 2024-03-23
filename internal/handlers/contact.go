package handlers

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"github.com/fouched/go-contact-app/internal/repo"
	"github.com/fouched/go-contact-app/internal/validation"
	"github.com/go-chi/chi/v5"
	"io"
	"math"
	"net/http"
	"strconv"
)

func (m *HandlerConfig) ContactsNewGet(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["contact"] = models.Contact{}

	render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: MakeUpsertMap("Add", "/contacts/new"),
	})
}

// ContactsNewPost persists a contact and redirects to the list page
func (m *HandlerConfig) ContactsNewPost(w http.ResponseWriter, r *http.Request) {

	contact := parseContactForm(r)
	form := isValidContact(r, 0)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact
		render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
			Form:      &form,
			Data:      data,
			StringMap: MakeUpsertMap("Add", "/contacts/new"),
		})
		return
	}

	_, err := repo.InsertContact(contact)

	if err != nil {
		fmt.Println("DB error, cannot insert contact", err)
	}

	m.App.Session.Put(r.Context(), "success", "Contact created")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// ContactsListGet displays contacts
func (m *HandlerConfig) ContactsListGet(w http.ResponseWriter, r *http.Request) {
	intMap := make(map[string]int)
	intMap["cp"] = 1
	intMap["tp"] = 3
	render.Template(w, r, "/contacts.list.gohtml", &models.TemplateData{
		IntMap: intMap,
	})
}

func (m *HandlerConfig) ContactsListPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Cannot parse form", err)
	}

	q := r.Form.Get("q")
	cp := r.Form.Get("cp")
	cpInt, _ := strconv.Atoi(cp)
	first := r.Form.Get("first")
	last := r.Form.Get("last")
	prev := r.Form.Get("prev")
	next := r.Form.Get("next")

	count, _ := repo.SelectContactCount(q)

	pageSize := 3
	o := 0
	pages := math.Ceil(float64(count) / float64(pageSize))
	tp := int(pages)

	if cpInt == 0 {
		cpInt = 1
	}

	if first != "" {
		cpInt = 1
	} else if last != "" {
		cpInt = tp
		o = (cpInt - 1) * pageSize
	} else if prev != "" {
		cpInt = cpInt - 1
		o = (cpInt - 1) * pageSize
	} else if next != "" {
		cpInt = cpInt + 1
		o = (cpInt - 1) * pageSize
	}

	contacts, err := repo.SelectContacts(q, o, pageSize)
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contacts"] = contacts
	intMap := make(map[string]int)
	intMap["cp"] = cpInt
	intMap["tp"] = tp

	render.TemplateSnippet(w, r, "/contacts.results.gohtml", &models.TemplateData{
		Data:   data,
		IntMap: intMap,
	})
}

func (m *HandlerConfig) ContactsViewGet(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	contact, err := repo.SelectContactById(contactId)
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

	contact, err := repo.SelectContactById(contactId)
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contact"] = contact
	render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
		Form:      validation.New(nil),
		Data:      data,
		StringMap: MakeUpsertMap("Edit", "/contacts/"+id+"/edit"),
	})
}

func (m *HandlerConfig) ContactsEditPost(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	contact := parseContactForm(r)
	form := isValidContact(r, contactId)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact
		render.Template(w, r, "/contacts.upsert.gohtml", &models.TemplateData{
			Form:      &form,
			Data:      data,
			StringMap: MakeUpsertMap("Edit", "/contacts/"+id+"/edit"),
		})
		return
	}

	contact.ID = contactId
	err = repo.UpdateContactById(contact)
	if err != nil {
		fmt.Println("Could not update contact", err)
		return
	}

	m.App.Session.Put(r.Context(), "success", "Contact updated")
	http.Redirect(w, r, "/contacts/", http.StatusSeeOther)
}

func (m *HandlerConfig) ContactsEmailValidation(w http.ResponseWriter, r *http.Request) {

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

	emailExists, err := repo.EmailExists(r.Form.Get("email"), contactId)
	if err != nil {
		fmt.Println("Error checking email", err)
		return
	}

	if emailExists {
		_, _ = io.WriteString(w, "Email already taken!")
	}
}

func (m *HandlerConfig) ContactsDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contactId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Contact ID not an Integer", err)
		return
	}

	err = repo.DeleteContactById(contactId)
	if err != nil {
		fmt.Println("Could not delete contact", err)
		return
	}

	m.App.Session.Put(r.Context(), "success", "Contact deleted")
	http.Redirect(w, r, "/contacts/", http.StatusSeeOther)
}

// parseContactForm creates an instance of the Contact struct
func parseContactForm(r *http.Request) models.Contact {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return models.Contact{}
	}

	contact := models.Contact{
		First: r.Form.Get("first"),
		Last:  r.Form.Get("last"),
		Phone: r.Form.Get("phone"),
		Email: r.Form.Get("email"),
	}
	return contact
}

// isValidContact validates the form
func isValidContact(r *http.Request, id int) validation.Form {

	// populate a new form with the post data
	form := validation.New(r.PostForm)
	// perform validation
	form.Required("first", "last", "phone", "email")
	form.IsEmail("email")
	form.MinLength("first", 2)
	form.MinLength("last", 2)

	return *form
}
