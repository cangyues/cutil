package qson

import (
	json "cjson/simplejson"
	"container/list"
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

func ParseJSONArray(starry string) *JSONArray {
	array := new(JSONArray)
	array.data = list.New()
	j, e := json.NewJson([]byte(starry))
	if e != nil {
		log.Fatal(e)
	}
	a, er := j.Array()
	if er != nil {
		log.Fatal(er)
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

func (js *JSONArray) ToString() string {
	a := js.ToArray()
	j := json.New()
	j.Set("_ks", a)
	s, e := j.Get("_ks").MarshalJSON()
	if e != nil {
		log.Fatal(e)
	}
	return string(s)
}

func (js *JSONArray) IsEmpty() bool {
	return js.data.Len() == 0
}
