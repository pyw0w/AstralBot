package web

import (
	"AstralBot/config"
	_ "AstralBot/docs" // Этот импорт нужен для генерации документации
	"AstralBot/internal/logger"
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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

var router *gin.Engine

func (s *Server) Start() {
	// Set the Gin mode to release
	if !s.config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	// Создаем новый роутер
	router = gin.New()
	router.Use(gzip.Gzip(gzip.BestSpeed))

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("web/templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	initializeRoutes()

	// Запускаем сервер
	addr := fmt.Sprintf(":%d", s.config.WebPort)
	s.logger.Info("AstralWeb", fmt.Sprintf("Starting web server on %s", addr))
	if err := router.Run(addr); err != nil {
		s.logger.Error("Could not start web server: %v", err)
	}
}
