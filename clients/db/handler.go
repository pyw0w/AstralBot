package db

import (
	"AstralBot/internal/database"
	"AstralBot/utils/config"
	"log"
)

func Connect() database.Database {
	cfg := config.LoadConfig()

	db, err := database.NewDatabase(cfg.DBType, cfg.DBName)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	return db
}
