package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
)

func NamespaceJsonHandler(c *gin.Context) {
	ns := c.Param("namespace")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Namespace: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.JSON(200, serializer.NewSuccessResponse(data))
}

func NamespaceTextHandler(c *gin.Context) {
	ns := c.Param("namespace")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Namespace: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.String(200, data.Text)
}

func NamespaceXmlHandler(c *gin.Context) {
	ns := c.Param("namespace")

	data, err := model.Client.Text.RandomRecord(model.Text{
		Namespace: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	c.XML(200, serializer.NewSuccessResponse(data))
}
