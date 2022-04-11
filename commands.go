package corde

import (
	"github.com/Karitham/corde/internal/rest"
)

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
	_, err := rest.DoJSON(m.Client, r.Get(m.authorize, rest.JSON), &commands)
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
