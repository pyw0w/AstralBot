package telegram

import (
	"AstralBot/handlers/telegram/events"
	"AstralBot/internal/cmd"
	"AstralBot/internal/logger"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot            *tgbotapi.BotAPI
	commandHandler *cmd.CommandHandler
	debug          bool
	logger         *logger.Logger
}

func NewHandler(token string, cmdHandler *cmd.CommandHandler, debug bool, logger *logger.Logger) (*Handler, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// Включаем дебаг режим и перенаправляем логи в наш логгер
	if debug {
		bot.Debug = true
		// Создаем новый логгер для Telegram API
		tgLogger := log.New(logger.NewAdapter("Telegram-Library"), "", 0)
		// Устанавливаем его как логгер для библиотеки
		tgbotapi.SetLogger(tgLogger)
	}

	handler := &Handler{
		bot:            bot,
		commandHandler: cmdHandler,
		debug:          debug,
		logger:         logger,
	}

	if debug {
		logger.Debug("Telegram", "Бот инициализирован как: "+bot.Self.UserName)
	}

	return handler, nil
}

func (h *Handler) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			go h.handleCommand(update)
		}
	}
}

func (h *Handler) handleCommand(update tgbotapi.Update) {
	if update.Message.From.IsBot {
		return
	}

	// Логируем сообщение
	events.LogMessage(update, h.logger)

	cmd := update.Message.Command()
	args := strings.Split(update.Message.Text, " ")[1:]

	response, _ := h.commandHandler.ExecuteCommand(cmd, args)
	switch resp := response.(type) {
	case string:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp)
		h.bot.Send(msg)
	case *discordgo.MessageEmbed:
		var messageText strings.Builder
		messageText.WriteString(resp.Description + "\n\n")
		for _, field := range resp.Fields {
			messageText.WriteString(field.Name + ": " + field.Value + "\n")
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText.String())
		h.bot.Send(msg)
	default:
		h.logger.Error("Telegram", "Failed to assert response to known type")
	}
}
