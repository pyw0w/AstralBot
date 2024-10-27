package commands

type Command struct {
	Name        string
	Description string
	Execute     func(args []string) (string, error)
}

type CommandHandler struct {
	commands map[string]Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		commands: make(map[string]Command),
	}
}

func (h *CommandHandler) RegisterCommand(cmd Command) {
	h.commands[cmd.Name] = cmd
}

func (h *CommandHandler) ExecuteCommand(name string, args []string) (string, error) {
	if cmd, exists := h.commands[name]; exists {
		return cmd.Execute(args)
	}
	return "❌ Неверная команда", nil
}
