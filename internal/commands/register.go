package commands

import (
	"AstralBot/internal/cmd"
	"AstralBot/internal/commands/help"
	"AstralBot/internal/commands/ping"
	"AstralBot/internal/commands/steam"
)

func RegisterCommands(cmdHandler *cmd.CommandHandler) {
	ping.RegisterPingCommand(cmdHandler)
	steam.RegisterSteamInfoCommand(cmdHandler)
	cmdHandler.RegisterCommand(&help.HelpCommand{CommandHandler: cmdHandler})
}
