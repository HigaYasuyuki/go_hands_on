package main

import (
	"html/template"
	"net/http"
	"strconv"

	"./model"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	tasks := model.GetTaskList()

	t := template.Must(template.ParseFiles("./templates/base.html", "./templates/index.html"))
	t.ExecuteTemplate(w, "layout", tasks)
}

func detail(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	var id int
	id, _ = strconv.Atoi(rawID)
	task, notFound := model.GetTask(id)

	if notFound {
		http.NotFound(w, r)
	} else {
		t := template.Must(template.ParseFiles("./templates/base.html", "./templates/detail.html"))
		t.ExecuteTemplate(w, "layout", task)
	}

}

func edit(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	var id int
	id, _ = strconv.Atoi(rawID)

	r.ParseForm()
	description := r.Form.Get("description")
	isComplete := r.Form.Get("isComplete")

	notFound := model.UpdateTask(id, description, isComplete == "1")

	if notFound {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/base.html", "./templates/create.html"))
	t.ExecuteTemplate(w, "layout", "")
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	description := r.Form.Get("description")

	model.RegisterTask(description)

	http.Redirect(w, r, "/", http.StatusFound)
}

func delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rawID := r.Form.Get("id")
	var id int
	id, _ = strconv.Atoi(rawID)

	notFound := model.DeleteTask(id)

	if notFound {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func complete(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func incompleteList(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	model.Init()
	defer model.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/detail/{id:[0-9]+}", detail)
	r.HandleFunc("/edit/{id:[0-9]+}", edit).Methods("POST")
	r.HandleFunc("/create", create)
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/delete", delete).Methods("POST")
	r.HandleFunc("/complete/{id:[0-9]+}", complete)
	r.HandleFunc("/incompleteList", incompleteList)

	http.ListenAndServe(":8085", r)
}
