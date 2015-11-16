package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type stringUtils struct {
}

var StringUtils stringUtils = stringUtils{}

func (self *stringUtils) ClearEnterWrap(str string) string {
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func (self *stringUtils) ClearBlank(str string) string {
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	return str
}

func (self *stringUtils) Test() {

}

func (self *stringUtils) GetValidParams(pm []string) []string {
	if pm == nil {
		fmt.Println("over the exe.")
	}
	fmt.Println(pm)
	result := make([]string, 0)
	for _, value := range pm {
		fmt.Println(value)
		isTrue, _ := regexp.Match("\\s+", []byte(value))
		if !isTrue {
			result = append(result, value)
		}
	}
	return result
}

func (self *stringUtils) ToMap(str string) map[string][]string {
	params := regexp.MustCompile("&+").Split(str, -1)
	result := make(map[string][]string)
	for _, v := range params {
		if strings.Contains(v, "=") {
			keyValue := regexp.MustCompile("=").Split(v, -1)
			result[keyValue[0]] = append(result[keyValue[0]], keyValue[1])
		}
	}
	return result
}
