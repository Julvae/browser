package utils

import (
	"Regexp/model"
)

type MapExtend struct {
}

var MapUtils MapExtend = MapExtend{}

func (self *MapExtend) KeySet(m map[string][]*model.UrlParam) (keys []string) {
	keys = make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func (self *MapExtend) CheckAllKey(m map[string][]*model.UrlParam, expect string) bool {
	keys := self.KeySet(m)
	length := 0
	for _, key := range keys {
		if key == expect {
			length++
		}
	}
	return length == len(keys) || length == 0
}

func (self *MapExtend) Copy(src, des map[string]string) {
	for k, v := range src {
		des[src[k]] = v
	}

}
