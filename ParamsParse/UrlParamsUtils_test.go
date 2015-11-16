package ParamsParse

import (
	"fmt"
	"testing"
)

func TestGetUrlLinkByParams(t *testing.T) {
	urlLink := UrlUtils.GetUrlLinkByParams("../config/myconfig")
	fmt.Println(urlLink)
}

func TestGetUrlLinkByRandomParams(t *testing.T) {
	urlLink := UrlUtils.GetUrlLinkByRandomParams("../config/RandomParamsConfig")
	fmt.Println(urlLink)
	urlLink = UrlUtils.GetUrlLinkByRandomParams("../config/RandomParamsConfig")
	fmt.Println(urlLink)
}

func TestGetCmdConfig(t *testing.T) {
	UrlUtils.GetCmdConfig("../config/CmdConfig")
}
