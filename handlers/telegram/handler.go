package telegram

import (
	"AstralBot/handlers/telegram/events"
	"AstralBot/internal/commands"
	"AstralBot/internal/logger"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot            *tgbotapi.BotAPI
	commandHandler *commands.CommandHandler
	debug          bool
	logger         *logger.Logger
}

func NewHandler(token string, cmdHandler *commands.CommandHandler, debug bool, logger *logger.Logger) (*Handler, error) {
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
	events.LogMessage(update)

	cmd := update.Message.Command()
	args := strings.Split(update.Message.Text, " ")[1:]

	if h.debug {
		h.logger.Debug("Telegram", "Команда: "+cmd+" Аргументы: "+strings.Join(args, " "))
	}

	response, _ := h.commandHandler.ExecuteCommand(cmd, args)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	h.bot.Send(msg)
}
