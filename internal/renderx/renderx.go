package renderx

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/config"
	"github.com/fouched/go-contact-app/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates/htmx"
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles(pathToTemplates + tmpl)
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
