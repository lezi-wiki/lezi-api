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

func NewDataClient(db *gorm.DB) *DataClient {
	c := &DataClient{
		db: db,
	}

	c.Setting = NewSettingService(c.db)
	c.Text = NewTextService(c.db)

	return c
}
