package commands

import (
	"time"
)

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Проверка работоспособности бота"
}

func (c *PingCommand) Execute(args []string) (interface{}, error) {
	startTime := time.Now()
	time.Sleep(1 * time.Millisecond) // Задержка для имитации проверки
	duration := time.Since(startTime)
	return "Pong! Время выполнения: " + duration.String(), nil
}

func RegisterPingCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&PingCommand{})
}
