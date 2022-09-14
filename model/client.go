package model

import (
	"gorm.io/gorm"
)

var Client *DataClient

type DataClient struct {
	db *gorm.DB

	Setting   SettingService
	Text      TextService
	User      UserService
	Namespace NamespaceService
}

func NewDataClient(db *gorm.DB) *DataClient {
	c := &DataClient{
		db: db,
	}

	c.Setting = NewSettingService(c.db)
	c.Text = NewTextService(c.db)
	c.User = NewUserService(c.db)
	c.Namespace = NewNamespaceService(c.db)

	return c
}
