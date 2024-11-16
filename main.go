package main

import (
	"AstralBot/clients"
	"AstralBot/internal"
	"AstralBot/internal/logger"
	"AstralBot/utils"
	"AstralBot/utils/config"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)
	// Initialize the database
	// handlers.InitDB(cfg, log)

	internal.CheckForNewVersion(log)
	web := clients.InitializeWebServer(cfg, log)
	cmdHandler := clients.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := clients.InitializeHandlers(cfg, cmdHandler, log)
	clients.StartHandlers(tgHandler, discordHandler, web, log)
	utils.WaitForShutdown(discordHandler, log)
}
