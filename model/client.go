package model

import (
	"gorm.io/gorm"
)

var Client *DataClient

type DataClient struct {
	db *gorm.DB

	Setting SettingService
	Text    TextService
}

func NewDataClient(db *gorm.DB, settingService SettingService, textService TextService) *DataClient {
	c := &DataClient{
		db: db,
	}

	c.Setting = settingService
	c.Text = textService

	return c
}
