package corde

// Embed builder
// https://regex101.com/r/gmVH2A/3
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

func (b *EmbedB) B() Embed { return b.Embed }

func (b *EmbedB) Author(a Author) *EmbedB {
	b.Embed.Author = a
	return b
}

func (b *EmbedB) Footer(f Footer) *EmbedB {
	b.Embed.Footer = f
	return b
}

func (b *EmbedB) Title(s string) *EmbedB {
	b.Embed.Title = s
	return b
}

func (b *EmbedB) Description(s string) *EmbedB {
	b.Embed.Description = s
	return b
}

func (b *EmbedB) Thumbnail(i Image) *EmbedB {
	b.Embed.Thumbnail = i
	return b
}

func (b *EmbedB) Image(i Image) *EmbedB {
	b.Embed.Image = i
	return b
}

func (b *EmbedB) URL(s string) *EmbedB {
	b.Embed.URL = s
	return b
}

func (b *EmbedB) Fields(f ...Field) *EmbedB {
	b.Embed.Fields = append(b.Embed.Fields, f...)
	return b
}

func (b *EmbedB) Color(i int64) *EmbedB {
	b.Embed.Color = i
	return b
}
