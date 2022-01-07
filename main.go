package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of byte %d", n))
	})
	_ = http.ListenAndServe(":8080", nil)
}
