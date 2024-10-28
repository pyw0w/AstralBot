package main

import (
	"os"
	"os/signal"
	"syscall"

	"AstralBot/config"
	"AstralBot/handlers/discord"
	"AstralBot/handlers/telegram"
	"AstralBot/internal/commands"
	"AstralBot/utils"
)

func main() {
	cfg := config.LoadConfig()
	logger := utils.NewLogger(cfg.DebugMode)

	// Инициализация обработчика команд
	cmdHandler := commands.NewCommandHandler()

	// Инициализация команд
	commands.RegisterCommands(cmdHandler, logger, cfg.DebugMode) // Передаем логгер и флаг дебага

	// Инициализация обработчиков для разных платформ
	tgHandler, err := telegram.NewHandler(cfg.TelegramToken, cmdHandler, cfg.DebugMode, logger)
	if err != nil {
		logger.Error("System", "Ошибка инициализации Telegram: ", err)
		return
	}

	discordHandler, err := discord.NewHandler(
		cfg.DiscordToken,
		cmdHandler,
		cfg.DebugMode,
		logger,
		cfg.DetailedAPILogs,
	)
	if err != nil {
		logger.Error("System", "Ошибка инициализации Discord: ", err)
		return
	}

	// Запуск обработчиков
	go tgHandler.Start()
	go discordHandler.Start()

	logger.Info("Бот запущен. Нажмите Ctrl+C для завершения")

	// Ожидание сигнала завершения
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Отключение Discord бота
	if err := discordHandler.Session.Close(); err != nil {
		logger.Error("System", "Ошибка при отключении Discord:", err)
	}

	logger.Info("Завершение работы бота...")
}
