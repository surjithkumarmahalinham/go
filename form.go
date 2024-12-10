package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Company string
	Data    string
}

func main() {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		p := PageData{Title: "Go", Company: "Msk Multinational", Data: "A paragraph is a group of sentences that are related to a single topic and work together to develop a main idea. Paragraphs are a fundamental part of writing, and are used to organize essays and other longer pieces of writing."}

		t.Execute(w, p)
	}

	http.HandleFunc("/", myHandler)
	http.HandleFunc("/action", myActionfunc)
	http.ListenAndServe(":9001", nil)
}

func myActionfunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "action.html")
	} else {
		user := r.FormValue("username")
		pass := r.FormValue("password")

		fmt.Fprintf(w, "Username %s and Password %s", user, pass)

	}

}
