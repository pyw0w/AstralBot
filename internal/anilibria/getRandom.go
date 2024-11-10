package anilibria

import (
	"AstralBot/internal/httpclient"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTitle() (string, error) {
	client := httpclient.NewClient()

	resp, err := client.Get("https://api.anilibria.tv/v3/title/random")
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неудачный HTTP статус: %d", resp.StatusCode)
	}

	var result any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("ошибка форматирование JSON: %w", err)
	}

	return string(prettyJSON), nil
}
