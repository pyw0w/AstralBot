package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken      string
	DiscordToken       string
	DebugMode          bool
	CommandPrefix      string
	DetailedAPILogs    bool
	SteamAPIKey        string
	WebPort            int
	DBType             string // Тип базы данных (например, "mongodb", "sqlite", "mysql")
	DBConnectionString string // Строка подключения к базе данных
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return &Config{
		TelegramToken:      os.Getenv("TELEGRAM_TOKEN"),
		DiscordToken:       os.Getenv("DISCORD_TOKEN"),
		DebugMode:          os.Getenv("DEBUG_MODE") == "true",
		CommandPrefix:      os.Getenv("COMMAND_PREFIX"),
		DetailedAPILogs:    os.Getenv("DETAILED_API_LOGS") == "true",
		SteamAPIKey:        os.Getenv("STEAM_API_KEY"),
		WebPort:            8080,
		DBType:             os.Getenv("DB_TYPE"),
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}
}
