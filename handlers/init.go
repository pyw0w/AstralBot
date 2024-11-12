package handlers

import (
	"AstralBot/handlers/discord"
	"AstralBot/handlers/telegram"
	"AstralBot/internal/commands"
	"AstralBot/internal/database"
	"AstralBot/internal/database/models"
	"AstralBot/internal/logger"
	"AstralBot/utils/config"
	"AstralBot/web"
	"os"

	"gorm.io/gorm"
)

func InitializeCommandHandler(log *logger.Logger, debugMode bool) *commands.CommandHandler {
	cmdHandler := commands.NewCommandHandler()
	commands.RegisterCommands(cmdHandler, log, debugMode)
	return cmdHandler
}

func InitializeHandlers(cfg *config.Config, cmdHandler *commands.CommandHandler, log *logger.Logger) (*telegram.Handler, *discord.Handler) {
	tgHandler, err := telegram.NewHandler(cfg.TelegramToken, cmdHandler, cfg.DebugMode, log)
	if err != nil {
		log.Error("System", "Ошибка инициализации Telegram: ", err)
		os.Exit(1)
	}

	discordHandler, err := discord.NewHandler(cfg.DiscordToken, cmdHandler, cfg.DebugMode, log, cfg.DetailedAPILogs)
	if err != nil {
		log.Error("System", "Ошибка инициализации Discord: ", err)
		os.Exit(1)
	}

	return tgHandler, discordHandler
}

func InitializeWebServer(cfg *config.Config, log *logger.Logger) *web.Server {
	webServer := web.NewServer(cfg, log)
	return webServer
}

func InitDB(cfg *config.Config, log *logger.Logger) {
	db, err := database.NewDatabase(cfg.DBType, cfg.DBConnectionString)
	if err != nil {
		log.Error("AstralBot", "Ошибка инициализации базы данных: ", err)
	}

	err = db.Connect()
	if err != nil {
		log.Error("AstralBot", "Ошибка подключения к базе данных: ", err)
	}

	if cfg.DBType != "mongodb" {
		sqlDB := db.GetDB().(*gorm.DB)
		sqlDB.AutoMigrate(&models.User{})
	}
	if cfg.DBType == "sqlite" {
		sqlDB := db.GetDB().(*gorm.DB)
		sqlDB.AutoMigrate(&models.User{})
	}
}

func StartHandlers(tgHandler *telegram.Handler, discordHandler *discord.Handler, webServer *web.Server, log *logger.Logger) {
	go tgHandler.Start()
	go discordHandler.Start()
	go webServer.Start()
	log.Info("AstralBot", "Бот запущен. Нажмите Ctrl+C для завершения")
}
