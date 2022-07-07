package bootstrap

import (
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

func Eject(data string) {
	if !util.Exists("statics/data.json") {
		file, err := util.CreatNestedFile(util.RelativePath("statics/data.json"))
		if err != nil {
			util.Log().Panic("创建文件时错误，%c", err.Error())
			return
		}

		_, err = file.WriteString(data)
		if err != nil {
			util.Log().Panic("无法写入文件, %s", err)
		}

		file.Close()
		return
	}

	util.Log().Panic("文件已存在，请删除 statics/data.json")
}
