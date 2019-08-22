package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type fizzBuzz struct {
	Number int
	Fizz   bool
	Buzz   bool
}

func display(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rawNum := r.Form.Get("number")
	num, _ := strconv.Atoi(rawNum)

	var list []fizzBuzz
	i := 1
	for i <= num {
		list = append(list, fizzBuzz{
			Number: i,
			Fizz:   i%3 == 0,
			Buzz:   i%5 == 0,
		})
		i++
	}
	t := template.Must(template.ParseFiles("base.html", "display.html"))
	t.ExecuteTemplate(w, "layout", list)
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
