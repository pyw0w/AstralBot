package discord

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"AstralBot/clients/discord/events"
	"AstralBot/internal/cmd"
	"AstralBot/internal/logger"

	"github.com/bwmarrin/discordgo"
)

var (
	startTime    = time.Now()
	commandCount int
)

func NewHandler(token string, cmdHandler *cmd.CommandHandler, debug bool, logger *logger.Logger, detailedLogs bool) (*Handler, error) {
	// Перехватываем логи библиотеки до создания сессии
	discordgo.Logger = func(msgL, caller int, format string, a ...interface{}) {
		if !detailedLogs {
			return // Пропускаем детальные логи если они отключены
		}
		msg := fmt.Sprintf(format, a...)
		logger.Debug("Discord-Library", msg)
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	// Настраиваем кастомный HTTP клиент только если включены детальные логи
	if detailedLogs {
		session.Client = &http.Client{
			Transport: &loggingTransport{
				underlying: http.DefaultTransport,
				logger:     logger,
			},
		}
	}

	// Настраиваем логгер Discord
	session.LogLevel = discordgo.LogError // Изменяем уровень логирования на Error
	session.SyncEvents = true
	session.StateEnabled = true
	session.Debug = detailedLogs // Используем detailedLogs вместо debug
	session.ShouldReconnectOnError = true
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	handler := &Handler{
		Session:        session,
		CommandHandler: cmdHandler,
		Debug:          debug,
		Logger:         logger,
	}

	return handler, nil
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Не логируем запросы если они не относятся к Discord API
	if !strings.Contains(req.URL.Host, "discord.com") {
		return t.underlying.RoundTrip(req)
	}

	// Если включен дебаг режим, логируем запрос через наш логгер
	if t.logger != nil {
		t.logger.Debug("Discord-HTTP", fmt.Sprintf("%s %s", req.Method, req.URL.Path))
	}

	return t.underlying.RoundTrip(req)
}

func (h *Handler) Start() error {
	h.Session.AddHandler(h.HandleEvent)
	err := h.Session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) HandleEvent(s *discordgo.Session, event interface{}) {
	switch e := event.(type) {
	case *discordgo.Ready:
		if h.Debug {
			h.Logger.Debug("Discord", "Бот инициализирован как: "+e.User.Username)
		}
		// Update the status with the number of unique users
		Status(s)
	case *discordgo.MessageCreate:
		if e.Author.Bot {
			return
		}

		if !strings.HasPrefix(e.Content, "!") {
			return
		}

		// Логируем сообщение
		events.LogMessage(s, e, h.Logger)
		h.CommandHandler.ExecuteDiscordCommand(s, e)
	}
}

// Add external event handlers
func (h *Handler) AddEventHandler(handler interface{}) {
	h.Session.AddHandler(handler)
}

func (h *Handler) Close() error {
	return h.Session.Close()
}
