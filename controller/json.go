package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"github.com/lezi-wiki/lezi-api/services"
	"math/rand"
	"time"
)

func JsonHandler(c *gin.Context) {
	ns := c.Param("namespace")

	arr, err := services.GetTextByNamespace(ns)
	if err != nil {
		response.NotFoundError(c)
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	response.JsonData(c, data)
}
