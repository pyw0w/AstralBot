package internal

import (
	"AstralBot/internal/logger"
	"encoding/json"
	"net/http"
	"time"
)

// Version хранит текущую версию приложения
const Version = "1.1.2" // Обновите версию на нужную

// GitHubVersionURL is the URL to the JSON file containing the latest version information
const GitHubVersionURL = "https://raw.githubusercontent.com/pyw0w/AstralBot/refs/heads/main/version.json"

// VersionInfo представляет структуру JSON файла
type VersionInfo struct {
	Version string `json:"version"`
}

// CheckForNewVersion проверяет, доступна ли новая версия
func CheckForNewVersion(log *logger.Logger) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(GitHubVersionURL)
	if err != nil {
		log.Error("AstralBot-Version", "Ошибка при получении информации о версии: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("AstralBot-Version", "Ошибка: получен код ответа, отличный от 200")
		return
	}

	var versionInfo VersionInfo
	if err := json.NewDecoder(resp.Body).Decode(&versionInfo); err != nil {
		log.Error("AstralBot-Version", "Ошибка при декодировании информации о версии: ", err)
		return
	}

	if versionInfo.Version != Version {
		log.Infof("AstralBot-Version", "Доступна новая версия: %s (текущая версия: %s)\n", versionInfo.Version, Version)
	} else {
		log.Info("AstralBot-Version", "Вы используете последнюю версию: ", Version)
	}
}
