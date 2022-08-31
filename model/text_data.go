package model

type TextData struct {
	Namespace string `json:"ns" xml:"namespace"`
	Speaker   string `json:"speaker" xml:"speaker"`
	Text      string `json:"text" xml:"text"`
}
