package conf

import (
	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/sirupsen/logrus"
)

var cfg *ini.File

// Init 初始化配置文件
func Init(path string) {
	var err error

	if path == "" || !util.Exists(path) {
		// 创建初始配置文件
		confContent := defaultConf
		f, err := util.CreatNestedFile(path)
		if err != nil {
			log.Log().Panicf("无法创建配置文件, %s", err)
		}

		// 写入配置文件
		_, err = f.WriteString(confContent)
		if err != nil {
			log.Log().Panicf("无法写入配置文件, %s", err)
		}

		f.Close()
		log.Log().Infof("配置文件初始化完成，文件位于 %s", path)
	}

	log.Log().Infof("将从 %s 解析配置文件", path)
	cfg, err = ini.Load(path)
	if err != nil {
		log.Log().Panicf("无法解析配置文件 '%s': %s", path, err)
	}

	sections := map[string]interface{}{
		"System":     SystemConfig,
		"DataSource": DataSourceConfig,
	}
	for sectionName, sectionStruct := range sections {
		err = mapSection(sectionName, sectionStruct)
		if err != nil {
			log.Log().Panicf("配置文件 %s 分区解析失败: %s", sectionName, err)
		}
	}

	// 重设log等级
	if SystemConfig.Debug {
		log.GlobalLogger = nil
		log.Log().SetLevel(logrus.DebugLevel)
	}
}

// mapSection 将配置文件的 Section 映射到结构体上
func mapSection(section string, confStruct interface{}) error {
	err := cfg.Section(section).MapTo(confStruct)
	if err != nil {
		return err
	}

	// 验证合法性
	validate := validator.New()
	err = validate.Struct(confStruct)
	if err != nil {
		return err
	}

	return nil
}
