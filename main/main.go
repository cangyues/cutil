package main

import (
	"cjson/qson"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/task_mange?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf(err.Error())
	}
	db.LogMode(true)
	db.SingularTable(true)

	user := testTable{Name: "cangyue", Age: 24}
	db.Create(&user)

	var testTables []testTable

	db.Find(&testTables)

	fmt.Println(qson.ArrayToString(testTables))

	defer db.Close()
}

type testTable struct {
	id   int64
	Name string
	Age  int
}

func (t testTable) TableName() string {
	return "test_table"
}
