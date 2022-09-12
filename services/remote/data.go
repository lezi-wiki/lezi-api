package remote

import (
	"encoding/json"
	"errors"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/http"
	"github.com/lezi-wiki/lezi-api/pkg/log"
)

var Endpoint string

const (
	ErrNotValid = "json is not legal"
)

func GetDataFromGitHub() ([]model.Text, error) {
	raw, err := http.Get(Endpoint)
	if err != nil {
		log.Log().Errorf("Error when get data from GitHub, %s", err.Error())
		return nil, err
	}

	isValid := json.Valid(raw)
	if !isValid {
		log.Log().Errorf("Error when get data from GitHub, %s", ErrNotValid)
		return nil, errors.New(ErrNotValid)
	}

	var data []model.Text
	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Log().Errorf("Error when get data from GitHub, %s", err.Error())
		return nil, err
	}

	return data, nil
}
