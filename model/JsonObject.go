package model

type JsonObject struct {
	JObject map[string]string
	JArray  map[string]*JsonObject
}
