package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/controller"
	"github.com/lezi-wiki/lezi-api/middleware"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"net/http"
)

func InitRouter() *gin.Engine {
	// Gin router
	r := gin.Default()

	// cors
	r.Use(middleware.Cors())

	// Ping interface
	r.GET("/ping", func(c *gin.Context) {
		response.JsonData(c, &model.ApiData{
			Code: http.StatusOK,
			Msg:  http.StatusText(http.StatusOK),
			Data: "pong",
		})
	})

	// api v1
	v1 := r.Group("/api/v1")

	{
		v1.GET("global", controller.GlobalHandler)
		v1.POST("global", controller.GlobalHandler)

		v1.GET(":namespace/text", controller.TextHandler)
		v1.POST(":namespace/text", controller.TextHandler)

		v1.GET(":namespace/json", controller.JsonHandler)
		v1.POST(":namespace/json", controller.JsonHandler)
	}

	return r
}
