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

type TextService interface {
	Get(text Text) (*Text, error)
	GetTextByNamespace(namespace string) ([]Text, error)
	GetTextBySpeaker(speaker string) ([]Text, error)
	ListAll() ([]Text, error)
	CreateText(text Text) (*Text, error)
	DeleteText(id uint) error
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

func (t TextServiceImpl) DeleteText(id uint) error {
	return t.db.Model(&Text{}).Where("id = ?", id).Delete(&Text{}).Error
}
