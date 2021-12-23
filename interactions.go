package corde

import (
	"encoding/json"
	"strconv"
	"time"
)

type InteractionType int

const (
	PING InteractionType = iota + 1
	APPLICATION_COMMAND
	MESSAGE_COMPONENT
	APPLICATION_COMMAND_AUTOCOMPLETE
)

type Snowflake uint64

func SnowflakeFromString(s string) Snowflake {
	i, _ := strconv.ParseUint(s, 10, 64)
	return Snowflake(i)
}

func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

func (s Snowflake) MarshalJSON() ([]byte, error) {
	b := strconv.FormatUint(uint64(s), 10)
	return json.Marshal(b)
}

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

type User struct {
	Avatar        string    `json:"avatar"`
	Discriminator string    `json:"discriminator"`
	ID            Snowflake `json:"id"`
	PublicFlags   int       `json:"public_flags"`
	Username      string    `json:"username"`
}

type Interaction struct {
	Data          InteractionData `json:"data"`
	GuildID       Snowflake       `json:"guild_id"`
	ChannelID     Snowflake       `json:"channel_id"`
	Member        Member          `json:"member"`
	Message       string          `json:"message"`
	ApplicationID Snowflake       `json:"application_id"`
	ID            Snowflake       `json:"id"`
	Token         string          `json:"token"`
	Type          InteractionType `json:"type"`
	User          User            `json:"user"`
	Version       int             `json:"version"`
}

type InteractionData struct {
	ID            Snowflake     `json:"id"`
	Name          string        `json:"name"`
	Type          int           `json:"type"`
	Resolved      interface{}   `json:"resolved"` // ?
	Options       []Option      `json:"options"`
	CustomID      Snowflake     `json:"custom_id"`
	ComponentType int           `json:"component_type"`
	Values        []interface{} `json:"values"` // ?
	TagetID       Snowflake     `json:"target_id"`
}

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
