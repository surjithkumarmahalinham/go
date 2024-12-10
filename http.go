package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	Name    string
	Age     int
	Address string
}

func main() {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		p := pageData{Name: "msk", Age: 27, Address: "IsLand"}
		// data := "Hello html template first"
		t.Execute(w, p)
	}

	http.HandleFunc("/", myHandler)

	http.HandleFunc("/cat", myCatHandlerFunc)

	http.ListenAndServe(":9000", nil)
}

func myCatHandlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cat.html")

}
