package commands

func RegisterCommands(cmdHandler *CommandHandler) {
	RegisterPingCommand(cmdHandler)
	RegisterHelpCommand(cmdHandler)
	RegisterSteamInfoCommand(cmdHandler)
	RegisterTestCommand(cmdHandler)
	RegisterCoinFlipCommand(cmdHandler)
}
