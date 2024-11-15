package steam

import (
	"AstralBot/internal/steam"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SteamInfoCommand struct{}

func (c *SteamInfoCommand) Name() string {
	return "steaminfo"
}

func (c *SteamInfoCommand) Description() string {
	return "Получение информации о пользователе Steam"
}

func (c *SteamInfoCommand) Execute(args []string) (interface{}, error) {
	if len(args) < 1 {
		return "❌ Пожалуйста, укажите Steam ID или URL профиля.", nil
	}

	steamID := args[0]

	if strings.HasPrefix(steamID, "https://steamcommunity.com/id/") {
		steamID = strings.TrimPrefix(steamID, "https://steamcommunity.com/id/")
		steamID = strings.TrimSuffix(steamID, "/")
	}

	if strings.HasPrefix(steamID, "https://steamcommunity.com/profiles/") {
		steamID = strings.TrimPrefix(steamID, "https://steamcommunity.com/profiles/")
		steamID = strings.TrimSuffix(steamID, "/")
	}

	if len(steamID) <= 17 {
		var err error
		steamID, err = steam.ConvertToNumericSteamID(steamID)
		if err != nil {
			return "", err
		}
	}

	gameCount, err := steam.GetOwnedGamesCount(steamID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество игр: %d", gameCount), nil
}

func (c *SteamInfoCommand) ExecuteDiscord(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")[1:]
	response, _ := c.Execute(args)
	s.ChannelMessageSend(m.ChannelID, response.(string))
}

func (c *SteamInfoCommand) ExecuteTelegram(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	args := strings.Split(update.Message.Text, " ")[1:]
	response, _ := c.Execute(args)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response.(string))
	bot.Send(msg)
}
