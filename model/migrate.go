package model

import (
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
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
	if !needMigrate(db) && !conf.SystemConfig.Debug {
		log.Log().Info("跳过数据库迁移阶段")
		return
	}

	log.Log().Info("开始数据库迁移")

	err := db.AutoMigrate(&Setting{}, &Text{}, &User{}, &Namespace{})
	if err != nil {
		log.Log().Panicf("无法迁移数据库: %s", err)
	}

	addDefaultSettings()
	addDefaultUser()

	log.Log().Info("数据库迁移完成")
}

func addDefaultSettings() {
	for _, value := range defaultSettings {
		if !Client.Setting.Exists(value.Name, &value.Type) {
			err := Client.Setting.Set(value.Name, value.Type, value.Val)
			if err != nil {
				log.Log().Errorf("无法设置默认设置 %s: %s", value.Name, err)
				continue
			}
		}
	}
}

func addDefaultUser() {
	if !Client.User.Exists(User{}) {
		pass := util.RandStringRunes(16)
		_, err := Client.User.CreateUser(User{
			Username: "admin",
			Password: pass,
			Email:    "admin@code.lezi.wiki",
			Avatar:   nil,
			Texts:    nil,
		})
		if err != nil {
			log.Log().Errorf("无法创建默认用户: %s", err)
			return
		}

		log.Log().Infof("默认用户已创建。 用户名: admin  邮箱: admin@code.lezi.wiki  密码: %s", pass)
	}
}
