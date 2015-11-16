package systeminfo

import (
	"runtime"
)

var NEWLINE = "\n"
var OSTYPE = "windows"
var DELIM = '\n'

func init() {
	if runtime.GOOS == "windows" {
		NEWLINE = "\r\n"
		OSTYPE = "windows"
		DELIM = '\n'
	} else if runtime.GOOS == "darwin" {
		NEWLINE = "\r"
		OSTYPE = "darwin"
		DELIM = '\r'
	} else {
		NEWLINE = "\n"
		OSTYPE = "linux"
		DELIM = '\n'
	}
}
