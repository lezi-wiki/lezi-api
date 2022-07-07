package Services

import (
	"LeziAPI/JsonPkg"
	"github.com/gin-gonic/gin"
)

var InitJSON = JsonPkg.InitJSON()

func ApiJSON(c *gin.Context) {
	newRand := Random(InitJSON)
	GetTxt, GetAuthor := JsonPkg.GetTxt(newRand)
	c.JSON(200, gin.H{
		"from":   "LeziAPI",
		"txt":    GetTxt,
		"author": GetAuthor,
	})
}
func ApiTxt(c *gin.Context) {
	newRand := Random(InitJSON)
	GetTxt, _ := JsonPkg.GetTxt(newRand)
	c.String(200, GetTxt)
}
