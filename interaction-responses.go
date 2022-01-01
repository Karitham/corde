package corde

import "io"

// InteractionRespData is the payload for responding to an interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionRespData struct {
	Content         string           `json:"content,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           IntResponseFlags `json:"flags,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []Attachment     `json:"attachments,omitempty"`
}

type IntResponseFlags uint

const (
	RESPONSE_FLAGS_EPHEMERAL IntResponseFlags = 64
)

type AllowedMentions struct {
	Parse []string `json:"parse"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

// Attachement is the files attached to the request
type Attachment struct {
	Body        io.Reader `json:"-"`
	ID          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	Description string    `json:"description,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Size        int       `json:"size,omitempty"`
	URL         string    `json:"url,omitempty"`
	ProxyURL    string    `json:"proxy_url,omitempty"`
	Height      int       `json:"height,omitempty"`
	Width       int       `json:"width,omitempty"`
	Ephemeral   bool      `json:"ephemeral,omitempty"`
}

// Embed is the embed object
type Embed struct {
	Author      Author  `json:"author"`
	Footer      Footer  `json:"footer"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumbnail   Image   `json:"thumbnail"`
	Image       Image   `json:"image"`
	URL         string  `json:"url"`
	Fields      []Field `json:"fields"`
	Color       int64   `json:"color"`
}

// Author is the author object
type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

// Field is the field object inside an embed
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Footer is the footer of the embed
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// Image is an image possibly contained inside the embed
type Image struct {
	URL string `json:"url"`
}

type MessageReference struct {
	MessageID string `json:"message_id"`
}
