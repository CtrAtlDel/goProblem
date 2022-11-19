package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(templates.ParseGlob("templates/*.html"))
	templates.ExecuteTemplate(w, "index.html", nil)
	fmt.Fprintf(w, "Hello, Gophers")
}
