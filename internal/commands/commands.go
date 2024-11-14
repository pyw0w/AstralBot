package commands

import "github.com/bwmarrin/discordgo"

type Embed struct {
	Title       string
	Description string
	Color       int
	Fields      []Field
	Inline      bool
}

type Field struct {
	Name   string
	Value  string
	Inline bool
}

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
		Fields:      e.FieldsToDiscordFields(),
	}
}

func (e *Embed) FieldsToDiscordFields() []*discordgo.MessageEmbedField {
	var fields []*discordgo.MessageEmbedField
	for _, field := range e.Fields {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   field.Name,
			Value:  field.Value,
			Inline: field.Inline,
		})
	}
	return fields
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
