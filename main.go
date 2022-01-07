package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("application is running on port%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
