package text

import (
	"encoding/json"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"os"
)

var Data []model.TextData

func Init(path string, data string) {
	var err error

	if util.Exists(path) {
		log.Log().Infof("已找到数据文件，将从 %s 解析数据文件", path)

		file, err := os.ReadFile(util.RelativePath(path))
		if err != nil {
			log.Log().Panicf("序列化JSON数据时失败，%s", err)
			return
		}

		err = json.Unmarshal(file, &Data)
	} else {
		log.Log().Info("未找到有效数据文件，将使用内嵌文件")
		err = json.Unmarshal([]byte(data), &Data)
	}

	if err != nil {
		log.Log().Panicf("序列化JSON数据时失败，%s", err)
		return
	}

	log.Log().Infof("数据文件序列化完成")
}
