package main

import (
	_ "embed"
	"flag"
	"github.com/lezi-wiki/lezi-api/bootstrap"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/cron"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/lezi-wiki/lezi-api/routers"
)

var (
	isEject        bool
	confPath       string
	updateEndpoint string
)

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "配置文件路径")
	flag.StringVar(&updateEndpoint, "update", "https://raw.githubusercontent.com/lezi-wiki/lezi-api/master/data.json", "数据更新地址")
	flag.BoolVar(&isEject, "eject", false, "导出内置静态资源")
	flag.Parse()

	bootstrap.Init(confPath, updateEndpoint)
}

func main() {
	if updateEndpoint != "false" {
		cron.InitJobs()
	} else {
		log.Log().Warning("自动更新服务未启动，数据将无法从 GitHub 自动获取！")
	}

	r := routers.InitRouter()

	log.Log().Infof("应用将监听 %s", conf.SystemConfig.Listen)
	err := r.Run(conf.SystemConfig.Listen)
	if err != nil {
		log.Log().Panicf("尝试监听 %s 时发生错误，%s", conf.SystemConfig.Listen, err.Error())
		return
	}
}
