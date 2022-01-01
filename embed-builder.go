package corde

import "fmt"

// EmbedB is an Embed builder
// https://regex101.com/r/gmVH2A/4
type EmbedB struct {
	Embed
}

// NewEmbed returns a new embed builder ready for use
func NewEmbed() *EmbedB {
	return &EmbedB{
		Embed: Embed{
			Author:      Author{},
			Footer:      Footer{},
			Title:       "",
			Description: "",
			Thumbnail:   Image{},
			Image:       Image{},
			URL:         "",
			Fields:      []Field{},
			Color:       0,
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

// URL adds the url to the Embed
func (b *EmbedB) URL(s string) *EmbedB {
	b.Embed.URL = s
	return b
}

func (b *EmbedB) Fields(f ...Field) *EmbedB {
	b.Embed.Fields = append(b.Embed.Fields, f...)
	return b
}

// Color adds the color to the Embed
func (b *EmbedB) Color(i int64) *EmbedB {
	b.Embed.Color = i
	return b
}
