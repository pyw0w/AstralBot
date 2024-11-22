package run

import (
	"github.com/pyw0w/AstralBot/database"
	"github.com/pyw0w/AstralBot/discord"
	"github.com/pyw0w/AstralBot/pluginBase"
)

func init() {
	pluginBase.RegisterPlugin(&discord.Plugin{})
	pluginBase.RegisterPlugin(&database.Plugin{})
}
