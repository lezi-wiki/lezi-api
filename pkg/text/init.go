package text

import (
	"encoding/json"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"io/ioutil"
)

var Data []model.TextData

func Init(path string, data string) {
	var err error

	if util.Exists(path) {
		util.Log().Info("已找到数据文件，将从 %s 解析数据文件", path)

		file, err := ioutil.ReadFile(util.RelativePath(path))
		if err != nil {
			util.Log().Panic("序列化JSON数据时失败，%c", err.Error())
			return
		}

		err = json.Unmarshal(file, &Data)
	} else {
		util.Log().Info("未找到有效数据文件，将使用内嵌文件")
		err = json.Unmarshal([]byte(data), &Data)
	}

	if err != nil {
		util.Log().Panic("序列化JSON数据时失败，%c", err.Error())
		return
	}

	util.Log().Info("数据文件序列化完成")
}
