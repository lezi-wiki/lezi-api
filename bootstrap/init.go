package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/text"
	"github.com/lezi-wiki/lezi-api/services/remote"
)

func Init(dataPath string, data string, confPath string, updateEndpoint string) {
	printName()

	log.Log()

	// 初始化配置文件
	conf.Init(confPath)

	// 初始化数据
	text.Init(dataPath, data)

	// 设置更新
	remote.Endpoint = updateEndpoint

	// Debug 关闭时，切换为生产模式
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
