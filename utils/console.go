package utils

import (
	"AstralBot/clients/discord"
	"AstralBot/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown(discordHandler *discord.Handler, log *logger.Logger) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	if discordHandler != nil {
		err := discordHandler.Close()
		if err != nil {
			log.Error("Ошибка закрытия Discord клиента: %v\n", err)
			return
		}
	}
	log.Info("AstralBot", "Завершение работы бота...")
}
