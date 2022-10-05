package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/pkg/serializer/vo"
	"gorm.io/gorm"
)

func GlobalHandler(c *gin.Context) {
	var err error

	ns := c.Query("ns")
	speaker := c.Query("speaker")
	format := c.Query("format")

	text, err := model.Client.Text.RandomRecord(model.Text{
		Namespace: ns,
		Speaker:   speaker,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, serializer.NotFoundResponse())
			return
		}

		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(500, serializer.NewErrorResponse(500, "database error"))
		return
	}

	switch format {
	case "json":
		c.JSON(200, serializer.NewSuccessResponse(vo.BuildTextVO(text)))
	case "xml":
		c.XML(200, serializer.NewSuccessResponse(vo.BuildTextVO(text)))
	case "text":
		c.String(200, text.Text)
	default:
		c.String(200, text.Text)
	}
}
