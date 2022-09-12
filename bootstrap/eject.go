package bootstrap

import (
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

func Eject(data string) {
	if util.Exists("statics/data.json") {
		log.Log().Panicf("文件已存在，请删除 statics/data.json")
		return
	}

	file, err := util.CreatNestedFile(util.RelativePath("statics/data.json"))
	if err != nil {
		log.Log().Panicf("创建文件时错误，%c", err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Log().Panicf("无法写入文件, %s", err)
	}
}
