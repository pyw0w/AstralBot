package commands

import (
	"strings"
)

type HelpCommand struct {
	cmdHandler *CommandHandler
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Показать доступные команды"
}

func (c *HelpCommand) Execute(args []string) (string, error) {
	var availableCommands []string
	for _, command := range c.cmdHandler.Commands {
		availableCommands = append(availableCommands, command.Name()+": "+command.Description())
	}
	return "Доступные команды:\n" + strings.Join(availableCommands, "\n"), nil
}

func RegisterHelpCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&HelpCommand{cmdHandler: cmdHandler})
}
