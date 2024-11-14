package commands

type HelpCmd struct {
	cmdHandler *CommandHandler
}

func (c *HelpCmd) Name() string {
	return "help"
}

func (c *HelpCmd) Description() string {
	return "Показать доступные команды"
}

func (c *HelpCmd) Execute(args []string) (interface{}, error) {
	embed := Embed{
		Title:       "Доступные команды",
		Description: "Список всех доступных команд:",
		Color:       65280,
		Fields: []Field{
			{Name: "test", Value: "Команда для проверки возможности получения данных с Anilibria"},
			{Name: "ping", Value: "Проверка работоспособности бота"},
			{Name: "help", Value: "Показать доступные команды"},
			{Name: "steam", Value: "Получить информацию о профиле Steam"},
		},
	}
	return embed, nil
}

func RegisterHelpCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&HelpCmd{cmdHandler: cmdHandler})
}
