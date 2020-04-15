package cxml

import (
	"cjson/qson"
	"cxml/etree"
	"fmt"
)

func XmlToJSON(path string) *qson.JSONObject {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Println("xml文件加载失败....")
		return nil
	}
	ay := doc.Root().ChildElements()
	return analysisElement(ay)
}

func analysisElement(elements []*etree.Element) *qson.JSONObject {
	obj := qson.NewJSONObject()
	for _, v := range elements {
		_t := v.ChildElements()
		if len(_t) > 0 {
			obj.PutJSONObject(v.Tag, analysisElement(_t))
		} else {
			obj.PutString(v.Tag, v.Text())
		}
	}
	return obj
}
