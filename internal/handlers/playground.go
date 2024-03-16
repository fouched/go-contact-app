package handlers

import (
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"net/http"
)

func (m *HandlerConfig) PlaygroundGet(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "/playground.htmx.tmpl", &models.TemplateData{})
}
