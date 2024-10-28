package anilibria

import (
	"AstralBot/internal/structs/anilibria" // Импортируем структуры
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTitleByID(id int) (*anilibria.Title, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.anilibria.tv/v3/title?id=%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка при получении данных: %s", resp.Status)
	}

	var title anilibria.Title
	if err := json.NewDecoder(resp.Body).Decode(&title); err != nil {
		return nil, err
	}

	return &title, nil
}

func GetTitleByCode(code string) (*anilibria.Title, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.anilibria.tv/v3/title?code=%s", code))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка при получении данных: %s", resp.Status)
	}

	var title anilibria.Title
	if err := json.NewDecoder(resp.Body).Decode(&title); err != nil {
		return nil, err
	}

	return &title, nil
}
