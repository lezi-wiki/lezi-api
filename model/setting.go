package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model `json:"-" xml:"-" bson:"-"`
	Name       string      `json:"name" xml:"name" bson:"name" gorm:"not null;uniqueIndex"`
	Type       SettingType `json:"type" xml:"type" bson:"type" gorm:"not null;index"`
	Val        string      `json:"val" xml:"val" bson:"val" gorm:"not null"`
}

type SettingType string

const (
	SettingTypeSystem SettingType = "system"
)
