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

		namespace := v1.Group(":namespace")
		{
			namespace.GET("text", controller.NamespaceTextHandler)
			namespace.POST("text", controller.NamespaceTextHandler)

			namespace.GET("json", controller.NamespaceJsonHandler)
			namespace.POST("json", controller.NamespaceJsonHandler)

			namespace.GET("xml", controller.NamespaceXmlHandler)
			namespace.POST("xml", controller.NamespaceXmlHandler)
		}

		speaker := v1.Group("speaker/:speaker")
		{
			speaker.GET("text", controller.SpeakerTextHandler)
			speaker.POST("text", controller.SpeakerTextHandler)

			speaker.GET("json", controller.SpeakerJsonHandler)
			speaker.POST("json", controller.SpeakerJsonHandler)

			speaker.GET("xml", controller.SpeakerXmlHandler)
			speaker.POST("xml", controller.SpeakerXmlHandler)
		}
	}

	return r
}
