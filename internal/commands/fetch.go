package commands

import (
	"AstralBot/internal/httpclient"
	"fmt"
	"net/http"
)

func RegisterFetchCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "fetch",
		Description: "Получить данные из внешнего API",
		Execute: func(args []string) (string, error) {
			client := httpclient.NewClient()                        // Передаем логгер и флаг дебага
			resp, err := client.Get("https://api.example.com/data") // Замените на ваш URL
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return "", fmt.Errorf("ошибка при получении данных: %s", resp.Status)
			}

			return "Данные успешно получены!", nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
