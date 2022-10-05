package jobs

import (
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer/dto"
	"github.com/lezi-wiki/lezi-api/services/remote"
	"github.com/samber/lo"
)

func UpdateData() {
	log.Log().Infof("准备从 GitHub 更新数据集")
	data, err := remote.GetDataFromGitHub()
	if err != nil {
		log.Log().Errorf("更新数据集失败")
		return
	}

	models := lo.Map(data, func(item dto.TextJsonDTO, i int) model.Text {
		return dto.BuildTextJsonDTO(item)
	})

	for _, datum := range models {
		if exist := model.Client.Text.Exists(datum); exist {
			continue
		}

		if _, err := model.Client.Text.CreateText(datum); err != nil {
			log.Log().Errorf("对于命名空间 %s 同步发言人 %s 的数据 %s 失败", datum.Namespace, datum.Speaker, datum.Text)
			continue
		}
	}

	log.Log().Infof("数据集更新完成，远端获取数据 %d 条，当前数据 %d 条", len(data), model.Client.Text.Count())
}
