package dto

import (
	"encoding/json"
	"github.com/lezi-wiki/lezi-api/model"
)

type TextJsonDTO struct {
	Text      string `json:"text"`
	Namespace string `json:"ns"`
	Speaker   string `json:"speaker"`
}

func (t TextJsonDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

func UnmarshalTextJsonDTO(data []byte) (TextJsonDTO, error) {
	var r TextJsonDTO
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalTextJsonDTOs(data []byte) ([]TextJsonDTO, error) {
	var r []TextJsonDTO
	err := json.Unmarshal(data, &r)
	return r, err
}

func BuildTextJsonDTO(text TextJsonDTO) model.Text {
	return model.Text{
		Namespace: text.Namespace,
		Speaker:   text.Speaker,
		Text:      text.Text,
	}
}
