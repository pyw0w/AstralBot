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
	web := handlers.InitializeWebServer(cfg, log)
	handlers.StartHandlers(tgHandler, discordHandler, web, log)
	utils.WaitForShutdown(discordHandler, log)
}
