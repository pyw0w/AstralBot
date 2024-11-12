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
	// Создаем новый роутер
	r := gin.New()

	// Serve static files from the www directory
	r.Static("/static", "./www")

	// Load templates from the templates directory
	r.LoadHTMLGlob("www/templates/*")

	// Register route for the main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "base.html", gin.H{
			"Title":   "Astral Bot",
			"Content": "index.html",
		})
	})

	// Register API endpoints
	r.GET("/api/hello", func(c *gin.Context) {
		api.HelloHandler(c.Writer, c.Request)
	})

	// Register Swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	addr := fmt.Sprintf(":%d", s.config.WebPort)
	s.logger.Info("AstralWeb", fmt.Sprintf("Starting web server on %s", addr))
	if err := r.Run(addr); err != nil {
		s.logger.Error("Could not start web server: %v", err)
	}
}
