package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"gorm.io/gorm"
)

func GlobalHandler(c *gin.Context) {
	var err error
	var payload model.Text

	speaker := c.Query("speaker")
	if speaker != "" {
		payload.Speaker = speaker
	}

	ns := c.Query("ns")
	if ns != "" {
		payload.Namespace = ns
	}

	text, err := model.Client.Text.RandomRecord(payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, serializer.NotFoundResponse())
			return
		}

		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(500, serializer.NewErrorResponse(500, "database error"))
		return
	}

	format := c.Query("format")

	switch format {
	case "json":
		c.JSON(200, serializer.NewSuccessResponse(text))
	case "xml":
		c.XML(200, serializer.NewSuccessResponse(text))
	case "text":
		c.String(200, text.Text)
	default:
		c.String(200, text.Text)
	}
}
