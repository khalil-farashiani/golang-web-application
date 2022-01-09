package main

import (
	"fmt"
	"github.com/khalil-farashiani/golang-web-application/pkg/config"
	"github.com/khalil-farashiani/golang-web-application/pkg/handlers"
	"github.com/khalil-farashiani/golang-web-application/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	ts, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = ts
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("application is running on port%s  http://localhost:8080", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
