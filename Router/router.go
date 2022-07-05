package Router

import (
	"LeziAPI/Services"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	rt := gin.Default() //Init Gin

	//Ping test
	rt.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//API Router group
	v1 := rt.Group("/api")
	{
		v1.GET("/txt", Services.ApiTxt)
		v1.GET("/img", Services.ApiImg)
	}
	err := rt.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}