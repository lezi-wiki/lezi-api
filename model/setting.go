package model

import (
	"encoding/gob"
	"gorm.io/gorm"
)

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

func init() {
	gob.Register(&Setting{})
}

type SettingService interface {
	Get(name string, settingType SettingType) (Setting, error)
	Set(name string, settingType SettingType, val string) error
	Delete(name string, settingType SettingType) error
	List() ([]Setting, error)
	ListType(settingType SettingType) ([]Setting, error)
}

type SettingServiceImpl struct {
	db *gorm.DB
}

func NewSettingService(db *gorm.DB) SettingService {
	return &SettingServiceImpl{db: db}
}

func (s *SettingServiceImpl) Get(name string, settingType SettingType) (Setting, error) {
	var setting Setting
	err := s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).First(&setting).Error
	return setting, err
}

func (s *SettingServiceImpl) Set(name string, settingType SettingType, val string) error {
	return s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).Updates(&Setting{
		Val: val,
	}).Error
}

func (s *SettingServiceImpl) Delete(name string, settingType SettingType) error {
	return s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).Delete(&Setting{}).Error
}

func (s *SettingServiceImpl) List() ([]Setting, error) {
	var settings []Setting
	err := s.db.Model(&Setting{}).Find(&settings).Error
	return settings, err
}

func (s *SettingServiceImpl) ListType(settingType SettingType) ([]Setting, error) {
	var settings []Setting
	err := s.db.Model(&Setting{}).Where("type = ?", settingType).Find(&settings).Error
	return settings, err
}
