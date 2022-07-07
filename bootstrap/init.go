package bootstrap

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

var TextData []model.TextData

func Init(data string, confPath string) {
	util.Log()

	err := json.Unmarshal([]byte(data), &TextData)
	if err != nil {
		util.Log().Panic("序列化JSON数据时失败，%c", err.Error())
		return
	}

	conf.Init(confPath)

	// Debug 关闭时，切换为生产模式
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
