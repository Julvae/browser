package filemanger

import (
	"fmt"
	"io"
	"testing"
)

func TestReadString(t *testing.T) {
	file := FileTools.Open("../config/json")
	fmt.Println("../config/json")
	FileTools.OpenReader(file)
	for {
		msg, err := FileTools.ReadString('#')
		fmt.Println(msg)
		if io.EOF == err {
			break
		}
	}
	file.Close()
}

func TestReadLine(t *testing.T) {
	file := FileTools.Open("../config/test")
	fmt.Println("../config/test")
	FileTools.OpenReader(file)
	for {
		msg, err := FileTools.ReadLine()
		fmt.Println(msg)
		if io.EOF == err {
			break
		}
	}
	file.Close()
}
