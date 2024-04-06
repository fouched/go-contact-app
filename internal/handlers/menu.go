package handlers

import (
	"github.com/fouched/go-contact-app/internal/models"
	"github.com/fouched/go-contact-app/internal/render"
	"net/http"
)

func (m *HandlerConfig) Settings(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/settings.gohtml", &models.TemplateData{})
}

func (m *HandlerConfig) Help(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "/help.gohtml", &models.TemplateData{})
}
