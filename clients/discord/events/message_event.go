package events

import (
	"AstralBot/internal/logger"

	"github.com/bwmarrin/discordgo"
)

// MessageEvent представляет событие сообщения
type MessageEvent struct {
	Username string
	Content  string
}

// LogMessage логирует событие сообщения
func LogMessage(s *discordgo.Session, m *discordgo.MessageCreate, l *logger.Logger) {
	event := MessageEvent{
		Username: m.Author.Username,
		Content:  m.Content,
	}
	if m.Author.Bot {
		return
	}
	l.Info("Discord-Event", "Сообщение: "+event.Content+" | "+event.Username)
}
