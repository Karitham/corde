package corde

import (
	"io"
	"strings"
	"time"
)

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
	Choices         []Choice[any]    `json:"choices,omitempty"`
}

// InteractionRespData implements InteractionResponder interface
func (i *InteractionRespData) InteractionRespData() *InteractionRespData {
	return i
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
// https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	Timestamp   *Timestamp `json:"timestamp,omitempty"`
	Color       uint32     `json:"color,omitempty"`
	Footer      Footer     `json:"footer,omitempty"`
	Image       Image      `json:"image,omitempty"`
	Thumbnail   Image      `json:"thumbnail,omitempty"`
	Video       Video      `json:"video,omitempty"`
	Provider    Provider   `json:"provider,omitempty"`
	Author      Author     `json:"author,omitempty"`
	Fields      []Field    `json:"fields,omitempty"`
}

// Embed implements Embedder
func (e Embed) Embed() Embed {
	return e
}

// InteractionRespData implements InteractionResponder
func (e Embed) InteractionRespData() *InteractionRespData {
	return &InteractionRespData{
		Embeds: []Embed{e},
	}
}

// Timestamp is a discord timestamp
// It is represented as a string in the ISO 8601 format
type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).UTC().Format(time.RFC3339) + `"`), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" || s == "" {
		return nil
	}

	v, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	*t = Timestamp(v)
	return nil
}

func (t Timestamp) String() string {
	return time.Time(t).UTC().Format("2006-01-02T15:04:05-0700")
}

// Author is the author object
type Author struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// Field is the field object inside an embed
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Footer is the footer of the embed
type Footer struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// Image is an image possibly contained inside the embed
type Image struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type MessageReference struct {
	MessageID string `json:"message_id"`
}

// Video is an embed video
type Video struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// Provider is an embed provider
type Provider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
