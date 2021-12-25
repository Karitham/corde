package corde

import "fmt"

// ResponseBuilder
// https://regex101.com/r/tKfloG/1
type RespB struct {
	*InteractionRespData
}

// Returns a new response builder
func NewResp() *RespB {
	return &RespB{
		InteractionRespData: &InteractionRespData{
			Content:         "",
			TTS:             false,
			Embeds:          []Embed{},
			AllowedMentions: &AllowedMentions{},
			Flags:           0,
			Components:      []Component{},
			Attachements:    []Attachment{},
		},
	}
}
func (b *RespB) B() *InteractionRespData { return b.InteractionRespData }

func (b *RespB) Content(s string) *RespB {
	b.InteractionRespData.Content = s
	return b
}

func (b *RespB) Contentf(s string, args ...any) *RespB {
	b.InteractionRespData.Content = fmt.Sprintf(s, args...)
	return b
}

func (b *RespB) TTS(tts bool) *RespB {
	b.InteractionRespData.TTS = tts
	return b
}

func (b *RespB) Embeds(e ...Embed) *RespB {
	b.InteractionRespData.Embeds = append(b.InteractionRespData.Embeds, e...)
	return b
}

func (b *RespB) AllowedMentions(a *AllowedMentions) *RespB {
	b.InteractionRespData.AllowedMentions = a
	return b
}

func (b *RespB) Flags(i IntResponseFlags) *RespB {
	b.InteractionRespData.Flags = i
	return b
}

func (b *RespB) Ephemeral() *RespB {
	b.InteractionRespData.Flags = RESPONSE_FLAGS_EPHEMERAL
	return b
}

func (b *RespB) Components(c ...Component) *RespB {
	b.InteractionRespData.Components = append(b.InteractionRespData.Components, c...)
	return b
}

func (b *RespB) ActionRow(c ...Component) *RespB {
	b.InteractionRespData.Components = append(b.InteractionRespData.Components,
		Component{
			Type:       COMPONENT_ACTION_ROW,
			Components: c,
		},
	)
	return b
}

func (b *RespB) Attachements(a ...Attachment) *RespB {
	b.InteractionRespData.Attachements = append(b.InteractionRespData.Attachements, a...)
	return b
}
