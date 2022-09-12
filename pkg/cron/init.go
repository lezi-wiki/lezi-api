package cron

import (
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/robfig/cron/v3"
)

var task *cron.Cron

func init() {
	task = cron.New()
	task.Start()
}

func InitJobs() {
	// 每30分钟获取远端数据
	_, err := task.AddFunc("0 */30 * * *", getNewData)
	if err != nil {
		log.Log().Errorf("Cron jobs 错误， %s", err.Error())
		log.Log().Warning("自动更新服务未启动，数据将无法从 GitHub 自动获取！")
		return
	} else {
		log.Log().Infof("自动更新服务已启动，数据将从 GitHub 自动获取！")
	}
}
