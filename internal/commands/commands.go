package commands

type Command interface {
	Name() string
	Description() string
	Execute(args []string) (string, error)
}

type CommandHandler struct {
	Commands map[string]Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		Commands: make(map[string]Command),
	}
}

func (h *CommandHandler) RegisterCommand(cmd Command) {
	h.Commands[cmd.Name()] = cmd
}

func (h *CommandHandler) ExecuteCommand(name string, args []string) (string, error) {
	if cmd, exists := h.Commands[name]; exists {
		return cmd.Execute(args)
	}
	return "❌ Неверная команда", nil
}
