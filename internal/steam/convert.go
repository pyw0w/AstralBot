package steam

import (
	"AstralBot/internal/httpclient"
	"AstralBot/utils/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func ConvertToNumericSteamID(steamID string) (string, error) {
	client := httpclient.NewClient()
	cfg := config.LoadConfig() // Загрузка конфигурации
	resp, err := client.Get("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=" + cfg.SteamAPIKey + "&vanityurl=" + steamID)
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неудачный HTTP статус: %d", resp.StatusCode)
	}

	var result struct {
		Response struct {
			SteamID string `json:"steamid"`
			Success int    `json:"success"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	if result.Response.Success != 1 {
		return "", fmt.Errorf("не удалось преобразовать SteamID")
	}

	return result.Response.SteamID, nil
}
