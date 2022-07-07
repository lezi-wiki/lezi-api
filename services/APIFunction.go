package services

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/json"
)

var InitJSON = json.InitJSON()

func ApiJSON(c *gin.Context) {
	newRand := Random(InitJSON)
	GetTxt, GetAuthor := json.GetTxt(newRand)
	c.JSON(200, gin.H{
		"from":   "LeziAPI",
		"txt":    GetTxt,
		"author": GetAuthor,
	})
}
func ApiTxt(c *gin.Context) {
	newRand := Random(InitJSON)
	GetTxt, _ := json.GetTxt(newRand)
	c.String(200, GetTxt)
}
