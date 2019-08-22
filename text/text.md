# Go言語ハンズオン

## 今回の勉強会の目的
- Go言語に触れてみる
- Goを使ったwebアプリ開発の雰囲気を掴む

## Goってどんな言語?
- 文法がシンプル
- 静的型付け言語（型推論あり）
- バイナリだけ置けばデプロイ完了
- 動作が高速
- 並行処理が簡単に書ける

## 環境構築
- repository clone後、`docker-compose up -d`
- `docker-compose exec app bash` で作業用コンテナにログイン  
  windowsでgit bashを使っている人は `winpty docker-compose exec app bash`
- コンテナにログイン後、`go run 1_hello/hello.go` というコマンドを実行して Hello world と出力されるか確認
- workディレクトリがコンテナと同期されている
- http://localhost:8085 で コンテナの中のwebアプリにアクセス（今はまだアプリが未起動）
- http://localhost:8086 で コンテナの中のphpmyadminにアクセス

## プログラムの実行方法
- `go build -o hello hello.go` でビルドし、`./hello` で実行するのが基本
- `go run hello.go` でビルドコマンドを省略して実行できる

## 文法基礎
basics.md 参照

## 練習1
- goでFizz Buzzを書いてみよう  
- 1から100までの数字について、3で割り切れるときは`Fizz`、5で割り切れるときは`Buzz`、  
どちらでも割り切れるときは`FizzBuzz`、どちらでも割り切れないときはその数字を出力するプログラム。

出力例
```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
...
```


## Webアプリ（ToDoリスト）を作る
### net/http を使ってwebサーバを提供する
- Go言語組み込みのhttpサーバ機能
- 基礎的なルーティング
- パラメータを使用した高度なルーティング
`go get github.com/gorilla/mux`

### テンプレート
- Go言語組み込みのテンプレート機能
  - テンプレートの解析、実行
- アクション
- レイアウト

## 練習2
先程のFizzBuzzをテンプレートを使って出力してみましょう。  
フォームから入力された数値を受け取り、  
1 ~ その数値までをFizzBuzzした結果を  
| ID | Fizz | Buzz | FizzBuzz |
|----|------|------|----------|
|3|○|||
このようにテーブルとして出力します。
3_4_practice ディレクトリに雛形を用意したのでそれをもとに作成してみてください。


### DB
- ormライブラリgormの利用  
http://gorm.io/ja_JP/docs/

`go get github.com/go-sql-driver/mysql`
`go get github.com/jinzhu/gorm`

## 実際にtodoリストを作ってみよう!
- 5_todolist ディレクトリの中に雛形を作っておいたので、CRUD機能を実装してみましょう