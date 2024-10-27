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

	// Регистрация команд
	cmdHandler.RegisterCommand(commands.Command{
		Name:        "ping",
		Description: "Проверка работоспособности бота",
		Execute: func(args []string) (string, error) {
			return "Pong!", nil
		},
	})

	// Инициализация обработчиков для разных платформ
	tgHandler, err := telegram.NewHandler(cfg.TelegramToken, cmdHandler, cfg.DebugMode, logger)
	if err != nil {
		logger.Error("System", "Ошибка инициализации Telegram:", err)
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
		logger.Error("System", "Ошибка инициализации Discord:", err)
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

	logger.Info("Завершение работы бота...")
}
