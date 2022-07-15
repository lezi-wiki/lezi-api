package http

import (
	"errors"
	"github.com/idoubi/goz"
)

const (
	ErrHttp = "http error"
)

func Get(url string, opt ...goz.Options) ([]byte, error) {
	req, err := client.Get(url, opt...)
	if err != nil {
		return nil, err
	}

	code := req.GetStatusCode()
	if code != 200 {
		return nil, errors.New(ErrHttp)
	}

	data, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Post(url string, data interface{}, opts ...goz.Options) ([]byte, error) {
	newOpt := []goz.Options{
		{
			JSON: data,
		},
	}
	for _, option := range opts {
		newOpt = append(newOpt, option)
	}

	req, err := client.Post(url, newOpt...)
	if err != nil {
		return nil, err
	}

	resp, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
