package main

import (
	"fmt"
	"net/http"
)

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
	// Fprintfは出力先を明示的に指定できる関数
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	// パターンとそれにマッチしたリクエストを処理する関数を登録する
	http.HandleFunc("/", handler)
	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go")
	})

	// httpサーバを立ち上げる
	// 第一引数にリクエストを受け付けるアドレス、第二引数にルーティング処理を行うルーターを指定する
	// 第二引数がnilの場合、デフォルトのルーターが使用される
	http.ListenAndServe(":8085", nil)
	// 下記と同じ
	// server := http.Server{
	// 	Addr:    ":8085",
	// 	Handler: nil,
	// }
	// server.ListenAndServe()
}
