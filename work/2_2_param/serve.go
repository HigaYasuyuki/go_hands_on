package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// パラメーターの受け取り
	r.HandleFunc("/id/{id}", func(w http.ResponseWriter, r *http.Request) {
		// パラメータを取り出す
		id := mux.Vars(r)["id"]
		fmt.Fprintf(w, "id: %v", id)
	})

	// 正規表現の使用
	// /reg/aaa などとすると not found
	r.HandleFunc("/reg/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		fmt.Fprintf(w, "id: %v", id)
	})

	// 任意のパラメータ
	r.HandleFunc("/opt/{id:[0-9]*?}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		fmt.Fprintf(w, "id: %v", id)
	})

	// 階層化パラメータ
	r.HandleFunc("/nest/{category}/{id}", func(w http.ResponseWriter, r *http.Request) {
		category := mux.Vars(r)["category"]
		id := mux.Vars(r)["id"]
		fmt.Fprintf(w, "category: %v, id: %v", category, id)
	})

	// mux.NewRouter()で作成したルーターをListenAndServeにわたす
	http.ListenAndServe(":8085", r)
}
