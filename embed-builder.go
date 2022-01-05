package corde

import (
	"fmt"
	"time"
)

// EmbedB is an Embed builder
// https://regex101.com/r/gmVH2A/4
type EmbedB struct {
	Embed
}

// NewEmbed returns a new embed builder ready for use
func NewEmbed() *EmbedB {
	return &EmbedB{
		Embed: Embed{
			Title:       "",
			Description: "",
			URL:         "",
			Timestamp:   Timestamp{},
			Color:       0,
			Footer:      Footer{},
			Image:       Image{},
			Thumbnail:   Image{},
			Video:       Video{},
			Provider:    Provider{},
			Author:      Author{},
			Fields:      []Field{},
		},
	}
}

// B returns the built Embed
func (b *EmbedB) B() Embed { return b.Embed }

// Author adds the author to the Embed
func (b *EmbedB) Author(a Author) *EmbedB {
	b.Embed.Author = a
	return b
}

// Footer adds the footer to the Embed
func (b *EmbedB) Footer(f Footer) *EmbedB {
	b.Embed.Footer = f
	return b
}

// Title adds the title to the Embed
func (b *EmbedB) Title(s string) *EmbedB {
	b.Embed.Title = s
	return b
}

// Titlef adds the Title to the Embed
func (b *EmbedB) Titlef(format string, a ...any) *EmbedB {
	b.Embed.Title = fmt.Sprintf(format, a...)
	return b
}

// Description adds the description to the Embed
func (b *EmbedB) Description(s string) *EmbedB {
	b.Embed.Description = s
	return b
}

// Descriptionf adds the description to the Embed
func (b *EmbedB) Descriptionf(format string, a ...any) *EmbedB {
	b.Embed.Description = fmt.Sprintf(format, a...)
	return b
}

// Thumbnail adds the thumbnail to the Embed
func (b *EmbedB) Thumbnail(i Image) *EmbedB {
	b.Embed.Thumbnail = i
	return b
}

// Image adds the image to the Embed
func (b *EmbedB) Image(i Image) *EmbedB {
	b.Embed.Image = i
	return b
}

// ImageURL adds an image based off the url to the Embed
func (b *EmbedB) ImageURL(s string) *EmbedB {
	b.Embed.Image = Image{
		URL: s,
	}
	return b
}

// URL adds the url to the Embed
func (b *EmbedB) URL(s string) *EmbedB {
	b.Embed.URL = s
	return b
}

// Fields append the field to the Embed
func (b *EmbedB) Fields(f ...Field) *EmbedB {
	b.Embed.Fields = append(b.Embed.Fields, f...)
	return b
}

// Field adds a field to the Embed
func (b *EmbedB) Field(name, value string) *EmbedB {
	b.Embed.Fields = append(b.Embed.Fields, Field{
		Name:  name,
		Value: value,
	})
	return b
}

// FieldInline adds an inline field to the Embed
func (b *EmbedB) FieldInline(name, value string) *EmbedB {
	b.Embed.Fields = append(b.Embed.Fields, Field{
		Name:   name,
		Value:  value,
		Inline: true,
	})
	return b
}

// Provider adds a provider to the Embed
func (b *EmbedB) Provider(name string, url string) *EmbedB {
	b.Embed.Provider = Provider{
		Name: name,
		URL:  url,
	}
	return b
}

// Video adds the video to the Embed
func (b *EmbedB) Video(v Video) *EmbedB {
	b.Embed.Video = v
	return b
}

// Timestamp adds the timestamp to the Embed
func (b *EmbedB) Timestamp(t time.Time) *EmbedB {
	b.Embed.Timestamp = Timestamp(t)
	return b
}

// Color adds the color to the Embed
func (b *EmbedB) Color(i uint32) *EmbedB {
	b.Embed.Color = i
	return b
}
