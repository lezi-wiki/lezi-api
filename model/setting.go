package model

import (
	"encoding/gob"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Name string      `gorm:"not null;uniqueIndex"`
	Type SettingType `gorm:"not null;index"`
	Val  string      `gorm:"not null"`
}

type SettingType string

const (
	SettingTypeSystem SettingType = "system"
)

func init() {
	gob.Register(&Setting{})
}
