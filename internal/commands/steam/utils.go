package steam

import (
	"AstralBot/internal/cmd"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RegisterSteamInfoCommand(cmdHandler *cmd.CommandHandler) {
	cmdHandler.RegisterCommand(&SteamInfoCommand{})
}

func (c *SteamInfoCommand) SendEmbed(s *discordgo.Session, m *discordgo.MessageCreate, embed *discordgo.MessageEmbed) {
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func (c *SteamInfoCommand) SendTelegramMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
