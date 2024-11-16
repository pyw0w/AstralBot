package main

import (
	"AstralBot/clients"
	"AstralBot/internal"
	"AstralBot/internal/logger"
	"AstralBot/utils"
	"AstralBot/utils/config"
	"flag"
)

func main() {
	// Define a command-line flag
	runWebOnly := flag.Bool("web-only", false, "Run only the web server")
	flag.Parse()

	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)
	// Initialize the database
	// handlers.InitDB(cfg, log)

	internal.CheckForNewVersion(log)
	web := clients.InitializeWebServer(cfg, log)

	if *runWebOnly {
		web.Start()
		utils.WaitForShutdown(nil, log)
		return
	}

	cmdHandler := clients.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := clients.InitializeHandlers(cfg, cmdHandler, log)
	clients.StartHandlers(tgHandler, discordHandler, web, log)
}
