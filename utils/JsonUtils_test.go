package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestToJsonObject(t *testing.T) {
	file, _ := os.Open("../config/json")
	for {
		msg, err := bufio.NewReader(file).ReadBytes('#')
		if err != nil && io.EOF != err {
			fmt.Println("open file failed.")
		}
		if io.EOF == err {
			break
		}
		msgStr := strings.Replace(string(msg), "#", "", -1)
		msgStr = strings.Replace(msgStr, "\t", "", -1)
		msgStr = strings.Replace(msgStr, " ", "", -1)
		msgStr = StringUtils.ClearEnterWrap(msgStr)
		fmt.Println(msgStr)

		JsonParse.ToJsonObject(JsonParse.ToJsonStream(msgStr))
	}
	file.Close()
}
