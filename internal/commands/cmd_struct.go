package commands

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

type CommandHandler struct {
	Commands map[string]Command
}
