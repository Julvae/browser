package ParamsParse

import (
	"Regexp/filemanger"
	"Regexp/model"
	"Regexp/utils"
	"io/ioutil"
	"regexp"
	"strings"
)

type UrlParamUtils struct {
}

var ARGSIGN = "args"
var VALUESIGN = "@"
var RANDOMTYPE = "string"
var UrlUtils = UrlParamUtils{}

func (self *UrlParamUtils) jointParamPair(key, value string) string {
	return key + "=" + value + "&"
}

func (self *UrlParamUtils) GetUrlLinkByParams(fileName string) string {
	flieParase := filemanger.CreateFileParase(fileName)
	urlParams := flieParase.GetJsonObjectArray()
	return self.genUrlLink(urlParams)
}

func (self *UrlParamUtils) genUrlLink(urlParams *model.UrlParams) string {
	urlLink := ""
	for _, value := range urlParams.Params {
		urlLink = urlLink + self.genLinkByObject(value)
	}
	return urlLink
}

//object is like {"a":"1","args":[...]}
func (self *UrlParamUtils) genLinkByObject(urlParam *model.UrlParam) string {
	urlLink := ""
	for key, value := range urlParam.JObject {
		//fmt.Println("key=", key)
		randomIndex := int(utils.RandomUtils.RandomDigit()) % len(value)
		//fmt.Println("randomIndex=", randomIndex)
		//fmt.Println("len(value)=", len(value))
		if key == ARGSIGN {
			urlLink = urlLink + self.genLinkByObject(value[randomIndex])
		} else {
			urlLink = urlLink + self.jointParamPair(key, self.getFormatValueKeys(value[randomIndex]))
		}
	}
	return urlLink
}

func (self *UrlParamUtils) getFormatValueKeys(urlParam *model.UrlParam) string {
	for valueKey, _ := range urlParam.JObject {
		valueKey = utils.StringUtils.ClearBlank(valueKey)
		valueKeys := regexp.MustCompile(",").Split(valueKey, -1)
		randomIndex := int(utils.RandomUtils.RandomDigit()) % len(valueKeys)
		//fmt.Println("len(valueKeys)=", len(valueKeys))
		return valueKeys[randomIndex]
	}
	return ""
}

func (self *UrlParamUtils) GetUrlLinkByRandomParams(fileName string) string {
	flieParase := filemanger.CreateFileParase(fileName)
	urlParams := flieParase.GetJsonObjectArray()
	return self.genRandomLink(urlParams)
}

func (self *UrlParamUtils) genRandomLink(urlParams *model.UrlParams) string {
	urlLink := ""
	for _, value := range urlParams.Params {
		urlLink = urlLink + self.genOneLink(value)
	}
	return urlLink
}

func (self *UrlParamUtils) genOneLink(urlParam *model.UrlParam) string {
	urlLink := ""
	for key, value := range urlParam.JObject {
		nextKeyType := utils.MapUtils.KeySet(value[0].JObject)[0]
		length := utils.RandomUtils.RandomDigit() % 100
		if nextKeyType == "int" {
			urlLink += self.jointParamPair(key, utils.RandomUtils.RandomDigits(int(length)))
		} else if nextKeyType == "string" {
			urlLink += self.jointParamPair(key, utils.RandomUtils.RandomString(int(length)))
		} else if nextKeyType == "float" {
			urlLink += self.jointParamPair(key, utils.RandomUtils.RandomFloat(90))
		} else {
			urlLink += self.jointParamPair(key, utils.RandomUtils.RandomString(int(length)))
		}
	}
	return urlLink
}

func (self *UrlParamUtils) GetCmdConfig(filePath string) map[string]string {
	msg, _ := ioutil.ReadFile(filePath)
	//fmt.Println(msg)
	cmdMsg := utils.StringUtils.ClearBlank(string(msg))
	//fmt.Println(cmdMsg)
	cmds := strings.Split(cmdMsg, ",")
	//fmt.Println(cmds)
	cmdMap := make(map[string]string)
	for _, value := range cmds {
		if strings.Contains(value, "=") {
			//fmt.Println(value + "#")
			keyValue := strings.Split(value, "=")
			cmdMap[keyValue[0]] = keyValue[1]
			//fmt.Println("key=", keyValue[0], "#value=", keyValue[1])
		}
	}
	return cmdMap
}
