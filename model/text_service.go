package model

import (
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"gorm.io/gorm"
)

type TextService interface {
	Get(text Text) (*Text, error)
	List(text Text) ([]Text, error)
	GetTextByNamespace(namespace string) ([]Text, error)
	GetTextBySpeaker(speaker string) ([]Text, error)
	ListAll() ([]Text, error)
	CreateText(text Text) (*Text, error)
	UpdateText(text Text) (*Text, error)
	DeleteText(id uint) error
	Count() int64
	RandomRecord(rule Text) (*Text, error)
	Exists(text Text) bool
}

type TextServiceImpl struct {
	db *gorm.DB
}

func NewTextService(db *gorm.DB) TextService {
	return &TextServiceImpl{db: db}
}

func (t TextServiceImpl) Get(text Text) (*Text, error) {
	err := t.db.Model(&Text{}).Where(&text).First(&text).Error
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func (t TextServiceImpl) List(text Text) ([]Text, error) {
	var data []Text
	err := t.db.Model(&Text{}).Where(&text).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t TextServiceImpl) GetTextByNamespace(namespace string) ([]Text, error) {
	var data []Text
	err := t.db.Model(&Text{}).Where("namespace = ?", namespace).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t TextServiceImpl) GetTextBySpeaker(speaker string) ([]Text, error) {
	var data []Text
	err := t.db.Model(&Text{}).Where("speaker = ?", speaker).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t TextServiceImpl) ListAll() ([]Text, error) {
	var data []Text
	err := t.db.Model(&Text{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t TextServiceImpl) CreateText(text Text) (*Text, error) {
	err := t.db.Model(&Text{}).Create(&text).Error
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func (t TextServiceImpl) UpdateText(text Text) (*Text, error) {
	err := t.db.Model(&Text{}).Where("id = ?", text.ID).Updates(&text).Error
	if err != nil {
		return nil, err
	}

	return &text, nil
}

func (t TextServiceImpl) DeleteText(id uint) error {
	return t.db.Model(&Text{}).Where("id = ?", id).Delete(&Text{}).Error
}

func (t TextServiceImpl) Count() int64 {
	var count int64
	err := t.db.Model(&Text{}).Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

func (t TextServiceImpl) RandomRecord(rule Text) (*Text, error) {
	var count int64
	if err := t.db.Model(&Text{}).Where(&rule).Count(&count).Error; err != nil {
		return nil, err
	}

	var text Text
	if err := t.db.Model(&Text{}).Where(&rule).Offset(util.RandomInt(0, int(count-1))).First(&text).Error; err != nil {
		return nil, err
	}

	return &text, nil
}

func (t TextServiceImpl) Exists(text Text) bool {
	var count int64
	err := t.db.Model(&Text{}).Where(&text).Count(&count).Error
	if err != nil {
		return false
	}

	return count > 0
}
