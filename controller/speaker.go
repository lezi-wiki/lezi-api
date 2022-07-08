package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"github.com/lezi-wiki/lezi-api/services"
	"math/rand"
	"time"
)

func SpeakerJsonHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	arr, err := services.GetTextBySpeaker(speaker)
	if err != nil {
		response.NotFoundError(c)
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	response.JsonData(c, data)
}

func SpeakerTextHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	arr, err := services.GetTextBySpeaker(speaker)
	if err != nil {
		response.NotFoundError(c)
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	response.Data(c, data.Text)
}
