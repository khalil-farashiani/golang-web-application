package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/khalil-farashiani/golang-web-application/internal/config"
	"github.com/khalil-farashiani/golang-web-application/internal/models"
	"github.com/khalil-farashiani/golang-web-application/internal/render"
	"log"
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
	remoteAddress := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_addr", remoteAddress)
	render.RenderTemplates(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_addr")
	stringMap["remote_IP"] = remoteIP

	render.RenderTemplates(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Contact renders the room page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Availability renders the search-availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Massege string `json:"massege"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Massege: "Available!",
	}

	json, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}

	log.Println(string(json))

	w.Header().Set("Content-Type", "Application/json")
	w.Write(json)
}

// PostAvailability renders the search-availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("from %s to %s", start, end)))
	render.RenderTemplates(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}
