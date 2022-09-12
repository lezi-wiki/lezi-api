package serializer

import "net/http"

func NewErrorResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func NotFoundResponse() *Response {
	return &Response{
		Code: http.StatusNotFound,
		Msg:  http.StatusText(http.StatusNotFound),
	}
}
