package handlers

import (
	"github.com/khalil-farashiani/golang-web-application/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}
