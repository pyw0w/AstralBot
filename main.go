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
	// Define command-line flags
	runWebOnly := flag.Bool("web-only", false, "Run only the web server")
	runBotOnly := flag.Bool("bot-only", false, "Run only the bot")
	flag.Parse()

	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.DebugMode)
	// Initialize the database
	// handlers.InitDB(cfg, log)

	internal.CheckForNewVersion(log)
	web := clients.InitializeWebServer(cfg, log)
	cmdHandler := clients.InitializeCommandHandler(log, cfg.DebugMode)
	tgHandler, discordHandler := clients.InitializeHandlers(cfg, cmdHandler, log)

	if *runWebOnly {
		clients.StartHandlers(nil, nil, web, log)
		utils.WaitForShutdown(nil, log)
		return
	}

	if *runBotOnly {
		clients.StartHandlers(tgHandler, discordHandler, nil, log)
		utils.WaitForShutdown(nil, log)
		return
	}

	// If no arguments are provided, run everything
	if !*runWebOnly && !*runBotOnly {
		clients.StartHandlers(tgHandler, discordHandler, web, log)
		utils.WaitForShutdown(nil, log)
	}
}
