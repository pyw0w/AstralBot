package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramToken    string `yaml:"telegram_token"`
	DiscordToken     string `yaml:"discord_token"`
	DebugMode        bool   `yaml:"debug_mode"`
	CommandPrefix    string `yaml:"command_prefix"`
	DetailedAPILogs  bool   `yaml:"detailed_api_logs"`
	SteamAPIKey      string `yaml:"steam_api_key"`
	WebPort          int    `yaml:"web_port"`
	DiscordChannelID string `yaml:"discord_channel_id"`
	AccessToken      string `yaml:"access_token"`

	// Database
	DBType     string `yaml:"db_type"`
	DBLogin    string `yaml:"db_login"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
}

func LoadConfig() *Config {
	file, err := os.ReadFile("config.yml")

	if err != nil {
		defaultConfig := Config{
			TelegramToken:    "your_telegram_token",
			DiscordToken:     "your_discord_token",
			DebugMode:        false,
			CommandPrefix:    "!",
			DetailedAPILogs:  false,
			SteamAPIKey:      "your_steam_api_key",
			WebPort:          8080,
			DiscordChannelID: "your_discord_channel_id",
			AccessToken:      "your_access_token",
			DBType:           "mysql",
			DBLogin:          "root",
			DBPassword:       "password",
			DBHost:           "localhost",
			DBPort:           "3306",
			DBName:           "astralbot",
		}

		data, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			log.Fatal("Ошибка создания config.yml файла: ", err)
		}

		err = os.WriteFile("config.yml", data, 0644)
		if err != nil {
			log.Fatal("Ошибка записи config.yml файла: ", err)
		}

		return &defaultConfig
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Ошибка парсинга config.yml файла: ", err)
	}

	return &config
}
