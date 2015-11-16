package model

type UrlParam struct {
	JObject map[string][]*UrlParam
}

func (self *UrlParam) Add(key string, urlParam *UrlParam) {
	if self.JObject == nil {
		self.JObject = make(map[string][]*UrlParam)
	}
	self.JObject[key] = append(self.JObject[key], urlParam)
}

func (self *UrlParam) AddList(key string, values []*UrlParam) {
	for _, value := range values {
		self.Add(key, value)
	}
}
