package commands

import (
	"AstralBot/internal/anilibria"
	"fmt"
)

func RegisterTestCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "test",
		Description: "",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌", nil
			}

			fmt.Print(anilibria.GetTitle(map[string]string{"id": args[0]}))
			return "Test!" + args[0], nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
