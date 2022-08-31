package model

type ApiData struct {
	Code int         `json:"code" xml:"code"`
	Msg  string      `json:"message" xml:"msg"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}
