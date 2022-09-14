package cron

import (
	"github.com/robfig/cron/v3"
)

var task *cron.Cron

func init() {
	task = cron.New()
	task.Start()
}

func InitJobs() {
}
