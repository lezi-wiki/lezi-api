package main

import (
	_ "embed"
	"flag"
	"github.com/lezi-wiki/lezi-api/bootstrap"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/lezi-wiki/lezi-api/routers"
)

//go:embed data.json
var data string

var isEject bool
var confPath string

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "配置文件路径")
	flag.BoolVar(&isEject, "eject", false, "导出内置静态资源")
	flag.Parse()

	bootstrap.Init(data, confPath)
}

func main() {
	if isEject {
		bootstrap.Eject(data)
		return
	}

	r := routers.InitRouter()

	err := r.Run(conf.SystemConfig.Listen)
	if err != nil {
		return
	}
}
