package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/pkg/text"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	textService "github.com/lezi-wiki/lezi-api/services/text"
)

func GlobalHandler(c *gin.Context) {
	var err error

	ns := c.Query("ns")
	speaker := c.Query("speaker")
	format := c.Query("format")

	var arr = text.Data

	if ns != "" {
		arr, err = textService.GetTextByNamespace(ns)
		if err != nil {
			c.JSON(404, serializer.NotFoundResponse())
			return
		}
	}

	var newArr = arr

	if speaker != "" {
		newArr = []model.TextData{}

		for _, d := range arr {
			if d.Speaker == speaker {
				newArr = append(newArr, d)
			}
		}
	}

	data := util.RandomItemFromSlice(newArr)
	switch format {
	case "json":
		c.JSON(200, serializer.NewSuccessResponse(data))
	case "xml":
		c.XML(200, serializer.NewSuccessResponse(data))
	case "text":
		c.String(200, data.Text)
	default:
		c.String(200, data.Text)
	}
}
