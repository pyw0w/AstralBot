package utils

import (
	"AstralBot/handlers/discord"
	"AstralBot/internal"
	"AstralBot/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown(discordHandler *discord.Handler, log *logger.Logger) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	if err := discordHandler.Session.Close(); err != nil {
		log.Error("System", "Ошибка при отключении Discord:", err)
	}

	log.Info("AstralBot", "Завершение работы бота...")
}

func SetConsoleTitle(log *logger.Logger, debugMode bool) {
	err := SetTitle("AstralBot - Версия: " + internal.Version)
	if err != nil {
		log.Error("Ошибка установки заголовка: %v\n", err)
	} else if debugMode {
		log.Debug("AstralBot", "Заголовок успешно установлен")
	}
}
