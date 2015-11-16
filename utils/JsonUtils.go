package utils

import (
	"Regexp/model"
	"encoding/json"
	//"fmt"
)

type JsonUtils struct {
}

type JsonMap map[string]interface{}

var JsonParse JsonUtils = JsonUtils{}

func (self *JsonUtils) ToJsonObject(jsonMap JsonMap) *model.UrlParam {
	urlParam := &model.UrlParam{}
	self.JsonToObject(jsonMap, urlParam)
	return urlParam
}

func (self *JsonUtils) JsonToObject(jsonMap JsonMap, urlParam *model.UrlParam) {
	for key, value := range jsonMap {
		switch value.(type) {
		case string:
			//fmt.Println("string =" + value.(string))
			mapValues := &model.UrlParam{}
			mapValues.Add(value.(string), nil)
			urlParam.Add(key, mapValues)
		case []interface{}:
			//fmt.Println("array =", value.([]interface{}))
			urlParam.AddList(key, self.JsonArrayToObjects(value.([]interface{})))
		default:
			panic(key + "is of a type I don't know how to handle")
		}
	}
}

func (self *JsonUtils) JsonArrayToObjects(array []interface{}) []*model.UrlParam {
	nextValues := make([]*model.UrlParam, 0)
	for _, value := range array {
		nextJsonMap := (value).(map[string]interface{})
		urlParam := &model.UrlParam{}
		self.JsonToObject(nextJsonMap, urlParam)
		nextValues = append(nextValues, urlParam)
	}
	return nextValues
}

func (self *JsonUtils) ToJsonStream(jsonStr string) JsonMap {
	var f interface{}
	err := json.Unmarshal([]byte(jsonStr), &f)
	if f == nil {
		panic("error jsonString:@" + err.Error())
	}
	return f.(map[string]interface{})
}
