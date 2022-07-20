package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"github.com/lezi-wiki/lezi-api/services/text"
	"math/rand"
	"time"
)

func NamespaceJsonHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := text.GetTextByNamespace(ns)
	if err != nil {
		response.NotFoundError(c)
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	response.JsonData(c, data)
}

func NamespaceTextHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := text.GetTextByNamespace(ns)
	if err != nil {
		response.NotFoundError(c)
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	response.Data(c, data.Text)
}
