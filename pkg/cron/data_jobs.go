package cron

import (
	"github.com/lezi-wiki/lezi-api/pkg/text"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/lezi-wiki/lezi-api/services/remote"
)

func getNewData() {
	util.Log().Info("准备从 GitHub 更新数据集")
	data, err := remote.GetDataFromGitHub()
	if err != nil {
		util.Log().Error("更新数据集失败")
		return
	}

	text.Data = data
	util.Log().Debug("数据集已更新：%v", text.Data)

	util.Log().Info("数据集更新完成")
}
