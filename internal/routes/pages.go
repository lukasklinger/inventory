package routes

import (
	"html/template"
	"net/http"

	"cyaniccerulean.com/inventory/v2/internal/model"
)

type Pages struct {
	config model.Config
}

// pre-render templates
var templates = template.Must(template.ParseFiles("./web/template/index.html", "./web/template/entry.html"))

func InitPages(config model.Config) *Pages {
	return &Pages{config}
}

// handle requests for the index page
func (p Pages) IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", p.config)
}

// handle requests for individual entry pages
func (p Pages) EntryHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "entry.html", p.config)
}

// render the template using the service configuration
func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
