package qson

import (
	json "cjson/simplejson"
	"log"
)

type JSONObject struct {
	data *json.Json
}

func NewJSONObject() *JSONObject {
	js := new(JSONObject)
	js.data = json.New()
	return js
}

func ParseJSONObject(jsons string) *JSONObject {
	js := new(JSONObject)
	var err error
	js.data, err = json.NewJson([]byte(jsons))
	if err != nil {
		log.Fatal(err)
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
	v, err := jb.data.Get(key).String()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (jb *JSONObject) GetInt(key string) int {
	v, err := jb.data.Get(key).Int()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (jb *JSONObject) GetInt64(key string) int64 {
	v, err := jb.data.Get(key).Int64()
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (jb *JSONObject) GetArray(key string) []interface{} {
	v, err := jb.data.Get(key).Array()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(e)
	}
	return string(b)
}

func (jb *JSONObject) Keys() []string {
	m, e := jb.data.Map()
	if e != nil {
		log.Fatal(e)
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
		log.Fatal(e)
	}
	for k, v := range m {
		exec(k, v)
	}
}

func (jb *JSONObject) ToMap() map[string]interface{} {
	v, _ := jb.data.Map()
	return v
}
