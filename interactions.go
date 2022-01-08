package corde

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	ID              Snowflake    `json:"id"`
	Flags           int          `json:"flags"`
	Embeds          []Embed      `json:"embeds"`
	EditedTimestamp *time.Time   `json:"edited_timestamp"`
	Content         string       `json:"content"`
	Components      []Component  `json:"components"`
	ChannelID       Snowflake    `json:"channel_id"`
	Author          Author       `json:"author"`
	Attachments     []Attachment `json:"attachments"`
}

// MessageAuthor is a Discord MessageAuthor
type MessageAuthor struct {
	Username      string    `json:"username"`
	PublicFlags   int       `json:"public_flags"`
	ID            Snowflake `json:"id"`
	Discriminator string    `json:"discriminator"`
	Avatar        Hash      `json:"avatar"`
}

// Role is a user's role
// https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Permissions uint64    `json:"permissions,string"`
	Position    int       `json:"position"`
	Color       uint32    `json:"color"`
	Hoist       bool      `json:"hoist"`
	Managed     bool      `json:"managed"`
	Mentionable bool      `json:"mentionable"`
}

// InteractionData is data from a Discord Interaction
type InteractionData struct {
	ID            Snowflake           `json:"id"`
	Name          string              `json:"name"`
	Type          int                 `json:"type"`
	Resolved      Resolved            `json:"resolved,omitempty"`
	Options       OptionsInteractions `json:"options,omitempty"`
	CustomID      string              `json:"custom_id,omitempty"`
	ComponentType ComponentType       `json:"component_type"`
	Values        []any               `json:"values,omitempty"`
	TargetID      Snowflake           `json:"target_id,omitempty"`
}

// OptionsUser returns the resolved User for an Option
func (i InteractionData) OptionsUser(k string) (User, error) {
	var u User
	s, err := i.Options.Snowflake(k)
	if err != nil {
		return u, err
	}
	u, ok := i.Resolved.Users[s]
	if !ok {
		return u, fmt.Errorf("no user found for option %q", k)
	}
	return u, nil
}

// OptionsMember returns the resolved Member (and User) for an Option
func (i InteractionData) OptionsMember(k string) (Member, error) {
	var m Member
	s, err := i.Options.Snowflake(k)
	if err != nil {
		return m, err
	}
	m, ok := i.Resolved.Members[s]
	if !ok {
		return m, fmt.Errorf("no member found for option %q", k)
	}

	m.User, err = i.OptionsUser(k)
	if err != nil {
		return m, err
	}
	return m, nil
}

// OptionsRole returns the resolved Role for an Option
func (i InteractionData) OptionsRole(k string) (Role, error) {
	var r Role
	s, err := i.Options.Snowflake(k)
	if err != nil {
		return r, err
	}
	r, ok := i.Resolved.Roles[s]
	if !ok {
		return r, fmt.Errorf("no role found for option %q", k)
	}
	return r, nil
}

// OptionsMessage returns the resolved Message for an Option
func (i InteractionData) OptionsMessage(k string) (Message, error) {
	var m Message
	s, err := i.Options.Snowflake(k)
	if err != nil {
		return m, err
	}
	m, ok := i.Resolved.Messages[s]
	if !ok {
		return m, fmt.Errorf("no member message for option %q", k)
	}
	return m, nil
}

type ResolvedDataConstraint interface {
	User | Member | Role | Message
}

// ResolvedData is a generic mapping of Snowflakes to resolved data structs
type ResolvedData[T ResolvedDataConstraint] map[Snowflake]T

// First returns the first resolved data
func (r ResolvedData[T]) First() T {
	for _, v := range r {
		return v
	}
	return *new(T)
}

type Resolved struct {
	Users    ResolvedData[User]    `json:"users,omitempty"`
	Members  ResolvedData[Member]  `json:"members,omitempty"`
	Roles    ResolvedData[Role]    `json:"roles,omitempty"`
	Messages ResolvedData[Message] `json:"messages,omitempty"`
}

// OptionsInteractions is the options for an Interaction
type OptionsInteractions map[string]JsonRaw

// UnmarshalJSON implements json.Unmarshaler
func (o *OptionsInteractions) UnmarshalJSON(b []byte) error {
	type opt struct {
		Name    string     `json:"name"`
		Value   JsonRaw    `json:"value"`
		Type    OptionType `json:"type"`
		Options []opt      `json:"options"`
	}

	var opts []opt
	if err := json.Unmarshal(b, &opts); err != nil {
		return err
	}

	// max is 3 deep, as per discord's docs
	m := make(map[string]JsonRaw)
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

// Hash is a discord profile picture hash
// https://discord.com/developers/docs/reference#image-formatting
type Hash string

// Animated returns wether the url is an animated gif
func (h Hash) Animated() bool {
	return strings.HasPrefix(string(h), "a_")
}

// AvatarPNG returns a png url of this hash
func (u User) AvatarPNG() string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.png", u.ID, u.Avatar)
}

// AvatarURL returns a url of this user, being animated if it can
func (u User) AvatarURL() string {
	if u.Avatar.Animated() {
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.gif", u.ID, u.Avatar)
	}
	return u.AvatarPNG()
}

// Member is a Discord Member
// https://discord.com/developers/docs/resources/guild#guild-member-object
type Member struct {
	User                       User        `json:"user"`
	Nick                       string      `json:"nick,omitempty"`
	RoleIDs                    []Snowflake `json:"roles"`
	Avatar                     Hash        `json:"avatar,omitempty"`
	Joined                     Timestamp   `json:"joined_at"`
	BoostedSince               Timestamp   `json:"premium_since,omitempty"`
	CommunicationDisabledUntil Timestamp   `json:"communication_disabled_until,omitempty"`
	Deaf                       bool        `json:"deaf,omitempty"`
	Mute                       bool        `json:"mute,omitempty"`
	IsPending                  bool        `json:"pending,omitempty"`
}

// User is a Discord User
type User struct {
	ID            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Nick          string    `json:"nick,omitempty"`
	Avatar        Hash      `json:"avatar,omitempty"`
	Bot           bool      `json:"bot,omitempty"`
	System        bool      `json:"system,omitempty"`
	MFAEnabled    bool      `json:"mfa_enabled,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         string    `json:"email,omitempty"`
	Flags         int       `json:"flags,omitempty"`
	Banner        string    `json:"banner,omitempty"`
	AccentColor   uint32    `json:"accent_color,omitempty"`
	PremiumType   int       `json:"premium_type,omitempty"`
	PublicFlags   int       `json:"public_flags,omitempty"`
}

// ChannelType is the type of a channel
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType int

const (
	CHANNEL_GUILD_TEXT ChannelType = iota
	CHANNEL_DM
	CHANNEL_GUILD_VOICE
	CHANNEL_GROUP_DM
	CHANNEL_GUILD_CATEGORY
	CHANNEL_GUILD_NEWS
	CHANNEL_GUILD_STORE
	CHANNEL_GUILD_NEWS_THREAD = iota + 3
	CHANNEL_GUILD_PUBLIC_THREAD
	CHANNEL_GUILD_PRIVATE_THREAD
	CHANNEL_GUILD_STAGE_VOICE
)
