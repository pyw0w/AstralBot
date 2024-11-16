package web

import (
	log "AstralBot/internal/logger"
	"AstralBot/utils/config"
	"strings"
	"text/template"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func NewServer(cfg *config.Config, log *log.Logger) *Server {
	return &Server{
		config: cfg,
		logger: log,
	}
}

func (s *Server) Start() {
	// Create a new Iris application
	app := iris.New()

	// Add logger middleware
	app.Use(logger.New(
		logger.Config{
			Status: true,
			IP:     true,
			Method: true,
		},
	))

	// I18n support
	app.I18n.Loader.Funcs = func(current iris.Locale) template.FuncMap {
		return template.FuncMap{
			"uppercase": func(word string) string {
				return strings.ToUpper(word)
			},
		}
	}

	err := app.I18n.Load("./internal/web/locales/*/*.ini", "en-US", "el-GR")
	if err != nil {
		panic(err)
	}

	// Register the view engine
	app.RegisterView(iris.HTML("./internal/web/templates", ".html"))

	// Register the routes
	initializeRoutes(app)

	// Start the server on port 8080
	app.Listen(":8080")
}
