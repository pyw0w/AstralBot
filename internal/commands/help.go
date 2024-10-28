package commands

import (
	"strings"
)

func RegisterHelpCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "help",
		Description: "Показать доступные команды",
		Execute: func(args []string) (string, error) {
			var availableCommands []string
			for _, command := range cmdHandler.commands {
				availableCommands = append(availableCommands, command.Name+": "+command.Description)
			}
			return "Доступные команды:\n" + strings.Join(availableCommands, "\n"), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
