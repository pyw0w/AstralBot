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

func (c *HelpCmd) generateCommandsList() []Field {
	var fields []Field
	for _, cmd := range c.cmdHandler.Commands {
		fields = append(fields, Field{Name: cmd.Name(), Value: cmd.Description()})
	}
	return fields
}

func (c *HelpCmd) Execute(args []string) (interface{}, error) {
	embed := Embed{
		Title:       "Доступные команды",
		Description: "Список всех доступных команд:",
		Color:       65280,
		Fields:      c.generateCommandsList(),
	}
	return embed, nil
}

func RegisterHelpCommand(cmdHandler *CommandHandler) {
	cmdHandler.RegisterCommand(&HelpCmd{cmdHandler: cmdHandler})
}
