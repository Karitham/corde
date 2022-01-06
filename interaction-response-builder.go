package corde

import (
	"fmt"
	"io"
)

// RespB is an InteractionRespData builder
type RespB struct {
	resp *InteractionRespData
}

// NewResp Returns a new response builder
func NewResp() *RespB {
	return &RespB{resp: &InteractionRespData{}}
}

// Embedder returns an Embed
type Embedder interface {
	Embed() Embed
}

// InteractionRespData implements InteractionResponder
func (r *RespB) InteractionRespData() *InteractionRespData { return r.resp }

// Content adds the content to the InteractionRespData
func (r *RespB) Content(s string) *RespB {
	r.resp.Content = s
	return r
}

// Contentf adds the content to the InteractionRespData
func (r *RespB) Contentf(s string, args ...any) *RespB {
	r.resp.Content = fmt.Sprintf(s, args...)
	return r
}

// TTS adds the tts to the InteractionRespData
func (r *RespB) TTS(tts bool) *RespB {
	r.resp.TTS = tts
	return r
}

// Embeds adds embeds to the InteractionRespData
func (r *RespB) Embeds(e ...Embedder) *RespB {
	for _, eb := range e {
		r.resp.Embeds = append(r.resp.Embeds, eb.Embed())
	}
	return r
}

// AllowedMentions adds the allowed mentions to the InteractionRespData
func (r *RespB) AllowedMentions(a *AllowedMentions) *RespB {
	r.resp.AllowedMentions = a
	return r
}

// Flags adds the flags to the InteractionRespData
func (r *RespB) Flags(i IntResponseFlags) *RespB {
	r.resp.Flags = i
	return r
}

// Ephemeral adds the ephemeral flag to the InteractionRespData
func (r *RespB) Ephemeral() *RespB {
	r.resp.Flags = RESPONSE_FLAGS_EPHEMERAL
	return r
}

// Components adds components to the InteractionRespData
func (b *RespB) Components(c ...Component) *RespB {
	if b.resp.Components == nil {
		b.resp.Components = []Component{}
	}

	b.resp.Components = append(b.resp.Components, c...)
	return b
}

// ActionRow adds an action row to the InteractionRespData
func (b *RespB) ActionRow(c ...Component) *RespB {
	if b.resp.Components == nil {
		b.resp.Components = []Component{}
	}

	b.resp.Components = append(b.resp.Components,
		Component{
			Type:       COMPONENT_ACTION_ROW,
			Components: c,
		})
	return b
}

// Attachments adds attachments to the InteractionRespData
func (b *RespB) Attachments(a ...Attachment) *RespB {
	if b.resp.Attachments == nil {
		b.resp.Attachments = []Attachment{}
	}

	b.resp.Attachments = append(b.resp.Attachments, a...)
	return b
}

// Attachement adds an attachment to the InteractionRespData
func (b *RespB) Attachment(body io.Reader, filename string) *RespB {
	if b.resp.Attachments == nil {
		b.resp.Attachments = []Attachment{}
	}

	b.resp.Attachments = append(b.resp.Attachments, Attachment{
		Body:     body,
		Filename: filename,
	})
	return b
}

// Choice adds a choice to the InteractionRespData
func (b *RespB) Choice(name string, value any) *RespB {
	if b.resp.Choices == nil {
		b.resp.Choices = []Choice[any]{}
	}

	b.resp.Choices = append(b.resp.Choices, Choice[any]{
		Name:  name,
		Value: value,
	})

	return b
}

// Choices adds choices to the InteractionRespData
func (b *RespB) Choices(c ...Choice[any]) *RespB {
	if b.resp.Choices == nil {
		b.resp.Choices = []Choice[any]{}
	}

	b.resp.Choices = append(b.resp.Choices, c...)
	return b
}
