package corde

// ButtonStyle is the style of a button Component
type ButtonStyle int

const (
	BUTTON_PRIMARY   ButtonStyle = iota + 1 // BUTTON_PRIMARY blurple
	BUTTON_SECONDARY                        // BUTTON_SECONDARY grey
	BUTTON_SUCCESS                          // BUTTON_SUCCESS green
	BUTTON_DANGER                           // BUTTON_DANGER red
	BUTTON_LINK                             // BUTTON_LINK grey, navigate to URL
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

// Emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	ID   Snowflake `json:"id"`
	Name string    `json:"name"`
}
