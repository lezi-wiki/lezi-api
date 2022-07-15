package cron

import (
	"github.com/lezi-wiki/lezi-api/pkg/text"
	"github.com/lezi-wiki/lezi-api/services/remote"
)

func getNewData() {
	data, err := remote.GetDataFromGitHub()
	if err != nil {
		return
	}

	text.Data = data
}
