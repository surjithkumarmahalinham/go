package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow World !")
	})

	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello cat man")
	})

	http.ListenAndServe(":9000", nil)
}
