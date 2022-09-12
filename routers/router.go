package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/controller"
	"github.com/lezi-wiki/lezi-api/middleware"
	"github.com/lezi-wiki/lezi-api/pkg/serializer"
	"net/http"
)

func InitRouter() *gin.Engine {
	// Gin router
	r := gin.Default()

	// cors
	r.Use(middleware.Cors())

	// Ping interface
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, serializer.NewResponse(http.StatusOK, "pong", nil))
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://github.com/lezi-wiki/lezi-api/wiki")
	})

	// api v1
	v1 := r.Group("/api/v1")

	methodsSupport := []string{"GET", "POST", "HEAD", "OPTIONS"}

	{
		Handles(v1, methodsSupport, "global", controller.GlobalHandler)

		namespace := v1.Group(":namespace")
		{
			Handles(namespace, methodsSupport, "text", controller.NamespaceTextHandler)
			Handles(namespace, methodsSupport, "json", controller.NamespaceJsonHandler)
			Handles(namespace, methodsSupport, "xml", controller.NamespaceXmlHandler)
		}

		speaker := v1.Group("speaker/:speaker")
		{
			Handles(speaker, methodsSupport, "text", controller.SpeakerTextHandler)
			Handles(speaker, methodsSupport, "json", controller.SpeakerJsonHandler)
			Handles(speaker, methodsSupport, "xml", controller.SpeakerXmlHandler)
		}
	}

	return r
}
