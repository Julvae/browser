package httpClient

import (
	"Regexp/ParamsParse"
	"Regexp/systeminfo"
	"Regexp/utils"
	"fmt"
	"github.com/cheggaaa/pb"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var URLPARAM = ""
var URLRANDOMPARAM = ""

type httpUtils struct {
}
type Handler func(URL string) string

var HttpUtils httpUtils = httpUtils{}
var HandlerMap = map[string]Handler{"wget": HttpUtils.Get, "wpost": HttpUtils.Post}

func (self *httpUtils) Get(URL string) string {
	//fmt.Println(URL)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
		return "error"
	}
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return (string(body))
}

func (self *httpUtils) Post(URL string) string {
	params := utils.StringUtils.ToMap(strings.Split(URL, "?")[1])
	values := make(url.Values, 0)
	for key, value := range params {
		for _, elem := range value {
			values.Add(key, elem)
		}
	}
	response, err := http.PostForm(URL, values)
	if err != nil {
		fmt.Println(err.Error())
		return "error"
	}
	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return (string(body))
}

func (self *httpUtils) Run(handle Handler, URL string, times int) {
	fileName := utils.RandomUtils.GetFileName()
	fileObject, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("create file faile.")
	}
	bar := pb.StartNew(times)
	for i := 0; i < times; i++ {
		bar.Increment()
		URLLink := URL
		if !strings.Contains(URL, "?") {
			URLPARAM = ParamsParse.UrlUtils.GetUrlLinkByParams("./config/myconfig")
			URLRANDOMPARAM = ParamsParse.UrlUtils.GetUrlLinkByRandomParams("./config/RandomParamsConfig")
			URLLink = URL + "?" + URLPARAM + URLRANDOMPARAM
		}
		response := handle(URLLink)
		fileObject.Write([]byte(response + systeminfo.NEWLINE))
	}
	bar.FinishPrint("The End!")
	fileObject.Close()
}
