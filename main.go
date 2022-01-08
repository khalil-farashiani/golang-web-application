package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("application is running on port%s  http://localhost:8080", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
