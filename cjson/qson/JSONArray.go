package qson

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
)

type JSONArray struct {
	data *list.List
}

func NewJSONArray() *JSONArray {
	array := new(JSONArray)
	array.data = list.New()
	return array
}

func ArrayToString(data interface{}) string {
	bt, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return "[]"
	}
	return string(bt)
}

func ParseJSONArray(starry string) *JSONArray {
	array := new(JSONArray)
	array.data = list.New()
	j, e := NewJson([]byte(starry))
	if e != nil {
		fmt.Println(e)
	}
	a, er := j.Array()
	if er != nil {
		fmt.Println(er)
	}
	array.ArrayToJSONArray(a)
	return array
}

func (js *JSONArray) Add(value interface{}) *JSONArray {
	switch value.(type) {
	case *JSONObject:
		v, _ := value.(*JSONObject)
		js.data.PushBack(v.data)
		break
	case *JSONArray:
		v, _ := value.(*JSONArray)
		js.data.PushBack(v.ToArray())
		break
	default:
		js.data.PushBack(value)
	}
	return js
}

func (js *JSONArray) ToArray() []interface{} {
	array := make([]interface{}, js.data.Len())
	js.Each(func(index int, value interface{}) {
		array[index] = value
	})
	return array
}

func (js *JSONArray) Each(exec func(index int, value interface{})) {
	y := 0
	for i := js.data.Front(); i != nil; i = i.Next() {
		exec(y, i.Value)
		y++
	}
}

func (js *JSONArray) ArrayToJSONArray(array []interface{}) {
	js.data = list.New()
	for _, v := range array {
		js.data.PushBack(v)
	}
}

func (js *JSONArray) GetRow() *JSONObject {
	i := js.data.Front()
	if i == nil {
		return nil
	}
	v, e := i.Value.(map[string]interface{})
	if !e {
		v, _ := i.Value.(*Json)
		json := new(JSONObject)
		json.data = v
		return json
	}
	json := new(JSONObject)
	json.data = &Json{
		data: v,
	}
	return json
}

func (js *JSONArray) ToString() string {
	a := js.ToArray()
	j := New()
	j.Set("_ks", a)
	s, e := j.Get("_ks").MarshalJSON()
	if e != nil {
		fmt.Println(e)
	}
	return string(s)
}

func (js *JSONArray) IsEmpty() bool {
	return js.data.Len() == 0
}

func (js *JSONArray) Size() int{
	return js.data.Len()
}
