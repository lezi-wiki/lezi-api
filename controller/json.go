package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"github.com/lezi-wiki/lezi-api/services"
	"math/rand"
	"net/http"
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

	c.JSON(http.StatusOK, &model.ApiData{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusNotFound),
		Data: data,
	})
	c.Done()
}
