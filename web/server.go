package web

import (
	"AstralBot/config"
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
	http.HandleFunc("/", s.handleRoot)
	addr := fmt.Sprintf(":%d", s.config.WebPort)
	s.logger.Info("AstralWeb", fmt.Sprintf("Starting web server on %s", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		s.logger.Error("Could not start web server: %v", err)
	}
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to AstralBot Web Interface")
}
