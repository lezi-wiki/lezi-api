package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

func NamespaceJsonHandler(c *gin.Context) {
	ns := c.Param("namespace")

	namespace, err := model.Client.Namespace.Get(model.Namespace{
		Name: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	text := util.RandomItemFromSlice(namespace.Texts)

	c.JSON(200, serializer.NewSuccessResponse(text))
}

func NamespaceTextHandler(c *gin.Context) {
	ns := c.Param("namespace")

	namespace, err := model.Client.Namespace.Get(model.Namespace{
		Name: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	text := util.RandomItemFromSlice(namespace.Texts)

	c.String(200, text.Text)
}

func NamespaceXmlHandler(c *gin.Context) {
	ns := c.Param("namespace")

	namespace, err := model.Client.Namespace.Get(model.Namespace{
		Name: ns,
	})
	if err != nil {
		log.Log().Errorf("获取数据失败: %s", err)
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	text := util.RandomItemFromSlice(namespace.Texts)

	c.XML(200, serializer.NewSuccessResponse(text))
}
