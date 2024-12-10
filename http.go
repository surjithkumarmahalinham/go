package main

import (
	"fmt"
	"net/http"
)

func main() {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow World !")
	}

	http.HandleFunc("/", myHandler)

	http.HandleFunc("/cat", myCatHandlerFunc)

	http.ListenAndServe(":9000", nil)
}

func myCatHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello cat man")

}
