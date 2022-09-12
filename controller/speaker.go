package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"github.com/lezi-wiki/lezi-api/services/text"
	"math/rand"
	"time"
)

func SpeakerJsonHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	arr, err := text.GetTextBySpeaker(speaker)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.JSON(200, serializer.NewSuccessResponse(data))
}

func SpeakerXmlHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	arr, err := text.GetTextBySpeaker(speaker)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.XML(200, serializer.NewSuccessResponse(data))
}

func SpeakerTextHandler(c *gin.Context) {
	speaker := c.Param("speaker")

	arr, err := text.GetTextBySpeaker(speaker)
	if err != nil {
		c.JSON(404, serializer.NotFoundResponse())
		return
	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(arr))
	data := arr[index]

	c.String(200, data.Text)
}
