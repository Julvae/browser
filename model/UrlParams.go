package model

type UrlParams struct {
	Params []*UrlParam
}

func (self *UrlParams) Add(urlParam *UrlParam) {
	if self.Params == nil {
		self.Params = make([]*UrlParam, 0)
	}
	self.Params = append(self.Params, urlParam)
}
