package handlers

import (
	"net/http"

	"github.com/xuoxod/go-app-template/pkg/config"
	"github.com/xuoxod/go-app-template/pkg/models"
	"github.com/xuoxod/go-app-template/pkg/render"
)

// The repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["subheading"] = "Welcome to my awesome Golang web application"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["subheading"] = "Who we are is totally and completely irrelevant"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
