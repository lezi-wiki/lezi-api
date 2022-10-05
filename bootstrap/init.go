package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/cron/jobs"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/lezi-wiki/lezi-api/services/remote"
	"os"
)

func Init(confPath string, updateEndpoint string) {
	printName()

	log.Log()

	// 初始化配置文件
	conf.Init(confPath)

	if conf.SystemConfig.HashIDSalt == "" {
		log.Log().Warn("HashIDSalt 未设置，将使用随机值")
		conf.SystemConfig.HashIDSalt = util.RandStringRunes(32)
		_ = os.Setenv("HASHID_SALT", conf.SystemConfig.HashIDSalt)
	}

	// 初始化数据库
	model.Init()

	// 设置更新
	remote.Endpoint = updateEndpoint

	go jobs.UpdateData()

	// Debug 关闭时，切换为生产模式
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
