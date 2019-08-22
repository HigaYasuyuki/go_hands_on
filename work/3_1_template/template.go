package main

import (
	// テンプレートライブラリにはtext/template と html/template の２種類がある
	// html/template は自動的にエスケープ処理を実施してくれる
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// template.ParseFiles(テンプレートファイル名) でテンプレートファイルを解析する
		// template.ParseFiles は Template型(解析済みテンプレート)とerrorを返すが、
		// errorを無視したいときはtemplate.Must()にtemplate.ParseFilesをそのまま渡してやると
		// Ttemplate型だけを受け取ることができる
		tpl := template.Must(template.ParseFiles("template.html"))
		// テンプレートの内容を出力
		// 第二引数はテンプレートに渡すデータ
		tpl.Execute(w, "Hello, template")
	})

	http.ListenAndServe(":8085", r)
}
