package components

import (
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

// Interaction is a Discord Interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interactions
type Interaction struct {
	ID            snowflake.Snowflake `json:"id"`
	ApplicationID snowflake.Snowflake `json:"application_id"`
	Type          InteractionType     `json:"type"`
	Data          JsonRaw             `json:"data,omitempty"`
	Route         string              `json:"route,omitempty"`
	GuildID       snowflake.Snowflake `json:"guild_id,omitempty"`
	ChannelID     snowflake.Snowflake `json:"channel_id,omitempty"`
	Member        Member              `json:"member,omitempty"`
	User          *User               `json:"user,omitempty"`
	Token         string              `json:"token"`
	Version       int                 `json:"version"`
	Message       *Message            `json:"message,omitempty"`
	Locale        string              `json:"locale,omitempty"`
	GuildLocale   string              `json:"guild_locale,omitempty"`
}

type (
	InteractionDataConstraint interface {
		ButtonInteractionData |
			SelectInteractionData |
			ModalInteractionData |
			UserCommandInteractionData |
			MessageCommandInteractionData |
			SlashInteractionData |
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

	SlashInteractionData struct {
		ID   snowflake.Snowflake `json:"id"`
		Name string              `json:"name"`
		Type int                 `json:"type"`
		resolvedInteractionWithOptions
	}
)

func GetInteractionData[T InteractionDataConstraint](i Interaction) (T, error) {
	var data T
	return data, i.Data.UnmarshalTo(&data)
}
