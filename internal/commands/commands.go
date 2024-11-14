package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Description() string
	Execute(args []string) (interface{}, error)
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

func (e *Embed) ToDiscordEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       e.Title,
		Description: e.Description,
		Color:       e.Color,
	}

}

func (h *CommandHandler) ExecuteCommand(name string, args []string) (interface{}, error) {
	if cmd, exists := h.Commands[name]; exists {
		result, err := cmd.Execute(args)
		if err != nil {
			return nil, err
		}
		if embed, ok := result.(Embed); ok {
			return &embed, nil
		}
		return result, nil
	}
	return "❌ Неверная команда", nil
}
