package filemanger

import (
	"Regexp/systeminfo"
	"Regexp/utils"
	"bufio"
	"io"
	"os"
	"strings"
)

type FileUtils struct {
	reader *bufio.Reader
}

var FileTools FileUtils = FileUtils{}

func (self *FileUtils) Open(fileName string) (file *os.File) {
	fileObject, err := os.Open(fileName)
	if err != nil {
		panic("can't open the file " + fileName + ":" + err.Error())
	}
	return fileObject
}

func (self *FileUtils) Reset(file *os.File) {
	file.Seek(0, 0)
}
func (self *FileUtils) Close() {
	self.reader = nil
}
func (self *FileUtils) OpenReader(file *os.File) {
	self.reader = bufio.NewReader(file)
}
func (self *FileUtils) ReadLine() (string, error) {
	msg, err := self.reader.ReadBytes(byte(systeminfo.DELIM))
	if err != nil && err != io.EOF {
		panic("some error appear when reading file:" + err.Error())
	}
	return utils.StringUtils.ClearEnterWrap(string(msg)), err
}

func (self *FileUtils) ReadString(delim byte) (string, error) {
	msg, err := self.reader.ReadBytes(delim)
	if err != nil && err != io.EOF {
		panic("some error appear when reading file:" + err.Error())
	}
	msgStr := strings.Replace(string(msg), string(delim), "", -1)
	msgStr = strings.Replace(string(msgStr), "\t", "", -1)
	msgStr = strings.Replace(string(msgStr), " ", "", -1)
	return utils.StringUtils.ClearEnterWrap(msgStr), err
}
