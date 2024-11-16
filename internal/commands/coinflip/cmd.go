package coinflip

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CoinFlipCommand struct{}

func (c *CoinFlipCommand) Name() string {
	return "coinflip"
}

func (c *CoinFlipCommand) Description() string {
	return "Игра орёл и решка для двух пользователей"
}

func (c *CoinFlipCommand) Execute(args []string) (interface{}, error) {
	if len(args) < 1 {
		return "❌ Пожалуйста, укажите пользователя для игры.", nil
	}
	return "Игра началась! Ожидаем подтверждения от второго пользователя.", nil
}

var reactionMessageID string

func (c *CoinFlipCommand) ExecuteDiscord(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")[1:]
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, "❌ Пожалуйста, укажите пользователя для игры.")
		return
	}

	targetUserID := args[0]

	response, _ := c.Execute(args)
	s.ChannelMessageSend(m.ChannelID, response.(string))

	// Отправляем сообщение с реакцией для подтверждения
	msg, _ := s.ChannelMessageSend(m.ChannelID, "Пользователь, подтвердите участие, нажав на реакцию ниже.")
	s.MessageReactionAdd(m.ChannelID, msg.ID, "✅")
	reactionMessageID = msg.ID

	// Добавляем обработчик реакции
	s.AddHandler(func(s *discordgo.Session, r discordgo.MessageReactionAdd) {
		if r.UserID == s.State.User.ID {
			return
		}
		if r.MessageID == reactionMessageID {
			if r.UserID != targetUserID {
				s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
				return
			}
			if r.Emoji.Name == "✅" {
				coinFlipCmd := &CoinFlipCommand{}
				result := coinFlipCmd.flipCoin()
				s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("Результат: %s", result))
			}
		}
	})

}

func (c *CoinFlipCommand) ExecuteTelegram(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	response, _ := c.Execute(nil)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.(string))
	bot.Send(msg)
}

func (c *CoinFlipCommand) flipCoin() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return "Орёл"
	}
	return "Решка"
}
