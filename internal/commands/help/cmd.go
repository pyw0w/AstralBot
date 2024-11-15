package help

import (
	"AstralBot/internal/cmd"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HelpCommand struct {
	CommandHandler *cmd.CommandHandler
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Показывает список доступных команд"
}
func (c *HelpCommand) Execute(args []string) (interface{}, error) {
	var fields []*discordgo.MessageEmbedField
	for _, command := range c.CommandHandler.Commands {
		if command.Name() == "help" {
			continue
		}
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   command.Name(),
			Value:  command.Description(),
			Inline: false,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Доступные команды",
		Description: "Список всех доступных команд:",
		Fields:      fields,
		Color:       0x00ff00, // Green color
	}

	return embed, nil
}

func (c *HelpCommand) ExecuteDiscord(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, _ := c.Execute(nil)
	if embed, ok := response.(*discordgo.MessageEmbed); ok {
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	} else {
		s.ChannelMessageSend(m.ChannelID, response.(string))
	}
}

func (c *HelpCommand) ExecuteTelegram(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	response, _ := c.Execute(nil)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.(string))
	bot.Send(msg)
}
