package vo

import (
	"encoding/json"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/hashids"
)

type TextVO struct {
	ID        string `json:"id" xml:"Id" bson:"id" yaml:"id"`
	Namespace string `json:"ns" xml:"Namespace" bson:"namespace" yaml:"namespace"`
	Speaker   string `json:"speaker" xml:"Speaker" bson:"speaker" yaml:"speaker"`
	Text      string `json:"text" xml:"Text" bson:"text" yaml:"text"`
}

func (t TextVO) MarshalTextVOJSON() ([]byte, error) {
	return json.Marshal(t)
}

func UnmarshalTextVO(data []byte) (TextVO, error) {
	var r TextVO
	err := json.Unmarshal(data, &r)
	return r, err
}

func BuildTextVO(text *model.Text) TextVO {
	return TextVO{
		ID:        hashids.HashIDEncode(text.ID, hashids.TypeText),
		Namespace: text.Namespace,
		Speaker:   text.Speaker,
		Text:      text.Text,
	}
}
