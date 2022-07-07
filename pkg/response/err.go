package response

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/model"
	"net/http"
)

func Error(c *gin.Context, err int) {
	c.JSON(http.StatusOK, &model.ApiData{
		Code: err,
		Msg:  http.StatusText(err),
	})
}

func NotFoundError(c *gin.Context) {
	Error(c, http.StatusNotFound)
}

func ServerError(c *gin.Context) {
	Error(c, http.StatusInternalServerError)
}
