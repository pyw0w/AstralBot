package main

import (
	"AstralBot/config"
	"AstralBot/handlers"
	"AstralBot/internal/logger"
	"AstralBot/utils"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)

	utils.SetConsoleTitle(log, cfg.DebugMode)
	cmdHandler := handlers.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := handlers.InitializeHandlers(cfg, cmdHandler, log)

	handlers.StartHandlers(tgHandler, discordHandler, log)
	utils.WaitForShutdown(discordHandler, log)
}
