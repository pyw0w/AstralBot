package main

import (
	"AstralBot/handlers"
	"AstralBot/internal/logger"
	"AstralBot/utils"
	"AstralBot/utils/config"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)

	//handlers.InitDB(cfg, log)
	utils.SetConsoleTitle(log, cfg.DebugMode)
	web := handlers.InitializeWebServer(cfg, log)
	cmdHandler := handlers.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := handlers.InitializeHandlers(cfg, cmdHandler, log)
	handlers.StartHandlers(tgHandler, discordHandler, web, log)
	utils.WaitForShutdown(discordHandler, log)
}
