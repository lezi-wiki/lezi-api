package cron

import (
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"github.com/robfig/cron"
)

var task *cron.Cron

func init() {
	task = cron.New()
	task.Start()
}

func InitJobs() {
	// 每30分钟获取远端数据
	err := task.AddFunc("0 */30 * * *", getNewData)
	if err != nil {
		util.Log().Error("Cron jobs 错误， %s", err.Error())
		util.Log().Warning("自动更新服务未启动，数据将无法从 GitHub 自动获取！")
		return
	} else {
		util.Log().Info("自动更新服务已启动，数据将从 GitHub 自动获取！")
	}
}
