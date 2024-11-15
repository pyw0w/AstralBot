package cmd

import (
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command interface {
	Name() string
	Description() string
	Execute(args []string) (interface{}, error)
	ExecuteDiscord(s *discordgo.Session, m *discordgo.MessageCreate)
	ExecuteTelegram(bot *tgbotapi.BotAPI, update tgbotapi.Update)
}
