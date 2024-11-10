package anilibria

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetSchedule(params map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.anilibria.tv/v3/title/schedule", nil)
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %w", err)
	}

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
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

// Example usage of GetSchedule function
func main() {
	params := map[string]string{
		"filter":           "id,names,description",
		"days":             "monday,tuesday",
		"description_type": "plain",
		"playlist_type":    "object",
	}

	schedule, err := GetSchedule(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(schedule)
}
