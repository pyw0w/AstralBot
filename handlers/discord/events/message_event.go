package events

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// MessageEvent представляет событие сообщения
type MessageEvent struct {
	Username string
	Content  string
	Time     time.Time
}

// LogMessage логирует событие сообщения
func LogMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	event := MessageEvent{
		Username: m.Author.Username,
		Content:  m.Content,
		Time:     time.Now(),
	}
	fmt.Printf("Message from %s: %s | Time: %s\n", event.Username, event.Content, event.Time.Format(time.RFC3339))
}
