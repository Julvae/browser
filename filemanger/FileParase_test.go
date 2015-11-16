package filemanger

import (
	"testing"
)

func TestGetJsonObjectArray(t *testing.T) {
	fileParase := CreateFileParase("../config/json")
	urlParams := fileParase.GetJsonObjectArray()
	if urlParams == nil {
		t.Error("error")
	}
}
