package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "this is a home page wellcome")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "this is a about page")
}

func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("application is running on port%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
