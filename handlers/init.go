package handlers

import (
	"AstralBot/config"
	"AstralBot/handlers/discord"
	"AstralBot/handlers/telegram"
	"AstralBot/internal/commands"
	"AstralBot/internal/logger"
	"os"
)

func InitializeCommandHandler(log *logger.Logger, debugMode bool) *commands.CommandHandler {
	cmdHandler := commands.NewCommandHandler()
	commands.RegisterCommands(cmdHandler, log, debugMode)
	return cmdHandler
}

func InitializeHandlers(cfg *config.Config, cmdHandler *commands.CommandHandler, log *logger.Logger) (*telegram.Handler, *discord.Handler) {
	tgHandler, err := telegram.NewHandler(cfg.TelegramToken, cmdHandler, cfg.DebugMode, log)
	if err != nil {
		log.Error("System", "Ошибка инициализации Telegram: ", err)
		os.Exit(1)
	}

	discordHandler, err := discord.NewHandler(cfg.DiscordToken, cmdHandler, cfg.DebugMode, log, cfg.DetailedAPILogs)
	if err != nil {
		log.Error("System", "Ошибка инициализации Discord: ", err)
		os.Exit(1)
	}

	return tgHandler, discordHandler
}

func StartHandlers(tgHandler *telegram.Handler, discordHandler *discord.Handler, log *logger.Logger) {
	go tgHandler.Start()
	go discordHandler.Start()
	log.Info("AstralBot", "Бот запущен. Нажмите Ctrl+C для завершения")
}
