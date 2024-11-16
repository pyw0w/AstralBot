package web

import (
	"AstralBot/internal/logger"
	"AstralBot/utils/config"
)

type Server struct {
	config *config.Config
	logger *logger.Logger
}
