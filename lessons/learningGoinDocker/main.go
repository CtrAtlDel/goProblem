package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "<h1>Hello from docker!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
