package discord

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/pyw0w/AstralBot/logger"
)

type Plugin struct{}

func (p *Plugin) Init() error {
	return nil
}

func (p *Plugin) Run() error {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		logger.Error("Discord", "DISCORD_BOT_TOKEN is not set")
		return nil
	}

	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
	)
	if err != nil {
		return err
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		return err
	}

	logger.Info("discord", "Discord bot is running")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s

	return nil
}
