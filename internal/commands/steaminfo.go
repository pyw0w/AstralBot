package commands

import (
	"AstralBot/internal/steam"
	"fmt"
	"strings"
)

type SteamInfoCommand struct{}

func (c *SteamInfoCommand) Name() string {
	return "steam"
}

func (c *SteamInfoCommand) Description() string {
	return "Получить информацию о профиле Steam"
}

func (c *SteamInfoCommand) Execute(args []string) (string, error) {
	if len(args) < 1 {
		return "❌ Пожалуйста, укажите Steam ID или URL профиля.", nil
	}

	steamID := args[0]

	// Проверка, является ли steamID URL профиля и извлечение ID
	if strings.HasPrefix(steamID, "https://steamcommunity.com/id/") {
		steamID = strings.TrimPrefix(steamID, "https://steamcommunity.com/id/")
		steamID = strings.TrimSuffix(steamID, "/")
	}

	// Проверка, является ли steamID URL профиля и извлечение ID
	if strings.HasPrefix(steamID, "https://steamcommunity.com/profiles/") {
		steamID = strings.TrimPrefix(steamID, "https://steamcommunity.com/profiles/")
		steamID = strings.TrimSuffix(steamID, "/")
	}

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
}

func RegisterSteamInfoCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&SteamInfoCommand{})
}
