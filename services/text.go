package services

import (
	"errors"
	"github.com/lezi-wiki/lezi-api/bootstrap"
	"github.com/lezi-wiki/lezi-api/model"
)

func GetTextByNamespace(ns string) ([]model.TextData, error) {
	var arr []model.TextData
	for _, v := range bootstrap.TextData {
		if v.Namespace == ns {
			arr = append(arr, v)
		}
	}

	if len(arr) == 0 {
		return nil, errors.New("no enough data")
	}

	return arr, nil
}
