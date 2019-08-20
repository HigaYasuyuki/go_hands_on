package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("template.html"))
		tpl.Execute(w, "Hello, template")
	})

	http.ListenAndServe(":8085", r)
}
