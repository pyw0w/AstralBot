package web

import (
	"AstralBot/config"
	_ "AstralBot/docs" // Этот импорт нужен для генерации документации
	"AstralBot/internal/api"
	"AstralBot/internal/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config *config.Config
	logger *logger.Logger
}

// @title Astral Bot API
// @version 1.0
// @description This is the API documentation for Astral Bot.
// @host localhost:8080
// @BasePath /

func NewServer(cfg *config.Config, logger *logger.Logger) *Server {
	return &Server{
		config: cfg,
		logger: logger,
	}
}

func (s *Server) Start() {
	r := gin.Default()

	// Обслуживание статических файлов из каталога www
	r.Static("/static", "./www")

	// Загрузка шаблонов из каталога templates
	r.LoadHTMLGlob("www/templates/*")

	// Регистрация маршрута для главной страницы
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "base.html", gin.H{
			"Title":   "Astral Bot",
			"Content": "index.html",
		})
	})

	// Регистрация маршрута для страницы "About"
	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "base.html", gin.H{
			"Title":   "About Astral Bot",
			"Content": "about.html",
		})
	})

	// Регистрация API эндпоинтов
	r.GET("/api/hello", func(c *gin.Context) {
		api.HelloHandler(c.Writer, c.Request)
	})

	// Регистрация Swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	addr := fmt.Sprintf(":%d", s.config.WebPort)
	s.logger.Info("AstralWeb", fmt.Sprintf("Starting web server on %s", addr))
	if err := r.Run(addr); err != nil {
		s.logger.Error("Could not start web server: %v", err)
	}
}
