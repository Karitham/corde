package corde

import (
	"fmt"
	"io"
)

// RespB is an InteractionRespData builder
// https://regex101.com/r/tKfloG/2
type RespB struct {
	*InteractionRespData
}

// NewResp Returns a new response builder
func NewResp() *RespB {
	return &RespB{
		InteractionRespData: &InteractionRespData{
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

// B returns the build InteractionRespData
func (b *RespB) B() *InteractionRespData { return b.InteractionRespData }

// Content adds the content to the InteractionRespData
func (b *RespB) Content(s string) *RespB {
	b.InteractionRespData.Content = s
	return b
}

// Contentf adds the content to the InteractionRespData
func (b *RespB) Contentf(s string, args ...any) *RespB {
	b.InteractionRespData.Content = fmt.Sprintf(s, args...)
	return b
}

// TTS adds the tts to the InteractionRespData
func (b *RespB) TTS(tts bool) *RespB {
	b.InteractionRespData.TTS = tts
	return b
}

// Embeds adds embeds to the InteractionRespData
func (b *RespB) Embeds(e ...Embed) *RespB {
	b.InteractionRespData.Embeds = append(b.InteractionRespData.Embeds, e...)
	return b
}

// AllowedMentions adds the allowed mentions to the InteractionRespData
func (b *RespB) AllowedMentions(a *AllowedMentions) *RespB {
	b.InteractionRespData.AllowedMentions = a
	return b
}

// Flags adds the flags to the InteractionRespData
func (b *RespB) Flags(i IntResponseFlags) *RespB {
	b.InteractionRespData.Flags = i
	return b
}

// Ephemeral adds the ephemeral flag to the InteractionRespData
func (b *RespB) Ephemeral() *RespB {
	b.InteractionRespData.Flags = RESPONSE_FLAGS_EPHEMERAL
	return b
}

// Components adds components to the InteractionRespData
func (b *RespB) Components(c ...Component) *RespB {
	b.InteractionRespData.Components = append(b.InteractionRespData.Components, c...)
	return b
}

// ActionRow adds an action row to the InteractionRespData
func (b *RespB) ActionRow(c ...Component) *RespB {
	b.InteractionRespData.Components = append(b.InteractionRespData.Components,
		Component{
			Type:       COMPONENT_ACTION_ROW,
			Components: c,
		},
	)
	return b
}

// Attachments adds attachments to the InteractionRespData
func (b *RespB) Attachments(a ...Attachment) *RespB {
	b.InteractionRespData.Attachments = append(b.InteractionRespData.Attachments, a...)
	return b
}

// Attachement adds an attachment to the InteractionRespData
func (b *RespB) Attachment(body io.Reader, filename string) *RespB {
	b.InteractionRespData.Attachments = append(b.InteractionRespData.Attachments, Attachment{
		Body:     body,
		Filename: filename,
	})
	return b
}
