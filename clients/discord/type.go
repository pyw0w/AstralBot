package discord

import (
	"AstralBot/internal/cmd"
	"AstralBot/internal/logger"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// loggingTransport Структура для перехвата HTTP запросов
// underlying - базовый RoundTripper для выполнения HTTP запросов
// logger - логгер для записи информации о запросах
type loggingTransport struct {
	underlying http.RoundTripper
	logger     *logger.Logger
}

// Handler Структура для обработки команд
// Session - сессия Discord
// CommandHandler - обработчик команд
// Debug - флаг для включения режима отладки
// Logger - логгер для записи информации
type Handler struct {
	Session        *discordgo.Session
	CommandHandler *cmd.CommandHandler
	Debug          bool
	Logger         *logger.Logger
}
