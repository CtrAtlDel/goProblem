package main

import (
	"github.com/go-redis/redis"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange("comments", 0, 20).Result()
	if err!=nil {
		return 
	}

	templates = template.Must(templates.ParseGlob("templates/*.html"))
	templates.ExecuteTemplate(w, "index.html", comments)
}
