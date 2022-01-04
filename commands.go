package corde

import (
	"github.com/Karitham/corde/internal/rest"
)

type OptionType int

const (
	OPTION_SUB_COMMAND OptionType = iota + 1
	OPTION_SUB_COMMAND_GROUP
	OPTION_STRING
	OPTION_INTEGER
	OPTION_BOOLEAN
	OPTION_USER
	OPTION_CHANNEL
	OPTION_ROLE
	OPTION_MENTIONABLE
	OPTION_NUMBER
)

type CommandType int

const (
	COMMAND_CHAT_INPUT CommandType = iota + 1
	COMMAND_USER
	COMMAND_MESSAGE
)

// Command is a Discord application command
type Command struct {
	Name              string      `json:"name,omitempty"`
	ID                Snowflake   `json:"id,omitempty"`
	Type              CommandType `json:"type,omitempty"`
	ApplicationID     Snowflake   `json:"application_id,omitempty"`
	GuildID           Snowflake   `json:"guild_id,omitempty"`
	Description       string      `json:"description,omitempty"`
	Options           []Option    `json:"options,omitempty"`
	DefaultPermission bool        `json:"default_permission,omitempty"`
	Version           Snowflake   `json:"version,omitempty"`
}

// Option is an option for an application Command
type Option struct {
	Name        string        `json:"name"`
	Type        OptionType    `json:"type"`
	Value       any           `json:"value"`
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Options     []Option      `json:"options,omitempty"`
	Choices     []Choice[any] `json:"choices,omitempty"`
	Focused     bool          `json:"focused,omitempty"`
}

// Choice is an application Command choice
type Choice[T any] struct {
	Name  string `json:"name"`
	Value T      `json:"value"`
}

// CommandsOpt is an option for a Command
type CommandsOpt struct {
	guildID Snowflake
}

// GuildOpt is an option for setting the guild of a Command
func GuildOpt(guildID Snowflake) func(*CommandsOpt) {
	return func(opt *CommandsOpt) {
		opt.guildID = guildID
	}
}

// GetCommands returns a slice of Command from the Mux
func (m *Mux) GetCommands(options ...func(*CommandsOpt)) ([]Command, error) {
	opt := &CommandsOpt{}
	for _, option := range options {
		option(opt)
	}

	r := rest.Req("applications", m.AppID)
	if opt.guildID != 0 {
		r.Append("guilds", opt.guildID)
	}
	r.Append("commands")

	var commands []Command
	_, err := rest.DoJson(m.Client, r.Get(m.authorize, rest.JSON), &commands)
	if err != nil {
		return nil, err
	}

	return commands, nil
}

// RegisterCommand registers a new Command on discord
func (m *Mux) RegisterCommand(c CreateCommander, options ...func(*CommandsOpt)) error {
	opt := &CommandsOpt{}
	for _, option := range options {
		option(opt)
	}

	r := rest.Req("applications", m.AppID).JSONBody(c)
	if opt.guildID != 0 {
		r.Append("guilds", opt.guildID)
	}
	r.Append("commands")

	resp, err := m.Client.Do(r.Post(m.authorize, rest.JSON))
	if err != nil {
		return err
	}
	return rest.CodeBetween(resp, 200, 299)
}

// BulkRegisterCommand registers a slice of Command on discord
func (m *Mux) BulkRegisterCommand(c []CreateCommander, options ...func(*CommandsOpt)) error {
	opt := &CommandsOpt{}
	for _, option := range options {
		option(opt)
	}

	r := rest.Req("applications", m.AppID).JSONBody(c)
	if opt.guildID != 0 {
		r.Append("guilds", opt.guildID)
	}
	r.Append("commands")

	resp, err := m.Client.Do(r.Put(m.authorize, rest.JSON))
	if err != nil {
		return err
	}
	return rest.CodeBetween(resp, 200, 299)
}

// DeleteCommand deletes a Command from discord
func (m *Mux) DeleteCommand(ID Snowflake, options ...func(*CommandsOpt)) error {
	opt := &CommandsOpt{}
	for _, option := range options {
		option(opt)
	}

	r := rest.Req("applications", m.AppID)
	if opt.guildID != 0 {
		r.Append("guilds", opt.guildID)
	}
	r.Append("commands", ID)

	resp, err := m.Client.Do(r.Delete(m.authorize, rest.JSON))
	if err != nil {
		return err
	}

	return rest.ExpectCode(resp, 204)
}
