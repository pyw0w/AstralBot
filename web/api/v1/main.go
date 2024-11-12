package v1

import (
	"github.com/gin-gonic/gin"
)

func Register(route *gin.Engine) {
	v1 := route.Group("/v1")
	{
		v1.GET("/ping", Ping)
	}
}
