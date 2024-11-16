package events

import (
	"AstralBot/internal/logger"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// MessageEvent представляет событие сообщения
type MessageEvent struct {
	Username string
	Content  string
	Time     time.Time
}

// LogMessage логирует событие сообщения
func LogMessage(update tgbotapi.Update, l *logger.Logger) {
	event := MessageEvent{
		Username: update.Message.From.UserName,
		Content:  update.Message.Text,
	}

	l.Info("Telegram-Event", "Команда: "+event.Content+" | "+event.Username)
}
