package main

import (
	"AstralBot/clients"
	"AstralBot/clients/db"
	"AstralBot/internal"
	"AstralBot/internal/logger"
	"AstralBot/utils"
	"AstralBot/utils/config"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)

	// Connect to the database
	db.Connect()

	// Initialize the database
	// handlers.InitDB(cfg, log)

	internal.CheckForNewVersion(log)
	utils.SetConsoleTitle(log, cfg.DebugMode)
	web := clients.InitializeWebServer(cfg, log)
	cmdHandler := clients.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := clients.InitializeHandlers(cfg, cmdHandler, log)
	clients.StartHandlers(tgHandler, discordHandler, web, log)
	utils.WaitForShutdown(discordHandler, log)
}
