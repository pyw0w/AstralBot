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
	l.Info("Discord-Event", "Команда: "+event.Content+" | "+event.Username)
}
