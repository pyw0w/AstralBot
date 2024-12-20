package clients

import (
	"AstralBot/clients/discord"
	"AstralBot/clients/telegram"
	"AstralBot/internal/cmd"
	"AstralBot/internal/commands"
	"AstralBot/internal/logger"
	"AstralBot/internal/web"
	"AstralBot/utils/config"

	"os"
)

// InitializeCommandHandler инициализирует обработчик команд
func InitializeCommandHandler(log *logger.Logger, debugMode bool) *cmd.CommandHandler {
	cmdHandler := cmd.NewCommandHandler()
	commands.RegisterCommands(cmdHandler)
	return cmdHandler
}

// InitializeHandlers инициализирует обработчики
func InitializeHandlers(cfg *config.Config, cmdHandler *cmd.CommandHandler, log *logger.Logger) (*telegram.Handler, *discord.Handler) {
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

// InitializeWebServer инициализирует веб-сервер
func InitializeWebServer(cfg *config.Config, log *logger.Logger) *web.Server {
	webServer := web.NewServer(cfg, log)
	return webServer
}

// StartHandlers запускает обработчики
func StartHandlers(tgHandler *telegram.Handler, discordHandler *discord.Handler, webServer *web.Server, log *logger.Logger) {
	if tgHandler == nil {
		log.Info("System", "Запуск Telegram отключен")
	} else {
		go tgHandler.Start()
	}
	if discordHandler == nil {
		log.Info("System", "Запуск Discord отключен")
	} else {
		go discordHandler.Start()
	}
	if webServer == nil {
		log.Info("System", "Запуск веб-сервера отключен")
	} else {
		go webServer.Start()
	}
}
