package serializer

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Response struct {
	Code int         `json:"code" xml:"code"`
	Msg  string      `json:"message" xml:"msg"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (r *Response) Json() ([]byte, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (r *Response) Xml() ([]byte, error) {
	bytes, err := xml.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewResponseWithCode(code int) *Response {
	return &Response{
		Code: code,
		Msg:  http.StatusText(code),
	}
}

func NewResponseWithCodeAndData(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  http.StatusText(code),
		Data: data,
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: data,
	}
}
