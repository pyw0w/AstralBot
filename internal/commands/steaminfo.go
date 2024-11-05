package commands

import (
	"AstralBot/internal/steam"
	"fmt"
)

func RegisterSteamInfoCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "steam",
		Description: "Получить информацию о профиле Steam",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌ Пожалуйста, укажите Steam ID или URL профиля.", nil
			}

			steamID := args[0]
			// Проверка, является ли steamID буквенной формой и преобразование в цифровую
			if len(steamID) <= 17 { // Пример длины буквенного SteamID
				var err error
				steamID, err = steam.ConvertToNumericSteamID(steamID)
				if err != nil {
					return "", err
				}
			}

			gameCount, err := steam.GetOwnedGamesCount(steamID)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("Количество игр: %d", gameCount), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
