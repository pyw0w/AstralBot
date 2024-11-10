package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken   string
	DiscordToken    string
	DebugMode       bool
	CommandPrefix   string
	DetailedAPILogs bool
	SteamAPIKey     string // Добавлено новое поле для API ключа Steam
	WebPort         int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return &Config{
		TelegramToken:   os.Getenv("TELEGRAM_TOKEN"),
		DiscordToken:    os.Getenv("DISCORD_TOKEN"),
		DebugMode:       os.Getenv("DEBUG_MODE") == "true",
		CommandPrefix:   os.Getenv("COMMAND_PREFIX"),
		DetailedAPILogs: os.Getenv("DETAILED_API_LOGS") == "true",
		SteamAPIKey:     os.Getenv("STEAM_API_KEY"), // Загрузка API ключа
		WebPort:         8080,                       // Установите порт по умолчанию для веб-сервера
	}
}
