package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// template.ParseFilesに使用するテンプレートファイルをすべて渡す
		t := template.Must(template.ParseFiles("base.html", "sample_content.html", "partial.html"))
		// t.Execute ではなく、ExecuteTemplate を実行
		// 第二引数で出力するテンプレート名を指定
		// 指定されたテンプレート以外は直接出力されることはない
		t.ExecuteTemplate(w, "layout", "Layout")
	})
	http.ListenAndServe(":8085", r)
}
