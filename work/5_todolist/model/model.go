package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// gormがmysqlに接続するために必要
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Task タスク
type Task struct {
	gorm.Model
	Description string `gorm:"not null"`
	IsComplete  bool   `gorm:"not null"`
}

var db *gorm.DB

// Init DB接続初期化
func Init() {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "gotodo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	var err error
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	file, err := os.OpenFile("./sql.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))

	db.AutoMigrate(&Task{})
}

// CloseDB DB接続のclose処理
func CloseDB() {
	db.Close()
}

// GetTaskList タスク一覧取得
func GetTaskList() []Task {
	var tasks []Task
	db.Find(&tasks)

	return tasks
}
