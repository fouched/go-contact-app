package render

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/config"
	"github.com/fouched/go-contact-app/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	parsedTemplate, _ := template.ParseFiles(pathToTemplates+tmpl, pathToTemplates+"/base.layout.gohtml")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Success = app.Session.PopString(r.Context(), "success")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	return td
}
