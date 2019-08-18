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

## 実習
### 基礎1
1_hello/hello.go


## ORM

`go get github.com/go-sql-driver/mysql`  
`go get github.com/jinzhu/gorm`