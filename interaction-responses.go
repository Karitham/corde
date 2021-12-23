package corde

type InteractionResponseData struct {
	Content         *string                  `json:"content,omitempty"`
	TTS             bool                     `json:"tts,omitempty"`
	Embeds          []Embed                  `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions         `json:"allowed_mentions,omitempty"`
	Flags           InteractionResponseFlags `json:"flags,omitempty"`
}

type InteractionResponseFlags uint

const (
	RESPONSE_FLAGS_EPHEMERAL InteractionResponseFlags = 64
)

type AllowedMentions struct {
	Parse []string `json:"parse"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       Image     `json:"image"`
}

type MessageReference struct {
	MessageID string `json:"message_id"`
}

type Attachments struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}

// Opt returns a ptr to the type
func Opt[T any](t T) *T {
	return &t
}
