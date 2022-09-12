package text

import (
	"errors"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/text"
)

func GetTextByNamespace(ns string) ([]model.Text, error) {
	var arr []model.Text
	for _, v := range text.Data {
		if v.Namespace == ns {
			arr = append(arr, v)
		}
	}

	if len(arr) == 0 {
		return nil, errors.New("no enough data")
	}

	return arr, nil
}

func GetTextBySpeaker(speaker string) ([]model.Text, error) {
	var arr []model.Text
	for _, v := range text.Data {
		if v.Speaker == speaker {
			arr = append(arr, v)
		}
	}

	if len(arr) == 0 {
		return nil, errors.New("no enough data")
	}

	return arr, nil
}
