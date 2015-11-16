package filemanger

import (
	"Regexp/model"
	"Regexp/utils"
	//"fmt"
	"io"
	"os"
)

var JsonDelim = '#'

type FileParase struct {
	fileName   string
	fileObject *os.File
}

func CreateFileParase(name string) FileParase {
	return FileParase{name, FileTools.Open(name)}
}

func (self *FileParase) GetJsonObjectArray() *model.UrlParams {
	urlParams := &model.UrlParams{}
	FileTools.OpenReader(self.fileObject)
	for {
		msg, err := FileTools.ReadString(byte(JsonDelim))
		if len(msg) == 0 || io.EOF == err {
			break
		}
		//fmt.Println(msg)
		urlParam := utils.JsonParse.ToJsonObject(utils.JsonParse.ToJsonStream(msg))
		urlParams.Add(urlParam)
	}
	self.fileObject.Close()
	return urlParams
}
