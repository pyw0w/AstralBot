package web

import (
	"AstralBot/config"
	"AstralBot/internal/api"
	"AstralBot/internal/logger"
	"fmt"
	"net/http"
)

type Server struct {
	config *config.Config
	logger *logger.Logger
}

func NewServer(cfg *config.Config, logger *logger.Logger) *Server {
	return &Server{
		config: cfg,
		logger: logger,
	}
}

func (s *Server) Start() {
	// Обслуживание статических файлов из каталога www
	fs := http.FileServer(http.Dir("./www"))
	http.Handle("/", fs)

	// Регистрация API эндпоинтов
	http.HandleFunc("/api/hello", api.HelloHandler)

	addr := fmt.Sprintf(":%d", s.config.WebPort)
	s.logger.Info("AstralWeb", fmt.Sprintf("Starting web server on %s", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		s.logger.Error("Could not start web server: %v", err)
	}
}
