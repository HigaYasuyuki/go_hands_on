# 文法基礎
## 変数
```go
// 通常の宣言
var x int

// 型名省略可能
var x, y = 1, "string"

// 暗黙的な型宣言
x := 1

// 定数
const X = 1
```

## 関数
```go
// 構文
func 関数名(変数名 型名, 変数名 型名) 返り値型名 { }

// 大文字で始まる関数はパッケージ外からアクセス可能
func Add(x int, y int) int { }

// 小文字で始まる関数はパッケージ外からアクセス不可
func add(x int, y int) int { }

// 引数の型が同じ場合はまとめる
func add(x, y int) int { }

// 関数値
f := func(x, y int) int { }
f(1, 2)
```

## 基本制御構文
```go
// 繰り返し
for i := 0 ; i < 10 ; i++ { }

// while
// while自体は存在しない
for i < 10 { }

// 無限ループ
for { }

// 分岐
if x < 0 { }

// ifのショートステートメント
// 条件の前に文を書ける
if v := add(x, y); v < 10 { }

// switch 
switch v := add(x, y); v {
case 1:
  // 処理 (breakは必要無い)
default:
  // 処理
}

// 条件の無いswitch
v := add(x, y)
switch {
case v < 10:
  // 処理
}
```

## ポインタ
```go
// ポインタ
var p *int

// &演算子 (ポインタを引き出す)
i := 42
p = &i

// *演算子 (ポインタの指す変数を示す)
fmt.Println(*p) // 42
*p = 21 // ポインタpを通して変数iに代入
```

## 構造体
goにはclassがないので、構造体をclassのように使う
```go
// 構造体の定義
type Point struct {
  X int
  Y int
}

// 構造体の初期化
p := Point{1, 2}

// フィールドを指定して初期化
p := Point{X: 1, Y: 2}

// 省略して初期化 (ゼロ値が入る)
p := Point{X: 1}
p := Point{}

// ポインタを返して初期化
v := &Point{1, 2}

// フィールドへアクセス
p.X = 10

// ポインタを通してアクセス
v := &p
(*v).X = 15 
v.X = 5 // *演算子は省略可能
```

## メソッド
```go
// Point型を定義
type Point struct {
  X, Y int
}

// Point型にメソッドを定義
func (p Point) Add() int {
  return p.X + p.Y
}

// メソッド呼び出し
p := Point{1, 2}
p.Add()

// ポインタレシーバ (レシーバ自身を更新出来る)
func (p *Point) Up() {
  p.X += 1
}
```

## 配列
動的なサイズ変更不可
```go
// 宣言
var a [3]int

// 初期化
a := [3]int{1,2,3}

// アクセス
a[0] = 1
```

## スライス
```go
// 宣言
var s []int

// 既存の配列からスライスを作る
var s []int = a[1:4] // 1番目から3番目の要素で作る

// スライスの省略
var s[]int = a[:4] // 0番目から3番目
var s[]int = a[1:] // 1番目から最後尾まで
var s[]int = a[:] // 0番目から最後尾まで

// 新規の配列からスライスのみを取得する
s := []int{1, 2, 3}

// 長さ (スライスの要素数) の取得
len(s)

// 容量 (スライスの最初の要素から数えて、元配列の要素数) の取得
cap(s)

// 動的サイズの配列の作成
s := make([]int, 長さ, 容量)

// 多次元スライス
board := [][]string{
  []string{"_", "_", "_"},
  []string{"_", "_", "_"},
  []string{"_", "_", "_"},
}
board[0][0] = "X"

// 要素の追加 (長さは追加した数だけ増え、容量は足りない時に多めに自動確保する)
append(s, 追加したい値, 追加したい値...)

// リストの走査 (i はインデックス, v は値のコピー)
for i, v := range s { }
for _, v := range s { } // インデックスの破棄
for i := range s { } // 値の破棄
```

## マップ
連想配列
```go
// 宣言
m := make(map[string]int)

// リテラル
var m = map[string]int{
  "one": 1,
  "two": 2,
}

// アクセス
m["one"] = 1

// 要素の削除
delete(m, key)

// キーの存在確認
elem, ok := m[key] // 存在すれば ok は true, しなければ false
```

## インターフェース
```go
// 型定義
type Point struct {
  X, Y int
}

// インタフェース定義
type Calculator interface {
  Add(int, int) int
}

// インタフェース実装 (implementsキーワードは必要無い)
func (p Point) Add(x, y int) int {
  return p.X + p.Y
}

// 代入
var c Calculator
p := Point{1, 2}
c = &p
```

## エラー
```go
// error型の定義
type error interface {
  Error()  string
}

// エラーチェック (nilなら成功, nilでなければ失敗)
i, err := f()
if err != nil {
  fmt.Printf("%v¥n", err)
  return
}
```