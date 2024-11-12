package web

import (
	v1 "AstralBot/web/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	// определение роута главной страницы
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	// определение роутера для api запросов
	apiGroup := router.Group("/api")
	{
		// регистрация роутера для версии 1
		v1.Register(apiGroup)
	}
}
