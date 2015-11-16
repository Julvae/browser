package model

type RandomParams struct {
	ParamMap map[string]string
}

func (self *RandomParams) Add(key, value string) {
	if self.ParamMap == nil {
		self.ParamMap = make(map[string]string, 0)
	}
	self.ParamMap[key] = value
}

func (self *RandomParams) get(key string) (string, bool) {
	value, ok := self.ParamMap[key]
	return value, ok
}
