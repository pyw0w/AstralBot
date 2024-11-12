package commands

import (
	"AstralBot/internal/anilibria"
	"AstralBot/internal/anilibria/structs"
	"encoding/json"
)

func RegisterTestCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "test",
		Description: "",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌", nil
			}

			titleData, err := anilibria.GetTitle(map[string]string{"id": args[0]})
			if err != nil {
				return "Ошибка получения данных: " + err.Error(), nil
			}

			var title structs.Title
			if err := json.Unmarshal([]byte(titleData), &title); err != nil {
				return "Ошибка декодирования данных: " + err.Error(), nil
			}

			return "Название тайтла: " + title.Names.Ru, nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
