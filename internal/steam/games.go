package steam

import (
	"AstralBot/internal/httpclient"
	"AstralBot/utils/config"
	"encoding/json"
)

func GetOwnedGamesCount(steamID string) (int, error) {
	client := httpclient.NewClient()
	cfg := config.LoadConfig() // Загрузка конфигурации
	resp, err := client.Get("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + cfg.SteamAPIKey + "&steamid=" + steamID + "&format=json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			GameCount int `json:"game_count"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Response.GameCount, nil
}
