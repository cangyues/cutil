package main

import (
	"cjson/qson"
	"fmt"
)

func main() {

	jxson := qson.NewJSONObject()

	jxson.PutString("sex", "20")
	fmt.Println(jxson.GetInt("sex"))

}

type TE struct {
	id   int64
	Name string
	Age  int
}
