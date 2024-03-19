package models

import "github.com/fouched/go-contact-app/internal/validation"

// TemplateData holds repo sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Success   string
	Warning   string
	Error     string
	Form      *validation.Form
}
