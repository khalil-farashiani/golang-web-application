package handlers

import (
	"github.com/khalil-farashiani/golang-web-application/pkg/config"
	"github.com/khalil-farashiani/golang-web-application/pkg/models"
	"github.com/khalil-farashiani/golang-web-application/pkg/render"
	"net/http"
)

//Repo repository used by handlers
var Repo *Repository

//Repository is a Repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{})
}
