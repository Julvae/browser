package tools

import (
	"Regexp/ParamsParse"
	"Regexp/httpClient"
	"Regexp/systeminfo"
	"Regexp/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	//"strings"
	//"time"
)

var CmdConfigMap = make(map[string]string)
var CmdConsole = make(map[string]string)

func ToCmdMap(cmdArgs []string) {
	CmdConsole = make(map[string]string)
	utils.MapUtils.Copy(CmdConfigMap, CmdConsole)
	var argsName []string = []string{"cmd", "url", "times", "thread"}
	for index, value := range cmdArgs {
		CmdConsole[argsName[index]] = value
	}
}
func GetCmdArgs() []string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes(byte(systeminfo.DELIM))
	args := regexp.MustCompile("\\s+").Split(string(input), -1)
	return utils.StringUtils.GetValidParams(args)
}

func Help() {
	fmt.Println("*********************help*****************************")
	file, err := os.Open("./config/help")
	if err != nil {
		return
	}
	msg, _ := ioutil.ReadAll(file)
	fmt.Println(string(msg))
	fmt.Println("*********************help*****************************")
}

func LoadConfig() {
	CmdConfigMap = ParamsParse.UrlUtils.GetCmdConfig("./config/CmdConfig")
}

func initCmdMap(mp map[string]string) bool {
	_, ok := mp["url"]
	if !ok {
		fmt.Println("you should give a url for http connect.")
		return false
	}
	_, ok1 := mp["cmd"]
	if !ok1 {
		mp["cmd"] = "wget"
	}
	_, ok2 := mp["thread"]
	if !ok2 {
		mp["thread"] = "1"
	}
	_, ok3 := mp["times"]
	if !ok3 {
		mp["times"] = "1"
	}
	return true
}

func AutoRun() {
	if initCmdMap(CmdConfigMap) {
		threads, _ := strconv.Atoi(CmdConfigMap["thread"])
		times, _ := strconv.Atoi(CmdConfigMap["times"])
		handle := httpClient.HandlerMap[CmdConfigMap["cmd"]]
		URL := CmdConfigMap["url"]
		for i := 0; i < threads; i++ {
			go httpClient.HttpUtils.Run(handle, URL, times)
			//time.Sleep(time.Second)
		}
	}
}

func HandleRun() {
	if initCmdMap(CmdConsole) {
		if _, ok := CmdConsole["url"]; !ok {
			fmt.Println("you should give a url for http connect.")
			return
		}
		threads, _ := strconv.Atoi(CmdConsole["thread"])
		times, _ := strconv.Atoi(CmdConsole["times"])
		handle := httpClient.HandlerMap[CmdConsole["cmd"]]
		URL := CmdConfigMap["url"]
		for i := 0; i < threads; i++ {
			go httpClient.HttpUtils.Run(handle, URL, times)
		}
	}
}

func Run() {
	LoadConfig()
	for {
		fmt.Print("apollo>>")
		args := GetCmdArgs()
		fmt.Println("args[0]=", args[0])
		if len(args) == 0 {
			Help()
			continue
		}
		if args[0] == "refresh" {
			LoadConfig()
			continue
		}
		if args[0] == "quit" {
			break
		}

		if args[0] == "run" {
			AutoRun()
		} else if args[0] == "wget" || args[0] == "wpost" {
			ToCmdMap(args)
			HandleRun()
		} else {
			Help()
		}
	}
}

func IsAuto() bool {
	elem, ok := CmdConfigMap["auto"]
	return ok && elem == "true"
}
