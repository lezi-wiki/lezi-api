package model

import (
	"encoding/gob"
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Namespace string  `json:"ns" xml:"namespace" bson:"namespace" gorm:"not null;index"`
	Speaker   string  `json:"speaker" xml:"speaker" bson:"speaker" gorm:"not null;index"`
	Text      string  `json:"text" xml:"text" bson:"text" gorm:"not null;size:512"`
	Context   *string `json:"context,omitempty" xml:"context,omitempty" bson:"context,omitempty" gorm:"size:512"`
}

func init() {
	gob.Register(&Text{})
}
