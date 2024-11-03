package events

import (
	"fmt"
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
func LogMessage(update tgbotapi.Update) {
	event := MessageEvent{
		Username: update.Message.From.UserName,
		Content:  update.Message.Text,
		Time:     time.Now(),
	}
	fmt.Printf("Message from %s: %s | Time: %s\n", event.Username, event.Content, event.Time.Format(time.RFC3339))
}
