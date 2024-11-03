package main

import (
	"AstralBot/config"
	"AstralBot/handlers/discord"
	"AstralBot/handlers/telegram"
	"AstralBot/internal/commands"
	"AstralBot/utils"
	"os"
	"os/signal"
	"syscall"
)

// SetConsoleTitle устанавливает заголовок окна консоли
//func SetConsoleTitle(title string) {
//	kernel32 := syscall.NewLazyDLL("kernel32.dll")
//	setConsoleTitle := kernel32.NewProc("SetConsoleTitleW")
//	_, _, _ = setConsoleTitle.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
//}

func main() {
	cfg := config.LoadConfig()
	logger := utils.NewLogger(cfg.DebugMode)

	// Устанавливаем заголовок консоли
	//SetConsoleTitle("AstralBot - Версия: " + internal.Version)

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
