package commands

import (
	"AstralBot/internal/anilibria"
	"AstralBot/internal/anilibria/structs"
	"encoding/json"
)

type TestCommand struct{}

func (c *TestCommand) Name() string {
	return "test"
}

func (c *TestCommand) Description() string {
	return "Команда для проверки возможности получения данных с Anilibria"
}

func (c *TestCommand) Execute(args []string) (string, error) {
	if len(args) < 1 {
		titleData, err := anilibria.GetTitleRandom()
		if err != nil {
			return "Ошибка получения данных: " + err.Error(), nil
		}

		var title structs.Title
		if err := json.Unmarshal([]byte(titleData), &title); err != nil {
			return "Ошибка декодирования данных: " + err.Error(), nil
		}

		return "Название тайтла: " + title.Names.Ru, nil
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
}

func RegisterTestCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&TestCommand{})
}
