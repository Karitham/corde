package corde

import "encoding/json"

// Style is the style of a button Component
type Style int

const (
	BUTTON_PRIMARY   Style = iota + 1 // BUTTON_PRIMARY blurple
	BUTTON_SECONDARY                  // BUTTON_SECONDARY grey
	BUTTON_SUCCESS                    // BUTTON_SUCCESS green
	BUTTON_DANGER                     // BUTTON_DANGER red
	BUTTON_LINK                       // BUTTON_LINK grey, navigate to URL

	TEXT_SHORT     = 1
	TEXT_PARAGRAPH = 2
)

// ComponentType
// https://discord.com/developers/docs/interactions/message-components#button-object
type ComponentType int

const (
	COMPONENT_ACTION_ROW ComponentType = iota + 1
	COMPONENT_BUTTON
	COMPONENT_SELECT_MENU
	COMPONENT_TEXT
)

// Component
//
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
type Component struct {
	Type        ComponentType `json:"type"`
	CustomID    string        `json:"custom_id"`
	Style       Style         `json:"style,omitempty"`
	Disabled    bool          `json:"disabled,omitempty"`
	Label       string        `json:"label,omitempty"`
	Emoji       *Emoji        `json:"emoji,omitempty"`
	URL         string        `json:"url,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	MinValues   int           `json:"min_values,omitempty"`
	MaxValues   int           `json:"max_values,omitempty"`
	MinLength   int           `json:"min_length,omitempty"`
	MaxLength   int           `json:"max_length,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Value       string        `json:"value,omitempty"`
	Options     []Option      `json:"options,omitempty"`
	Components  []Component   `json:"components,omitempty"`
}

func (c Component) Component() Component {
	return c
}

// TextInputComponent
//
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInputComponent struct {
	CustomID    string `json:"custom_id"`
	Style       Style  `json:"style"`
	Label       string `json:"label"`
	MinLength   int    `json:"min_length,omitempty"`
	MaxLength   int    `json:"max_length,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Value       string `json:"value,omitempty"`
	Placeholder string `json:"place_holder,omitempty"`
}

func (t TextInputComponent) Component() Component {
	return Component{
		Type:        COMPONENT_TEXT,
		CustomID:    t.CustomID,
		Style:       t.Style,
		Label:       t.Label,
		MinLength:   t.MinLength,
		MaxLength:   t.MaxLength,
		Required:    t.Required,
		Value:       t.Value,
		Placeholder: t.Placeholder,
	}
}

// Modal represents a discord modal.
// when marshalled, it is wrapped in an action row.
// Thus, the components it should contain are the actual components you want displayed,
// rather than an action row wrapping them.
//
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-styles
type Modal struct {
	Title      string      `json:"title"`
	CustomID   string      `json:"custom_id"`
	Components []Component `json:"components"`
}

func (m Modal) MarshalJSON() ([]byte, error) {
	type M2 Modal
	m2 := Modal{
		Title:    m.Title,
		CustomID: m.CustomID,
		Components: []Component{{
			Type:       1,
			Components: m.Components,
		}},
	}

	return json.Marshal(M2(m2))
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
	Style    Style         `json:"style"`
	Label    string        `json:"label,omitempty"`
	Emoji    Emoji         `json:"emoji,omitempty"`
	CustomID string        `json:"custom_id,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

// Emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object
type Emoji struct {
	ID            Snowflake   `json:"id,omitempty"`
	Name          string      `json:"name"`
	Roles         []Snowflake `json:"roles,omitempty"`
	User          User        `json:"user,omitempty"`
	RequireColons bool        `json:"require_colons,omitempty"`
	Managed       bool        `json:"managed,omitempty"`
	Animated      bool        `json:"animated,omitempty"`
	Available     bool        `json:"available,omitempty"`
}
