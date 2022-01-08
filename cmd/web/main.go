package main

import (
	"fmt"
	"github.com/khalil-farashiani/golang-web-application/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("application is running on port%s  http://localhost:8080", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
