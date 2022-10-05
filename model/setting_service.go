package model

import (
	"github.com/lezi-wiki/lezi-api/pkg/cache"
	"gorm.io/gorm"
)

type SettingService interface {
	Get(name string, settingType SettingType) (string, error)
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

func (s *SettingServiceImpl) Get(name string, settingType SettingType) (string, error) {
	var setting Setting

	if value, ok := cache.Get("setting_" + name); ok {
		return value.(string), nil
	}

	err := s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).First(&setting).Error
	return setting.Val, err
}

func (s *SettingServiceImpl) Set(name string, settingType SettingType, val string) error {
	err := s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).Updates(&Setting{
		Val: val,
	}).Error
	if err != nil {
		return err
	}

	cache.Set("setting_"+name, val, 0)
	return nil
}

func (s *SettingServiceImpl) Delete(name string, settingType SettingType) error {
	err := s.db.Model(&Setting{}).Where(&Setting{
		Name: name,
		Type: settingType,
	}).Delete(&Setting{}).Error
	if err != nil {
		return err
	}

	cache.Deletes([]string{name}, "setting_")
	return nil
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
