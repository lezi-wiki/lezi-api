package model

import (
	"errors"

	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"gorm.io/gorm"
)

func needMigrate(db *gorm.DB) bool {
	var s Setting
	err := db.Model(&Setting{}).Where(&Setting{
		Name: "version",
		Type: SettingTypeSystem,
	}).First(&s).Error
	if err != nil {
		return true
	}

	return s.Val != conf.Version
}

func migrate(db *gorm.DB) {
	if !needMigrate(db) {
		log.Log().Info("跳过数据库迁移阶段")
		return
	}

	log.Log().Info("开始数据库迁移")

	err := db.AutoMigrate(&Setting{}, &Text{})
	if err != nil {
		log.Log().Panicf("无法迁移数据库: %s", err)
	}

	addDefaultSettings(db)

	log.Log().Info("数据库迁移完成")
}

func addDefaultSettings(db *gorm.DB) {
	for _, value := range defaultSettings {
		err := db.Where(&Setting{
			Name: value.Name,
			Type: value.Type,
		}).First(&Setting{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Model(&Setting{}).Create(&value)
		}
	}
}
