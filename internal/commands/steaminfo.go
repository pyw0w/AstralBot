package commands

import (
	"AstralBot/config"
	"AstralBot/internal/httpclient"
	"AstralBot/internal/structs/steam" // Импортируем структуры
	"encoding/json"
	"fmt"
)

func RegisterSteamInfoCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "steaminfo",
		Description: "Получить информацию о профиле Steam",
		Execute: func(args []string) (string, error) {
			if len(args) < 1 {
				return "❌ Пожалуйста, укажите Steam ID или URL профиля.", nil
			}

			steamID := args[0]
			// Проверка, является ли steamID буквенной формой и преобразование в цифровую
			if len(steamID) <= 17 { // Пример длины буквенного SteamID
				var err error
				steamID, err = convertToNumericSteamID(steamID)
				if err != nil {
					return "", err
				}
			}

			client := httpclient.NewClient()
			cfg := config.LoadConfig() // Загрузка конфигурации
			resp, err := client.Get("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + cfg.SteamAPIKey + "&steamids=" + steamID)
			if err != nil {
				return "Ошибка: %s", nil
			}
			defer resp.Body.Close()

			var profile steam.Profile
			if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
				return "", err
			}

			return fmt.Sprintf("Количество игр: %d", profile.Response.PlayerCount), nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}

func convertToNumericSteamID(steamID string) (string, error) {
	client := httpclient.NewClient()
	cfg := config.LoadConfig() // Загрузка конфигурации
	resp, err := client.Get("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=" + cfg.SteamAPIKey + "&vanityurl=" + steamID)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			SteamID string `json:"steamid"`
			Success int    `json:"success"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Response.Success == 1 {
		return result.Response.SteamID, nil
	}
	return "", fmt.Errorf("не удалось преобразовать SteamID")
}
