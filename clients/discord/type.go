package discord

import (
	"AstralBot/internal/cmd"
	"AstralBot/internal/logger"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Структура для перехвата HTTP запросов
type loggingTransport struct {
	underlying http.RoundTripper
	logger     *logger.Logger
}

type Handler struct {
	Session        *discordgo.Session
	CommandHandler *cmd.CommandHandler
	Debug          bool
	Logger         *logger.Logger
}
