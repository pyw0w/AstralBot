package commands

import (
	"AstralBot/internal/anilibria"
	"AstralBot/internal/structs/anilibria" // Импортируем структуры
	"fmt"
	"strconv"
)

func RegisterAnilibriaCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "anilibria",
		Description: "Получить информацию о тайтле по ID или коду",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌ Пожалуйста, укажите ID или код тайтла.", nil
			}

			var title *anilibria.Title
			var err error

			if id, err := strconv.Atoi(args[0]); err == nil {
				title, err = anilibria.GetTitleByID(id)
			} else {
				title, err = anilibria.GetTitleByCode(args[0])
			}

			if err != nil {
				return "", err
			}

			return fmt.Sprintf("Тайтл: %s\nСтатус: %s\nОписание: %s\nЖанры: %v", title.Names.Ru, title.Status.String, title.Description, title.Genres), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
