package cron

import (
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/text"
	"github.com/lezi-wiki/lezi-api/services/remote"
)

func getNewData() {
	log.Log().Infof("准备从 GitHub 更新数据集")
	data, err := remote.GetDataFromGitHub()
	if err != nil {
		log.Log().Errorf("更新数据集失败")
		return
	}

	text.Data = data
	log.Log().Debugf("数据集已更新：%v", text.Data)

	log.Log().Infof("数据集更新完成")
}
