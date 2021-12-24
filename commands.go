package corde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

const API = "https://discord.com/api/v8"

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

type Choice[T any] struct {
	Name  string `json:"name"`
	Value T      `json:"value"`
}

type commandsOpt struct {
	guildID Snowflake
}

func Guild(guildID Snowflake) func(*commandsOpt) {
	return func(opt *commandsOpt) {
		opt.guildID = guildID
	}
}

func (m *Mux) GetCommands(options ...func(*commandsOpt)) ([]Command, error) {
	opt := &commandsOpt{}
	for _, option := range options {
		option(opt)
	}

	guild := ""
	if opt.guildID != 0 {
		guild = fmt.Sprintf("/guilds/%d", opt.guildID)
	}

	url := fmt.Sprintf("%s/applications/%d%s/commands", API, m.AppID, guild)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	m.authorize(req)

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commands []Command
	if err := json.NewDecoder(resp.Body).Decode(&commands); err != nil {
		return nil, err
	}

	return commands, nil
}

func (m *Mux) RegisterCommand(c Command, options ...func(*commandsOpt)) error {
	opt := &commandsOpt{}
	for _, option := range options {
		option(opt)
	}

	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(c); err != nil {
		return err
	}

	guild := ""
	if opt.guildID != 0 {
		guild = fmt.Sprintf("/guilds/%s", opt.guildID)
	}

	url := fmt.Sprintf("%s/applications/%d%s/commands", API, m.AppID, guild)

	req, err := http.NewRequest(http.MethodPost, url, b)
	if err != nil {
		return err
	}
	reqOpts(req, m.authorize, contentTypeJSON)

	resp, err := m.Client.Do(req)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		buf := &bytes.Buffer{}
		buf.ReadFrom(resp.Body)
		return fmt.Errorf("error: %w, body: %s, code: %d", err, buf.String(), resp.StatusCode)
	}
	return nil
}
