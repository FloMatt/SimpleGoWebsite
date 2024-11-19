package handlers

import (
	"net/http"

	"github.com/FloMatt/SimpleGoWebsite/pkg/config"
	"github.com/FloMatt/SimpleGoWebsite/pkg/models"
	"github.com/FloMatt/SimpleGoWebsite/pkg/render"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Newhanders sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// NewHandlers creates a new handlers
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
