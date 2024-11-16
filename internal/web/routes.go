package web

import (
	"github.com/kataras/iris/v12"
)

func initializeRoutes(app *iris.Application) {
	// определение роута главной страницы
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("title", "Home Page")
		ctx.View("index.html")
	})

	app.Get("/about", func(ctx iris.Context) {
		ctx.ViewData("title", "About Page")
		ctx.View("about.html")
	})
}
