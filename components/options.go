package components

import (
	"encoding/json"

	"github.com/Karitham/corde/snowflake"
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
	Name              string              `json:"name,omitempty"`
	ID                snowflake.Snowflake `json:"id,omitempty"`
	Type              CommandType         `json:"type,omitempty"`
	ApplicationID     snowflake.Snowflake `json:"application_id,omitempty"`
	GuildID           snowflake.Snowflake `json:"guild_id,omitempty"`
	Description       string              `json:"description,omitempty"`
	Options           []Option            `json:"options,omitempty"`
	DefaultPermission bool                `json:"default_permission,omitempty"`
	Version           snowflake.Snowflake `json:"version,omitempty"`
}

// Option is an option for an application Command
type Option struct {
	Name        string        `json:"name"`
	Type        OptionType    `json:"type"`
	Value       JsonRaw       `json:"value"`
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

// OptionsInteractions is the options for an Interaction
type OptionsInteractions map[string]JsonRaw

// UnmarshalJSON implements json.Unmarshaler
func (o *OptionsInteractions) UnmarshalJSON(b []byte) error {
	type opt struct {
		Name    string     `json:"name"`
		Value   JsonRaw    `json:"value"`
		Type    OptionType `json:"type"`
		Options []opt      `json:"options"`
		Focused bool       `json:"focused"`
	}

	var opts []opt
	if err := json.Unmarshal(b, &opts); err != nil {
		return err
	}

	// max is 3 deep, as per discord's docs
	m := make(map[string]JsonRaw)
	for _, opt := range opts {
		switch {
		case OPTION_SUB_COMMAND_GROUP == opt.Type:
			opt.Value = []byte(opt.Name)
			opt.Name = RouteInteractionSubcommandGroup
		case OPTION_SUB_COMMAND == opt.Type:
			opt.Value = []byte(opt.Name)
			opt.Name = RouteInteractionSubcommand
		case opt.Focused:
			m[RouteInteractionFocused] = []byte(opt.Name)
		}

		m[opt.Name] = opt.Value
		for _, opt2 := range opt.Options {
			switch {
			case OPTION_SUB_COMMAND == opt2.Type:
				opt2.Value = []byte(opt2.Name)
				opt2.Name = RouteInteractionSubcommand
			case opt2.Focused:
				m[RouteInteractionFocused] = []byte(opt2.Name)
			}

			m[opt2.Name] = opt2.Value
			for _, opt3 := range opt2.Options {
				if opt3.Focused {
					m[RouteInteractionFocused] = []byte(opt3.Name)
				}
				m[opt3.Name] = opt3.Value
			}
		}
	}

	*o = m
	return nil
}

// MarshalJSON implements json.Marshaler
func (o OptionsInteractions) MarshalJSON() ([]byte, error) {
	type opt struct {
		Name  string `json:"name"`
		Value any    `json:"value"`
	}

	opts := make([]opt, len(o))
	for k, v := range o {
		opts = append(opts, opt{k, v})
	}
	b, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	return b, nil
}
