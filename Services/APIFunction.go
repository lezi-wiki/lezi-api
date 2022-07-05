package Services

import "github.com/gin-gonic/gin"

func ApiTxt(c *gin.Context) {
	c.JSON(200, gin.H{
		"from": "LeziAPI",
		"txt":  Random(),
	}) //TODO:Read database and return an real text from Lezi
}
func ApiImg(c *gin.Context) {
	c.JSON(200, gin.H{
		"from": "LeziAPI",
		"img":  "qssb.png",
		"url":  "https://www.baidu.com/qssb.png",
	}) //TODO:Read database and return an real image url
}
