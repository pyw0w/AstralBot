package commands

func RegisterTestCommand(cmdHandler *CommandHandler) {
	cmd := Command{
		Name:        "test",
		Description: "",
		Execute: func(args []string) (string, error) {
			return "Test!", nil
		},
	}

	// Регистрация команды
	cmdHandler.RegisterCommand(cmd)
}
