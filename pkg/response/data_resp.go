package response

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"net/http"
)

func JsonData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &model.ApiData{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: data,
	})
}

func XmlData(c *gin.Context, data interface{}) {
	c.XML(http.StatusOK, &model.ApiData{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: data,
	})
}

func Data(c *gin.Context, format string, args ...any) {
	c.String(http.StatusOK, format, args...)
}
