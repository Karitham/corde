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

type CreateOption struct {
	Name        string           `json:"name"`
	Type        OptionType       `json:"type"`
	Description string           `json:"description,omitempty"`
	Required    bool             `json:"required,omitempty"`
	Options     []CreateOptioner `json:"options,omitempty"`
	Choices     []Choice[any]    `json:"choices,omitempty"`
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

// StringOption is an option that is a string
type StringOption struct {
	Name        string
	Description string
	Required    bool
	Choices     []Choice[string]
}

// NewStringOption returns a new string option
func NewStringOption(name string, description string, required bool, choices ...Choice[string]) *StringOption {
	return &StringOption{
		Name:        name,
		Description: description,
		Required:    required,
		Choices:     choices,
	}
}

func (o *StringOption) createOption() CreateOption {
	return CreateOption{
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
		Type:        OPTION_STRING,
	}
}

func (o *StringOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}

// IntOption represents an option that is an integer
type IntOption struct {
	Name        string
	Description string
	Required    bool
	Choices     []Choice[int]
}

// NewIntOption returns a new integer option
func NewIntOption(name string, description string, required bool, choices ...Choice[int]) *IntOption {
	return &IntOption{
		Name:        name,
		Description: description,
		Required:    required,
		Choices:     choices,
	}
}

func (o *IntOption) createOption() CreateOption {
	return CreateOption{
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
		Type:        OPTION_INTEGER,
	}
}

func (o *IntOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}

// BoolOption is an option that is a boolean
type BoolOption struct {
	Name        string
	Description string
	Required    bool
	Choices     []Choice[bool]
}

// NewBoolOption returns a new boolean option
func NewBoolOption(name string, description string, required bool, choices ...Choice[bool]) *BoolOption {
	return &BoolOption{
		Name:        name,
		Description: description,
		Required:    required,
		Choices:     choices,
	}
}

func (o *BoolOption) createOption() CreateOption {
	return CreateOption{
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
		Type:        OPTION_BOOLEAN,
	}
}

func (o *BoolOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}

// SubcommandOption is an option that is a subcommand
type SubcommandOption struct {
	Name        string
	Description string
	Required    bool
	Options     []CreateOptioner
}

// NewSubcommand returns a new subcommand
func NewSubcommand(name string, description string, required bool, options ...CreateOptioner) *SubcommandOption {
	return &SubcommandOption{
		Name:        name,
		Description: description,
		Required:    required,
		Options:     options,
	}
}

func (o *SubcommandOption) createOption() CreateOption {
	return CreateOption{
		Options:     o.Options,
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
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
	Required    bool
	Options     []CreateOptioner
}

// NewSubcommandGroup returns a new subcommand group
func NewSubcommandGroup(name string, description string, required bool, options ...CreateOptioner) *SubcommandGroupOption {
	return &SubcommandGroupOption{
		Name:        name,
		Description: description,
		Required:    required,
		Options:     options,
	}
}

func (o *SubcommandGroupOption) createOption() CreateOption {
	return CreateOption{
		Options:     o.Options,
		Name:        o.Name,
		Description: o.Description,
		Required:    o.Required,
		Type:        OPTION_SUB_COMMAND_GROUP,
	}
}

func (o *SubcommandGroupOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.createOption())
}
