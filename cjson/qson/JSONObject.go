package qson

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type JSONObject struct {
	data *Json
}

func NewJSONObject() *JSONObject {
	js := new(JSONObject)
	js.data = New()
	return js
}

func ObjectToString(data interface{}) string {
	bt, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return "{}"
	}
	return string(bt)
}

func ParseJSONObject(jsons string) *JSONObject {
	js := new(JSONObject)
	var err error
	js.data, err = NewJson([]byte(jsons))
	if err != nil {
		return NewJSONObject()
	}
	return js
}

func (jb *JSONObject) put(key string, value interface{}) *JSONObject {
	jb.data.Set(key, value)
	return jb
}

func (jb *JSONObject) PutString(key string, value string) *JSONObject {
	jb.put(key, value)
	return jb
}

func (jb *JSONObject) PutInt(key string, value int) *JSONObject {
	jb.put(key, value)
	return jb
}

func (jb *JSONObject) PutFloat(key string, value float64) *JSONObject {
	jb.put(key, value)
	return jb
}

func (jb *JSONObject) PutInt64(key string, value int64) *JSONObject {
	jb.put(key, value)
	return jb
}

func (jb *JSONObject) PutJSONObject(key string, value *JSONObject) *JSONObject {
	jb.put(key, value.data)
	return jb
}

func (jb *JSONObject) PutJSONArray(key string, value *JSONArray) *JSONObject {
	jb.put(key, value.ToArray())
	return jb
}

func (jb *JSONObject) GetString(key string) string {
	v := jb.GetInterface(key)
	switch v.(type) {
	case json.Number:
		t, _ := v.(json.Number)
		return string(t)
	case string:
		t, _ := v.(string)
		return t
	case int:
		t := v.(int)
		return strconv.Itoa(t)
	case int32:
		t := v.(int32)
		return strconv.Itoa(int(t))
	case int64:
		t := v.(int64)
		return strconv.Itoa(int(t))
	case float32:
		t := v.(float32)
		return strconv.FormatFloat(float64(t), 'E', -1, 32)
	case float64:
		t := v.(float64)
		return strconv.FormatFloat(t, 'E', -1, 64)
	default:
		log.Println("get string error！")
		return ""
	}
}

func (jb *JSONObject) GetFloat(key string) float64 {
	v := jb.GetInterface(key)
	switch v.(type) {
	case string:
		t, _ := v.(string)
		_t, err := strconv.ParseFloat(t, 64)
		if err != nil {
			log.Println(err.Error())
		}
		return _t
	case float64:
		t, _ := v.(float64)
		return t
	case float32:
		t, _ := v.(float32)
		return float64(t)
	}
	log.Println("get string error！")
	return 0
}

func (jb *JSONObject) GetInt(key string) int {
	v := jb.GetInterface(key)
	switch v.(type) {
	case json.Number:
		t, _ := v.(json.Number)
		_t, err := strconv.Atoi(string(t))
		if err != nil {
			log.Println(err.Error())
		}
		return _t
	case int:
		t, _ := v.(int)
		return t
	case string:
		t, _ := v.(string)
		_t, err := strconv.Atoi(t)
		if err != nil {
			log.Println(err.Error())
		}
		return _t
	case float32:
		t, _ := v.(float32)
		return int(t)
	case float64:
		t, _ := v.(float32)
		return int(t)
	default:
		log.Println("get int error!")
		return 0
	}
}

func (jb *JSONObject) GetInterface(key string) interface{} {
	return jb.data.Get(key).Interface()
}

func (jb *JSONObject) GetInt64(key string) int64 {
	v, err := jb.data.Get(key).Int64()
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func (jb *JSONObject) GetArray(key string) []interface{} {
	v, err := jb.data.Get(key).Array()
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func (jb *JSONObject) GetJSONArray(key string) *JSONArray {
	a := new(JSONArray)
	a.ArrayToJSONArray(jb.GetArray(key))
	return a
}

func (jb *JSONObject) GetJSONObject(key string) *JSONObject {
	v := jb.data.Get(key)
	_t := new(JSONObject)
	_t.data = v
	return _t
}

func (jb *JSONObject) ToString() string {
	b, e := jb.data.MarshalJSON()
	if e != nil {
		fmt.Println(e)
	}
	return string(b)
}

func (jb *JSONObject) Keys() []string {
	m, e := jb.data.Map()
	if e != nil {
		fmt.Println(e)
	}
	array := make([]string, len(m))
	index := 0
	for k, _ := range m {
		array[index] = k
		index++
	}
	return array
}

func (jb *JSONObject) Each(exec func(key string, value interface{})) {
	m, e := jb.data.Map()
	if e != nil {
		fmt.Println(e)
	}
	for k, v := range m {
		exec(k, v)
	}
}

func (jb *JSONObject) ToMap() map[string]interface{} {
	v, _ := jb.data.Map()
	return v
}
