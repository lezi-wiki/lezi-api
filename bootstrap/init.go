package bootstrap

import (
	"encoding/json"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

var TextData []model.TextData

func Init(data string) {
	util.Log()

	err := json.Unmarshal([]byte(data), &TextData)
	if err != nil {
		util.Log().Panic("序列化JSON数据时失败，%c", err.Error())
		return
	}
}
