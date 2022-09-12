package routers

import "github.com/gin-gonic/gin"

func Handles(r *gin.RouterGroup, methods []string, path string, handler ...gin.HandlerFunc) {
	for _, method := range methods {
		r.Handle(method, path, handler...)
	}
}
