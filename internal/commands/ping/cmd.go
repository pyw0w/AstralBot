package ping

import (
	"time"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Проверка работоспособности бота"
}

func (c *PingCommand) Execute(args []string) (interface{}, error) {
	startTime := time.Now()
	time.Sleep(1 * time.Millisecond) // Задержка для имитации проверки
	duration := time.Since(startTime)
	return "Pong! Время выполнения: " + duration.String(), nil
}

func (c *PingCommand) ExecuteDiscord(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, _ := c.Execute(nil)
	s.ChannelMessageSend(m.ChannelID, response.(string))
}

func (c *PingCommand) ExecuteTelegram(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	response, _ := c.Execute(nil)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.(string))
	bot.Send(msg)
}
