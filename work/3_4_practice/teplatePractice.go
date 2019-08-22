package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func display(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func input(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("base.html", "input.html"))
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", input)
	r.HandleFunc("/display", display).Methods("POST")
	http.ListenAndServe(":8085", r)
}
