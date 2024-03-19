package apix

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"github.com/fouched/go-contact-app/internal/repository"
	"net/http"
)

func (m *HtmxApiConfig) ContactsList(w http.ResponseWriter, r *http.Request) {

	pe := r.ParseForm()
	if pe != nil {
		fmt.Println("Cannot parse form", pe)
		return
	}

	contacts, err := repository.SelectContacts(r.Form.Get("q"))
	if err != nil {
		fmt.Println("DB error, cannot query contacts", err)
	}

	data := make(map[string]interface{})
	data["contacts"] = contacts

	render.TemplateSnippet(w, r, "/contacts.results.gohtml", &models.TemplateData{
		Data: data,
	})
}
