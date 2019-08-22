package main

import (
	"html/template"
	"net/http"

	"./model"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	tasks := model.GetTaskList()

	t := template.Must(template.ParseFiles("./templates/base.html", "./templates/index.html"))
	t.ExecuteTemplate(w, "layout", tasks)
}

func main() {
	model.Init()
	defer model.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	// r.HandleFunc("/detail/{id:[0-9]+}", detail)
	// r.HandleFunc("/edit/{id:[0-9]+}", edit).Methods("POST")
	// r.HandleFunc("/create", create)
	// r.HandleFunc("/register", register).Methods("POST")
	// r.HandleFunc("/delete", delete).Methods("POST")

	http.ListenAndServe(":8085", r)
}
