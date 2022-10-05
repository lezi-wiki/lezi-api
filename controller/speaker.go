package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/pkg/serializer/vo"
)

func SpeakerJsonHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Speaker: speaker,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.JSON(200, serializer.NewSuccessResponse(vo.BuildTextVO(data)))
}

func SpeakerXmlHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Speaker: speaker,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.XML(200, serializer.NewSuccessResponse(vo.BuildTextVO(data)))
}

func SpeakerTextHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Speaker: speaker,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.String(200, data.Text)
}
