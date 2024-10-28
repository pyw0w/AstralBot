package commands

import (
	"time"
)

func RegisterPingCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "ping",
		Description: "Проверка работоспособности бота",
		Execute: func(args []string) (string, error) {
			startTime := time.Now()
			// Здесь можно добавить логику проверки работоспособности бота
			time.Sleep(1 * time.Millisecond) // Задержка для имитации проверки
			duration := time.Since(startTime)
			return "Pong! Время выполнения: " + duration.String(), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
