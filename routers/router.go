package routers

import (
	"LeziAPI/Services"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Gin router
	r := gin.Default()

	// Ping interface
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// api v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("txt", services.ApiTxt)
		v1.GET("json", services.ApiJSON)
	}

	return r
}
