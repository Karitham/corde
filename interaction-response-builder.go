package corde

import (
	"fmt"
	"io"
)

// RespB is an InteractionRespData builder
// https://regex101.com/r/tKfloG/2
type RespB struct {
	resp *InteractionRespData
}

// NewResp Returns a new response builder
func NewResp() *RespB {
	return &RespB{
		resp: &InteractionRespData{
			Content:         "",
			TTS:             false,
			Embeds:          []Embed{},
			AllowedMentions: &AllowedMentions{},
			Flags:           0,
			Components:      []Component{},
			Attachments:     []Attachment{},
		},
	}
}

// EmbedBuilder is an Embed builder
type EmbedBuilder interface {
	Embed() Embed
}

// InteractionRespData implements InteractionResponseDataBuilder
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
func (r *RespB) Embeds(e ...EmbedBuilder) *RespB {
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
func (r *RespB) Components(c ...Component) *RespB {
	r.resp.Components = append(r.resp.Components, c...)
	return r
}

// ActionRow adds an action row to the InteractionRespData
func (r *RespB) ActionRow(c ...Component) *RespB {
	r.resp.Components = append(r.resp.Components,
		Component{
			Type:       COMPONENT_ACTION_ROW,
			Components: c,
		},
	)
	return r
}

// Attachments adds attachments to the InteractionRespData
func (r *RespB) Attachments(a ...Attachment) *RespB {
	r.resp.Attachments = append(r.resp.Attachments, a...)
	return r
}

// Attachment adds an attachment to the InteractionRespData
func (r *RespB) Attachment(body io.Reader, filename string) *RespB {
	r.resp.Attachments = append(r.resp.Attachments, Attachment{
		Body:     body,
		Filename: filename,
	})
	return r
}
