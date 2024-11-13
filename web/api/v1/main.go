package v1

import (
	"github.com/gin-gonic/gin"
)

func Register(route *gin.RouterGroup) {
	v1 := route.Group("/v1")
	{
		v1.GET("/ping", Ping)
		v1.GET("/generate_key", GenerateKey)
	}
}
