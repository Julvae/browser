package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestClearEnterWrap(t *testing.T) {
	file, _ := os.Open("../config/enterwrap")
	msg, _ := ioutil.ReadAll(file)
	fmt.Println(string(msg))
	msgStr := StringUtils.ClearEnterWrap(string(msg))
	fmt.Println(msgStr)
	file.Close()

	filea, _ := os.Open("../config/json")
	msga, _ := ioutil.ReadAll(filea)
	fmt.Println(string(msga))
	msgStr = StringUtils.ClearEnterWrap(string(msga))
	fmt.Println(msgStr)
	filea.Close()

	fileb, _ := os.Open("../config/json")
	msgb, _ := bufio.NewReader(fileb).ReadBytes('#')
	fmt.Println(string(msgb))
	msgStr = StringUtils.ClearEnterWrap(string(msgb))
	fmt.Println(msgStr)
	fileb.Close()
}
