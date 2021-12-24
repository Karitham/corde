package corde

type ButtonStyle int

const (
	BUTTON_PRIMARY ButtonStyle = iota + 1
	BUTTON_SECONDARY
	BUTTON_SUCCESS
	BUTTON_DANGER
	BUTTON_LINK
)

// ComponentType
// https://discord.com/developers/docs/interactions/message-components#button-object
type ComponentType int

const (
	COMPONENT_ACTION_ROW ComponentType = iota + 1
	COMPONENT_BUTTON
	COMPONENT_SELECT_MENU
)

// Component
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
type Component struct {
	Type        ComponentType `json:"type"`
	CustomID    string        `json:"custom_id,omitempty"`
	Style       ButtonStyle   `json:"style,omitempty"`
	Disabled    bool          `json:"disabled,omitempty"`
	Label       string        `json:"label,omitempty"`
	Emoji       *Emoji        `json:"emoji,omitempty"`
	URL         string        `json:"url,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	MinValues   int           `json:"min_values,omitempty"`
	MaxValues   int           `json:"max_values,omitempty"`
	Options     []Option      `json:"options,omitempty"`
	Components  []Component   `json:"components,omitempty"`
}

// SelectOption
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default,omitempty"`
}

// Button
// https://discord.com/developers/docs/interactions/message-components#button-object
type Button struct {
	Type     ComponentType `json:"type"`
	Style    ButtonStyle   `json:"style"`
	Label    string        `json:"label,omitempty"`
	Emoji    Emoji         `json:"emoji,omitempty"`
	CustomID string        `json:"custom_id,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

type Emoji struct {
	ID   Snowflake `json:"id"`
	Name string    `json:"name"`
}
