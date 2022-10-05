package model

import (
	"encoding/gob"
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Namespace string  `gorm:"not null;index"`
	Speaker   string  `gorm:"not null;index"`
	Text      string  `gorm:"not null;size:512"`
	Context   *string `gorm:"size:512"`
}

func init() {
	gob.Register(&Text{})
}
