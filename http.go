package main

import (
	"fmt"
	"net/http"
)

func main() {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "<h1>Hellow World !</h1>")
	}

	http.HandleFunc("/", myHandler)

	http.HandleFunc("/cat", myCatHandlerFunc)

	http.ListenAndServe(":9000", nil)
}

func myCatHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cat man")

}
