package main

import (
	"flag"

	"github.com/pyw0w/AstralBot/logger"
	"github.com/pyw0w/AstralBot/pluginBase"
	"github.com/pyw0w/AstralBot/web"
)

func main() {
	runWeb := flag.Bool("web", false, "Run web server")
	flag.Parse()

	if *runWeb {
		go web.StartServer()
	}

	logger.Info("main", "Starting plugins...")
	if err := pluginBase.Start(); err != nil {
		logger.Error("main", "Failed to start: "+err.Error())
	}
}
