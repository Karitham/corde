package corde

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
	// components
	ActionRowInteraction InnerInteractionType = iota + 1
	ButtonInteraction
	SelectMenuInteraction
	TextInputInteraction

	// autocomplete
	AutocompleteInteraction

	// commands
	SlashCommandInteraction
	UserCommandInteraction
	MessageCommandInteraction
)

// Interaction is a Discord Interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interactions
type Interaction[T InteractionDataConstraint] struct {
	ID            Snowflake       `json:"id"`
	ApplicationID Snowflake       `json:"application_id"`
	Type          InteractionType `json:"type"`
	Data          T               `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id,omitempty"`
	ChannelID     Snowflake       `json:"channel_id,omitempty"`
	Member        Member          `json:"member,omitempty"`
	User          *User           `json:"user,omitempty"`
	Token         string          `json:"token"`
	Version       int             `json:"version"`
	Message       *Message        `json:"message,omitempty"`
	Locale        string          `json:"locale,omitempty"`
	GuildLocale   string          `json:"guild_locale,omitempty"`

	Route                string               `json:"-"`
	InnerInteractionType InnerInteractionType `json:"-"`
}

type (
	_basicT struct {
		Type InteractionType `json:"type"`
		Data []byte          `json:"data,omitempty"`
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

type (
	// InteractionDataConstraint is the constraint for the interaction data
	// It contains all the possible values for interaction data
	InteractionDataConstraint interface {
		JsonRaw |
			ButtonInteractionData |
			SelectInteractionData |
			TextInputInteractionData |
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
		ID      Snowflake           `json:"id"`
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

	TextInputInteractionData struct {
		CustomID    string      `json:"custom_id"`
		Title       string      `json:"title"`
		Style       int         `json:"style"`
		Label       string      `json:"label"`
		MinLenght   int         `json:"min_length,omitempty"`
		MaxLenght   int         `json:"max_length,omitempty"`
		Required    bool        `json:"required,omitempty"`
		Value       string      `json:"value,omitempty"`
		Placeholder string      `json:"placeholder,omitempty"`
		Components  []Component `json:"components"`
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
		ID       Snowflake `json:"id"`
		TargetID Snowflake `json:"target_id,omitempty"`
		Name     string    `json:"name"`
		Type     int       `json:"type"`
		resolvedInteractionWithOptions
	}

	MessageCommandInteractionData struct {
		ID       Snowflake `json:"id"`
		TargetID Snowflake `json:"target_id,omitempty"`
		Name     string    `json:"name"`
		Type     int       `json:"type"`
		resolvedInteractionWithOptions
	}

	SlashCommandInteractionData struct {
		ID   Snowflake `json:"id"`
		Name string    `json:"name"`
		Type int       `json:"type"`
		resolvedInteractionWithOptions
	}

	PartialRoutingType struct {
		ID            Snowflake `json:"id"`
		Type          int       `json:"type"`
		ComponentType int       `json:"component_type"`
		Name          string    `json:"name"`
		CustomID      string    `json:"custom_id"`
		resolvedInteractionWithOptions
	}
)
