package corde

import "encoding/json"

// CreateCommander is a command that can be registered
type CreateCommander interface {
	createCommand() CreateCommand
}

// CreateOptioner is an interface for all options
type CreateOptioner interface {
	createOption() CreateOption
}

// CreateOption is the base option type for creating any sort of option
type CreateOption struct {
	Name         string           `json:"name"`
	Type         OptionType       `json:"type"`
	Description  string           `json:"description,omitempty"`
	Required     bool             `json:"required,omitempty"`
	Choices      []Choice[any]    `json:"choices,omitempty"`
	Options      []CreateOptioner `json:"options,omitempty"`
	ChannelTypes []ChannelType    `json:"channel_types,omitempty"`
	MinValue     float64          `json:"min_value,omitempty"`
	MaxValue     float64          `json:"max_value,omitempty"`
	Autocomplete bool             `json:"autocomplete,omitempty"`
}

func (c CreateOption) createOption() CreateOption {
	return c
}

// CreateCommand is a slash command that can be registered to discord
type CreateCommand struct {
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Type        CommandType      `json:"type,omitempty"`
	Options     []CreateOptioner `json:"options,omitempty"`
}

// CreateCommand is a slash command that can be registered to discord
type SlashCommand struct {
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Type        CommandType      `json:"type,omitempty"`
	Options     []CreateOptioner `json:"options,omitempty"`
}

// NewSlashCommand returns a new slash command
func NewSlashCommand(name string, description string, options ...CreateOptioner) CreateCommand {
	return CreateCommand{
		Name:        name,
		Description: description,
		Options:     options,
		Type:        COMMAND_CHAT_INPUT,
	}
}

func (c CreateCommand) createCommand() CreateCommand {
	return CreateCommand{
		Name:        c.Name,
		Description: c.Description,
		Options:     c.Options,
		Type:        c.Type,
	}
}

// SubcommandOption is an option that is a subcommand
type SubcommandOption struct {
	Name        string
	Description string
	Options     []CreateOptioner
}

// NewSubcommand returns a new subcommand
func NewSubcommand(name string, description string, options ...CreateOptioner) *SubcommandOption {
	return &SubcommandOption{
		Name:        name,
		Description: description,
		Options:     options,
	}
}

func (o *SubcommandOption) createOption() CreateOption {
	return CreateOption{
		Options:     o.Options,
		Name:        o.Name,
		Description: o.Description,
		Type:        OPTION_SUB_COMMAND,
	}
}

func (o *SubcommandOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}

// SubcommandGroupOption is an option that is a subcommand group
type SubcommandGroupOption struct {
	Name        string
	Description string
	Options     []CreateOptioner
}

// NewSubcommandGroup returns a new subcommand group
func NewSubcommandGroup(name string, description string, options ...CreateOptioner) *SubcommandGroupOption {
	return &SubcommandGroupOption{
		Name:        name,
		Description: description,
		Options:     options,
	}
}

func (o *SubcommandGroupOption) createOption() CreateOption {
	return CreateOption{
		Options:     o.Options,
		Name:        o.Name,
		Description: o.Description,
		Type:        OPTION_SUB_COMMAND_GROUP,
	}
}

func (o *SubcommandGroupOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}
