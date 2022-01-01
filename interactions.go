package corde

import (
	"encoding/json"
	"strconv"
	"time"
)

// InteractionType is the type of interaction
type InteractionType int

const (
	PING InteractionType = iota + 1
	APPLICATION_COMMAND
	MESSAGE_COMPONENT
	APPLICATION_COMMAND_AUTOCOMPLETE
)

// Snowflake is a Discord snowflake ID
type Snowflake uint64

// SnowflakeFromString returns a Snowflake from a string
func SnowflakeFromString(s string) Snowflake {
	i, _ := strconv.ParseUint(s, 10, 64)
	return Snowflake(i)
}

// String implements fmt.Stringer
func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

// MarshalJSON implements json.Marshaler
func (s Snowflake) MarshalJSON() ([]byte, error) {
	b := strconv.FormatUint(uint64(s), 10)
	return json.Marshal(b)
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Snowflake) UnmarshalJSON(b []byte) error {
	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	*s = Snowflake(i)
	return nil
}

// User is a Discord User
type User struct {
	Avatar        string    `json:"avatar"`
	Discriminator string    `json:"discriminator"`
	ID            Snowflake `json:"id"`
	PublicFlags   int       `json:"public_flags"`
	Username      string    `json:"username"`
}

// Interaction is a Discord Interaction
type Interaction struct {
	Data          InteractionData `json:"data"`
	GuildID       Snowflake       `json:"guild_id"`
	ChannelID     Snowflake       `json:"channel_id"`
	Member        Member          `json:"member"`
	Message       Message         `json:"message"`
	ApplicationID Snowflake       `json:"application_id"`
	ID            Snowflake       `json:"id"`
	Token         string          `json:"token"`
	Type          InteractionType `json:"type"`
	User          User            `json:"user"`
	Version       int             `json:"version"`
}

// Message is a Discord Message
type Message struct {
	Type            int          `json:"type"`
	Tts             bool         `json:"tts"`
	Timestamp       time.Time    `json:"timestamp"`
	Pinned          bool         `json:"pinned"`
	MentionEveryone bool         `json:"mention_everyone"`
	ID              string       `json:"id"`
	Flags           int          `json:"flags"`
	Embeds          []Embed      `json:"embeds"`
	EditedTimestamp *time.Time   `json:"edited_timestamp"`
	Content         string       `json:"content"`
	Components      []Component  `json:"components"`
	ChannelID       string       `json:"channel_id"`
	Author          Author       `json:"author"`
	Attachments     []Attachment `json:"attachments"`
}

// MessageAuthor is a Discord MessageAuthor
type MessageAuthor struct {
	Username      string `json:"username"`
	PublicFlags   int    `json:"public_flags"`
	ID            string `json:"id"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
}

// InteractionData is data from a Discord Interaction
type InteractionData struct {
	ID            Snowflake           `json:"id"`
	Name          string              `json:"name"`
	Type          int                 `json:"type"`
	Resolved      any                 `json:"resolved"` // ?
	Options       OptionsInteractions `json:"options"`
	CustomID      string              `json:"custom_id"`
	ComponentType ComponentType       `json:"component_type"`
	Values        []any               `json:"values"` // ?
	TagetID       Snowflake           `json:"target_id"`
}

// OptionsInteractions is the options for an Interaction
type OptionsInteractions map[string]any

// UnmarshalJSON implements json.Unmarshaler
func (o *OptionsInteractions) UnmarshalJSON(b []byte) error {
	type opt struct {
		Name    string     `json:"name"`
		Value   any        `json:"value"`
		Type    OptionType `json:"type"`
		Options []opt      `json:"options"`
	}

	var opts []opt
	if err := json.Unmarshal(b, &opts); err != nil {
		return err
	}

	// max is 3 deep, as per discord's docs
	m := make(map[string]any)
	for _, opt := range opts {
		m[opt.Name] = opt.Value
		for _, opt2 := range opt.Options {
			m[opt2.Name] = opt2.Value
			for _, opt3 := range opt2.Options {
				m[opt3.Name] = opt3.Value
			}
		}
	}

	*o = m
	return nil
}

// MarshalJSON implements json.Marshaler
func (o OptionsInteractions) MarshalJSON() ([]byte, error) {
	type opt struct {
		Name  string `json:"name"`
		Value any    `json:"value"`
	}

	opts := make([]opt, len(o))
	for k, v := range o {
		opts = append(opts, opt{k, v})
	}
	b, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Member is a Discord Member
type Member struct {
	User         User       `json:"user"`
	Roles        []string   `json:"roles"`
	PremiumSince *Snowflake `json:"premium_since"`
	Permissions  Snowflake  `json:"permissions"`
	Pending      bool       `json:"pending"`
	Nick         *string    `json:"nick"`
	Mute         bool       `json:"mute"`
	JoinedAt     time.Time  `json:"joined_at"`
	IsPending    bool       `json:"is_pending"`
	Deaf         bool       `json:"deaf"`
}
