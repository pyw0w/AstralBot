package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Commands: make(map[string]Command),
	}
}

func (h *CommandHandler) RegisterCommand(cmd Command) {
	h.Commands[cmd.Name()] = cmd
}

func (h *CommandHandler) ExecuteCommand(name string, args []string) (interface{}, error) {
	if cmd, exists := h.Commands[name]; exists {
		return cmd.Execute(args)
	}
	return "❌ Неверная команда", nil
}

func (h *CommandHandler) ExecuteDiscordCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmdName := strings.TrimPrefix(strings.Split(m.Content, " ")[0], "!")
	if cmd, exists := h.Commands[cmdName]; exists {
		cmd.ExecuteDiscord(s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "❌ Неверная команда")
	}
}

func (h *CommandHandler) ExecuteTelegramCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmdName := strings.Split(update.Message.Text, " ")[0]
	if cmd, exists := h.Commands[cmdName]; exists {
		cmd.ExecuteTelegram(bot, update)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ Неверная команда")
		bot.Send(msg)
	}
}
