package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type task struct {
	// 構造体のフィールドに gorm.Model を含めると
	// ID, CreatedAt, UpdatedAt, DeletedAt を自動的に定義してくれる。
	gorm.Model
	// Model gorm.Model の省略形
	// task.ID で task.Model.ID にアクセスできる
	Description string `gorm:"not null"`
	IsComplete  bool   `gorm:"not null"`
	// Goでは構造体のフィールドの後に文字列を続けることでそのフィールドにタグと呼ばれるメタ情報を設定できる
	// タグはそれだけでは意味を持たないが、プログラム中から参照されることで処理の分岐などに使える
	// gorm ではテーブルの制約などの定義に利用する
	// 詳しくは http://gorm.io/ja_JP/docs/models.html 参照
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "gotodo"
	// ?parseTime=true をつけてやらないと、gormはMySQLのDATETIMEをtime.Time型ではなくstringにパースしようとしてしまう
	// また、loc=Asia%2FTokyo の指定がないとUTC時間に書き直されてしまう
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// 実行されたSQLを確認するためにログを取る
	file, err := os.OpenFile("./sql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))

	return db
}

func main() {
	// DBに接続
	db := gormConnect()
	// deferキーワードをつけることで、呼び出した関数（この場合はmain）の最後に
	// 必ずdeferをつけた関数呼び出しが実行される
	defer db.Close()

	// 構造体の定義に基づき自動的にテーブルを作成し、カラムを定義してくれる
	// ただし、既存のテーブルのカラム定義の変更は明示的にマイグレーション用のメソッドを呼び出さなければ実行できない
	// http://gorm.io/ja_JP/docs/migration.html 参照
	db.AutoMigrate(&task{})

	// Create
	task1 := task{Description: "test1", IsComplete: false}
	db.Create(&task1)

	task2 := task{Description: "test2", IsComplete: false}
	db.Create(&task2)

	// Read
	var tasks []task
	db.Find(&tasks)
	fmt.Printf("all tasks : %v\n", tasks)

	var task3 task
	// Descriptionを指定してfind
	db.Find(&task3, "Description = ?", "test1")
	fmt.Printf("a task : %v\n", task3)

	// Update
	db.Model(&task1).Update("IsComplete", true)
	fmt.Printf("after update: %v\n", task1)

	// Delete
	// 構造体のフィールドに DeletedAt が含まれている場合、
	// gorm は自動的に論理削除する
	// 論理削除されたレコードを参照する方法や物理削除する方法は
	// http://gorm.io/docs/delete.html#Soft-Delete 参照
	db.Delete(&task2)
	db.Find(&tasks)
	fmt.Printf("after delete : %v\n", tasks)

}
