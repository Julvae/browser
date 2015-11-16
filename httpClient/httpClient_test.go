package httpClient

import (
	"Regexp/ParamsParse"
	"fmt"
	"testing"
)

const (
	URLPREIX = "http://as.xiaojukeji.com/toggles?"
)

func TestGet(t *testing.T) {
	URLParam := ParamsParse.UrlUtils.GetUrlLinkByParams("../config/myconfig")
	URLRandomParam := ParamsParse.UrlUtils.GetUrlLinkByRandomParams("../config/RandomParamsConfig")
	URL := URLPREIX + URLParam + URLRandomParam
	fmt.Println(URL)
	HttpUtils.Get(URL)
}

func TestPost(t *testing.T) {
	URLParam := ParamsParse.UrlUtils.GetUrlLinkByParams("../config/myconfig")
	URLRandomParam := ParamsParse.UrlUtils.GetUrlLinkByRandomParams("../config/RandomParamsConfig")
	URL := URLPREIX + URLParam + URLRandomParam
	fmt.Println(URL)
	HttpUtils.Post(URL)
}

func TestRun(t *testing.T) {
	URLParam := ParamsParse.UrlUtils.GetUrlLinkByParams("../config/myconfig")
	URLRandomParam := ParamsParse.UrlUtils.GetUrlLinkByRandomParams("../config/RandomParamsConfig")
	URL := URLPREIX + URLParam + URLRandomParam
	fmt.Println(URL)
	HttpUtils.Run(HttpUtils.Get, URL, 20)
}
