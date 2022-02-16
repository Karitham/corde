package components

import (
	"encoding/json"
	"path"
	"strings"

	"github.com/Karitham/corde/snowflake"
)

const (
	// RouteInteractionSubcommandGroup represents the map key for a subcommand group route
	RouteInteractionSubcommandGroup = "$group"
	// RouteInteractionSubcommand reprensents the map key for a subcommand route
	RouteInteractionSubcommand = "$command"

	// RouteInteractionFocused represents the map key for a focused route.
	// This is useful for autocomplete interactions so we can route on focused keys
	// Such, we route on `$group/$command/$focused`,
	RouteInteractionFocused = "$focused"
)

// InteractionType is the type of interaction
type InteractionType int

const (
	INTERACTION_TYPE_PING InteractionType = iota + 1
	INTERACTION_TYPE_APPLICATION_COMMAND
	INTERACTION_TYPE_MESSAGE_COMPONENT
	INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE
)

// InnerInteractionType is the inner type of interactions,
// and not just command, component, autocomplete etc.
type InnerInteractionType int

const (
	ButtonInteraction InnerInteractionType = iota + 1

	SelectInteraction
	ModalInteraction

	AutocompleteInteraction

	SlashCommandInteraction
	UserCommandInteraction
	MessageCommandInteraction
)

// Interaction is a Discord Interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interactions
type Interaction[DataT InteractionDataConstraint] struct {
	ID            snowflake.Snowflake `json:"id"`
	ApplicationID snowflake.Snowflake `json:"application_id"`
	Type          InteractionType     `json:"type"`
	Data          DataT               `json:"data,omitempty"`
	GuildID       snowflake.Snowflake `json:"guild_id,omitempty"`
	ChannelID     snowflake.Snowflake `json:"channel_id,omitempty"`
	Member        Member              `json:"member,omitempty"`
	User          *User               `json:"user,omitempty"`
	Token         string              `json:"token"`
	Version       int                 `json:"version"`
	Message       *Message            `json:"message,omitempty"`
	Locale        string              `json:"locale,omitempty"`
	GuildLocale   string              `json:"guild_locale,omitempty"`

	Route                string               `json:"-"`
	InnerInteractionType InnerInteractionType `json:"-"`
}

type (
	_basicT struct {
		Type InteractionType `json:"type"`
		Data JsonRaw         `json:"data,omitempty"`
	}

	_appCommandT struct {
		Type    int                 `json:"type"`
		Name    string              `json:"name"`
		Options OptionsInteractions `json:"options"`
	}

	_messageComponentT struct {
		Type    int                 `json:"type"`
		Name    string              `json:"name"`
		Options OptionsInteractions `json:"options"`
	}
)

func (i *Interaction[JsonRaw]) UnmarshalJSON(b []byte) error {
	// TODO(@Karitham): Finish unmarshalling stuff

	var bt _basicT
	if err := json.Unmarshal(b, &bt); err != nil {
		return err
	}

	switch bt.Type {
	case INTERACTION_TYPE_PING:
	case INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE:
		var data AutocompleteInteractionData
		if err := json.Unmarshal(bt.Data, &data); err != nil {
			return err
		}

		group := data.Options[RouteInteractionSubcommandGroup]
		cmd := data.Options[RouteInteractionSubcommand]
		focused := data.Options[RouteInteractionFocused]

		i.InnerInteractionType = AutocompleteInteraction

		i.Route = path.Join(data.Name, group.String(), cmd.String(), focused.String())
	case INTERACTION_TYPE_APPLICATION_COMMAND:
		i.Type = INTERACTION_TYPE_APPLICATION_COMMAND

		var data _appCommandT
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}

		switch data.Type {
		case 1:
			i.InnerInteractionType = SlashCommandInteraction
		case 2:
			i.InnerInteractionType = UserCommandInteraction
		case 3:
			i.InnerInteractionType = MessageCommandInteraction
		default:
			i.InnerInteractionType = SlashCommandInteraction
		}

		if data.Type != 1 {
			data.Name = path.Join(strings.Fields(data.Name)...)
		}

		group := data.Options[RouteInteractionSubcommandGroup]
		cmd := data.Options[RouteInteractionSubcommand]
		i.Route = path.Join(data.Name, group.String(), cmd.String())
	case INTERACTION_TYPE_MESSAGE_COMPONENT:
		i.Type = INTERACTION_TYPE_MESSAGE_COMPONENT

		var data _messageComponentT
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}
	}

	// This should 100% be valid
	// If we remove it, we get a compiler crash
	// I haven't been able to figure out why, and I can't manage to reproduce this, nor make sense of it.
	// It seems the compiler is having issues with rewriting the IR for generic types, notably assignments
	i.Data = bt.Data

	// Error is
	// # github.com/Karitham/corde/components
	// components/interactions.go:152:11: cannot use bt.Data (variable of type JsonRaw) as type JsonRaw in assignment

	return nil
}

type (
	// InteractionDataConstraint is the constraint for the interaction data
	// It contains all the possible values for interaction data
	InteractionDataConstraint interface {
		JsonRaw |
			ButtonInteractionData |
			SelectInteractionData |
			ModalInteractionData |
			UserCommandInteractionData |
			MessageCommandInteractionData |
			SlashCommandInteractionData |
			AutocompleteInteractionData |
			PartialCommandInteraction
	}

	resolvedInteractionWithOptions struct {
		Resolved Resolved            `json:"resolved,omitempty"`
		Options  OptionsInteractions `json:"options,omitempty"`
	}

	AutocompleteInteractionData struct {
		ID      snowflake.Snowflake `json:"id"`
		Name    string              `json:"name"`
		Type    int                 `json:"type"`
		Options OptionsInteractions `json:"options,omitempty"`
	}

	ButtonInteractionData struct {
		CustomID      string        `json:"custom_id,omitempty"`
		ComponentType ComponentType `json:"component_type"`
	}

	SelectInteractionData struct {
		CustomID      string        `json:"custom_id,omitempty"`
		ComponentType ComponentType `json:"component_type"`
		Values        []any         `json:"values,omitempty"`
	}

	ModalInteractionData struct {
		CustomID   string      `json:"custom_id,omitempty"`
		Components []Component `json:"components,omitempty"`
	}

	PartialCommandInteraction struct {
		Type int `json:"type"`
		JsonRaw
	}

	UserCommandInteractionData struct {
		ID       snowflake.Snowflake `json:"id"`
		TargetID snowflake.Snowflake `json:"target_id,omitempty"`
		Name     string              `json:"name"`
		Type     int                 `json:"type"`
		resolvedInteractionWithOptions
	}

	MessageCommandInteractionData struct {
		ID       snowflake.Snowflake `json:"id"`
		TargetID snowflake.Snowflake `json:"target_id,omitempty"`
		Name     string              `json:"name"`
		Type     int                 `json:"type"`
		resolvedInteractionWithOptions
	}

	SlashCommandInteractionData struct {
		ID   snowflake.Snowflake `json:"id"`
		Name string              `json:"name"`
		Type int                 `json:"type"`
		resolvedInteractionWithOptions
	}
)
