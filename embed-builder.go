package corde

import (
	"fmt"
	"time"
)

// EmbedB is an Embed builder
// https://regex101.com/r/gmVH2A/4
type EmbedB struct {
	embed Embed
}

// NewEmbed returns a new embed builder ready for use
func NewEmbed() *EmbedB {
	return &EmbedB{
		embed: Embed{
			Fields: []Field{},
		},
	}
}

// Embed returns the built Embed
func (e *EmbedB) Embed() Embed { return e.embed }

// InteractionRespData implements InteractionResponder
func (e *EmbedB) InteractionRespData() *InteractionRespData {
	return &InteractionRespData{
		Embeds: []Embed{e.Embed()},
	}
}

// Author adds the author to the Embed
func (e *EmbedB) Author(a Author) *EmbedB {
	e.embed.Author = a
	return e
}

// Footer adds the footer to the Embed
func (e *EmbedB) Footer(f Footer) *EmbedB {
	e.embed.Footer = f
	return e
}

// Title adds the title to the Embed
func (e *EmbedB) Title(s string) *EmbedB {
	e.embed.Title = s
	return e
}

// Titlef adds the Title to the Embed
func (e *EmbedB) Titlef(format string, a ...any) *EmbedB {
	e.embed.Title = fmt.Sprintf(format, a...)
	return e
}

// Description adds the description to the Embed
func (e *EmbedB) Description(s string) *EmbedB {
	e.embed.Description = s
	return e
}

// Descriptionf adds the description to the Embed
func (e *EmbedB) Descriptionf(format string, a ...any) *EmbedB {
	e.embed.Description = fmt.Sprintf(format, a...)
	return e
}

// Thumbnail adds the thumbnail to the Embed
func (e *EmbedB) Thumbnail(i Image) *EmbedB {
	e.embed.Thumbnail = i
	return e
}

// Image adds the image to the Embed
func (e *EmbedB) Image(i Image) *EmbedB {
	e.embed.Image = i
	return e
}

// ImageURL adds an image based off the url to the Embed
func (e *EmbedB) ImageURL(s string) *EmbedB {
	e.embed.Image = Image{
		URL: s,
	}
	return e
}

// URL adds the url to the Embed
func (e *EmbedB) URL(s string) *EmbedB {
	e.embed.URL = s
	return e
}

// Fields append the field to the Embed
func (e *EmbedB) Fields(f ...Field) *EmbedB {
	e.embed.Fields = append(e.embed.Fields, f...)
	return e
}

// Field adds a field to the Embed
func (e *EmbedB) Field(name, value string) *EmbedB {
	e.embed.Fields = append(e.embed.Fields, Field{
		Name:  name,
		Value: value,
	})
	return e
}

// FieldInline adds an inline field to the Embed
func (e *EmbedB) FieldInline(name, value string) *EmbedB {
	e.embed.Fields = append(e.embed.Fields, Field{
		Name:   name,
		Value:  value,
		Inline: true,
	})
	return e
}

// Provider adds a provider to the Embed
func (e *EmbedB) Provider(name string, url string) *EmbedB {
	e.embed.Provider = Provider{
		Name: name,
		URL:  url,
	}
	return e
}

// Video adds the video to the Embed
func (e *EmbedB) Video(v Video) *EmbedB {
	e.embed.Video = v
	return e
}

// Timestamp adds the timestamp to the Embed
func (e *EmbedB) Timestamp(t time.Time) *EmbedB {
	e.embed.Timestamp = opt(Timestamp(t))
	return e
}

// Color adds the color to the Embed
func (e *EmbedB) Color(i uint32) *EmbedB {
	e.embed.Color = i
	return e
}

// Message returns the embed wrapped in a message
func (e *EmbedB) Message() Message {
	return Message{
		Embeds: []Embed{e.embed},
	}
}

func opt[T any](v T) *T {
	return &v
}
