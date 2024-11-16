package discord

import (
	"fmt"
	"net/http"
	"strings"

	"AstralBot/clients/discord/events"
	"AstralBot/internal/cmd"
	"AstralBot/internal/logger"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	Session        *discordgo.Session
	CommandHandler *cmd.CommandHandler
	Debug          bool
	Logger         *logger.Logger
	handlers       map[string]func(*discordgo.Session, interface{})
}

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
		handlers:       make(map[string]func(*discordgo.Session, interface{})),
	}

	return handler, nil
}

// Структура для перехвата HTTP запросов
type loggingTransport struct {
	underlying http.RoundTripper
	logger     *logger.Logger
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
	h.Session.AddHandler(h.handleEvent)

	err := h.Session.Open()
	if err != nil {
		return err
	}

	if h.Debug {
		h.Logger.Debug("Discord", "Бот инициализирован как: "+h.Session.State.User.Username)
	}

	return nil
}

func (h *Handler) handleEvent(s *discordgo.Session, event interface{}) {
	switch e := event.(type) {
	case *discordgo.MessageCreate:
		h.message(s, e)
	case *discordgo.MessageReactionAdd:
		h.reactionAdd(s, e)
	}
}

func (h *Handler) message(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	// Логируем сообщение
	events.LogMessage(s, m, h.Logger)
	events.OnReady(s, &s.State.Ready)
	h.CommandHandler.ExecuteDiscordCommand(s, m)
}

func (h *Handler) reactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if handler, exists := h.handlers["reactionAdd"]; exists {
		handler(s, r)
	}
}

func (h *Handler) AddHandler(eventType string, handler func(*discordgo.Session, interface{})) {
	h.handlers[eventType] = handler
}

func (h *Handler) Close() error {
	return h.Session.Close()
}
