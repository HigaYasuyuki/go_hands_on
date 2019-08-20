package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type user struct {
	Name string
	Age  int
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	// goの時間書式にはこの日付を使わなくてはならない 理由は
	// https://qiita.com/ruiu/items/5936b4c3bd6eb487c182
	// を参照
	return t.Format(layout)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		funcMap := template.FuncMap{"fdate": formatDate}

		tpl := template.Must(template.New("template.html").Funcs(funcMap).ParseFiles("template.html"))

		tpl.Execute(w, map[string]interface{}{
			"user":  user{Name: "a", Age: 19},
			"map":   map[string]string{"key1": "val1", "key2": "val2"},
			"slice": []string{"element1", "element2"},
			"bool":  "",
			"time":  time.Now(),
		})
	})

	http.ListenAndServe(":8085", r)
}
