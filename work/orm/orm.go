package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type task struct {
	gorm.Model
	Description string `gorm:"not null"`
	IsComplete  bool   `gorm:"not null"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "gotodo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	db.AutoMigrate(&task{})

	singleTask := task{Description: "test", IsComplete: false}
	db.Create(&singleTask)

	var tasks []task
	db.Find(&tasks)
	fmt.Println(tasks)
	defer db.Close()
}
