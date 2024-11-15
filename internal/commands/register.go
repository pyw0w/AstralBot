package commands

import (
	"AstralBot/internal/cmd"
	"AstralBot/internal/commands/ping"
	"AstralBot/internal/commands/steam"
)

func RegisterCommands(cmdHandler *cmd.CommandHandler) {
	ping.RegisterPingCommand(cmdHandler)
	steam.RegisterSteamInfoCommand(cmdHandler)
}
