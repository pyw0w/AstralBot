package commands

import (
	"AstralBot/internal/cmd"
	"AstralBot/internal/commands/coinflip"
	"AstralBot/internal/commands/help"
	"AstralBot/internal/commands/ping"
	"AstralBot/internal/commands/steam"
)

func RegisterCommands(cmdHandler *cmd.CommandHandler) {
	ping.RegisterPingCommand(cmdHandler)
	steam.RegisterSteamInfoCommand(cmdHandler)
	coinflip.RegisterCoinFlipCommand(cmdHandler)
	cmdHandler.RegisterCommand(&help.HelpCommand{CommandHandler: cmdHandler})
}
