package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/services/text"
	"math/rand"
	"time"
)

func NamespaceJsonHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := text.GetTextByNamespace(ns)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.JSON(200, serializer.NewSuccessResponse(data))
}

func NamespaceTextHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := text.GetTextByNamespace(ns)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.String(200, data.Text)
}

func NamespaceXmlHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := text.GetTextByNamespace(ns)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.XML(200, serializer.NewSuccessResponse(data))
}
