package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Println("Web server is running")
	r.Run()
}
