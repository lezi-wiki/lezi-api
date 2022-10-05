package remote

import (
	"encoding/json"
	"errors"
	"github.com/lezi-wiki/lezi-api/pkg/http"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer/dto"
)

var Endpoint string

const (
	ErrNotValid = "json is not legal"
)

func GetDataFromGitHub() ([]dto.TextJsonDTO, error) {
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

	data, err := dto.UnmarshalTextJsonDTOs(raw)
	if err != nil {
		log.Log().Errorf("Error when get data from GitHub, %s", err.Error())
		return nil, err
	}

	return data, nil
}
