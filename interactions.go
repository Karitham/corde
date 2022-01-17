package corde

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// InteractionType is the type of interaction
type InteractionType int

const (
	INTERACTION_TYPE_PING InteractionType = iota + 1
	INTERACTION_TYPE_APPLICATION_COMMAND
	INTERACTION_TYPE_MESSAGE_COMPONENT
	INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE
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
// https://discord.com/developers/docs/interactions/receiving-and-responding#interactions
type Interaction struct {
	ID            Snowflake       `json:"id"`
	ApplicationID Snowflake       `json:"application_id"`
	Type          InteractionType `json:"type"`
	Data          InteractionData `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id,omitempty"`
	ChannelID     Snowflake       `json:"channel_id,omitempty"`
	Member        Member          `json:"member,omitempty"`
	User          User            `json:"user,omitempty"`
	Token         string          `json:"token"`
	Version       int             `json:"version"`
	Message       Message         `json:"message,omitempty"`
	Locale        string          `json:"locale,omitempty"`
	GuildLocale   string          `json:"guild_locale,omitempty"`
}

// Message is a Discord Message
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                Snowflake        `json:"id"`
	ChannelID         Snowflake        `json:"channel_id"`
	GuildID           Snowflake        `json:"guild_id,omitempty"`
	Author            User             `json:"user,omitempty"`
	Member            Member           `json:"member,omitempty"`
	Content           string           `json:"content"`
	Timestamp         Timestamp        `json:"timestamp"`
	Edited            Timestamp        `json:"edited_timestamp,omitempty"`
	TTS               bool             `json:"tts"`
	Mention           bool             `json:"mention_everyone"`
	Mentions          []User           `json:"mentions,omitempty"`
	MentionRoles      []Snowflake      `json:"mention_roles,omitempty"`
	MentionChannels   []Channel        `json:"mention_channels,omitempty"`
	Attachments       []Attachment     `json:"attachments,omitempty"`
	Embeds            []Embed          `json:"embeds,omitempty"`
	Reactions         []Reaction       `json:"reactions,omitempty"`
	Nonce             string           `json:"nonce,omitempty"`
	Pinned            bool             `json:"pinned,omitempty"`
	WebhookID         Snowflake        `json:"webhook_id,omitempty"`
	Type              MessageType      `json:"type"`
	Activity          Activity         `json:"activity,omitempty"`
	Application       Application      `json:"application,omitempty"`
	ApplicationID     Snowflake        `json:"application_id,omitempty"`
	MessageReference  MessageReference `json:"message_reference,omitempty"`
	Flags             MessageFlag      `json:"flags,omitempty"`
	ReferencedMessage *Message         `json:"referenced_message,omitempty"`
	Interaction       *Interaction     `json:"interaction,omitempty"`
	Thread            Channel          `json:"thread,omitempty"`
	Components        []Component      `json:"components,omitempty"`
	StickerItems      []StickerItem    `json:"sticker_items,omitempty"`
	Stickers          []Sticker        `json:"stickers,omitempty"`
}

// MessageType
// https://discord.com/developers/docs/resources/channel#message-object-message-types
type MessageType int

// We could have used iota, but this scales better
// https://regex101.com/r/a3XWNs/1
const (
	MESSAGE_TYPE_DEFAULT                                      MessageType = 0
	MESSAGE_TYPE_RECIPIENT_ADD                                MessageType = 1
	MESSAGE_TYPE_RECIPIENT_REMOVE                             MessageType = 2
	MESSAGE_TYPE_CALL                                         MessageType = 3
	MESSAGE_TYPE_CHANNEL_NAME_CHANGE                          MessageType = 4
	MESSAGE_TYPE_CHANNEL_ICON_CHANGE                          MessageType = 5
	MESSAGE_TYPE_CHANNEL_PINNED_MESSAGE                       MessageType = 6
	MESSAGE_TYPE_GUILD_MEMBER_JOIN                            MessageType = 7
	MESSAGE_TYPE_USER_PREMIUM_GUILD_SUBSCRIPTION              MessageType = 8
	MESSAGE_TYPE_USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_1       MessageType = 9
	MESSAGE_TYPE_USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_2       MessageType = 10
	MESSAGE_TYPE_USER_PREMIUM_GUILD_SUBSCRIPTION_TIER_3       MessageType = 11
	MESSAGE_TYPE_CHANNEL_FOLLOW_ADD                           MessageType = 12
	MESSAGE_TYPE_GUILD_DISCOVERY_DISQUALIFIED                 MessageType = 14
	MESSAGE_TYPE_GUILD_DISCOVERY_REQUALIFIED                  MessageType = 15
	MESSAGE_TYPE_GUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING MessageType = 16
	MESSAGE_TYPE_GUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING   MessageType = 17
	MESSAGE_TYPE_THREAD_CREATED                               MessageType = 18
	MESSAGE_TYPE_REPLY                                        MessageType = 19
	MESSAGE_TYPE_CHAT_INPUT_COMMAND                           MessageType = 20
	MESSAGE_TYPE_THREAD_STARTER_MESSAGE                       MessageType = 21
	MESSAGE_TYPE_GUILD_INVITE_REMINDER                        MessageType = 22
	MESSAGE_TYPE_CONTEXT_MENU_COMMAND                         MessageType = 23
)

// MessageFlag
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
type MessageFlag int

// https://regex101.com/r/1SsWbP/1
const (
	// MESSAGE_FLAG_CROSSPOSTED means this message has been published to subscribed channels (via Channel Following)
	MESSAGE_FLAG_CROSSPOSTED = 1 << 0
	// MESSAGE_FLAG_IS_CROSSPOST means this message originated from a message in another channel (via Channel Following)
	MESSAGE_FLAG_IS_CROSSPOST = 1 << 1
	// MESSAGE_FLAG_SUPPRESS_EMBEDS means do not include any embeds when serializing this message
	MESSAGE_FLAG_SUPPRESS_EMBEDS = 1 << 2
	// MESSAGE_FLAG_SOURCE_MESSAGE_DELETED means the source message for this crosspost has been deleted (via Channel Following)
	MESSAGE_FLAG_SOURCE_MESSAGE_DELETED = 1 << 3
	// MESSAGE_FLAG_URGENT means this message came from the urgent message system
	MESSAGE_FLAG_URGENT = 1 << 4
	// MESSAGE_FLAG_HAS_THREAD means this message has an associated thread, with the same id as the message
	MESSAGE_FLAG_HAS_THREAD = 1 << 5
	// MESSAGE_FLAG_EPHEMERAL means this message is only visible to the user who invoked the Interaction
	MESSAGE_FLAG_EPHEMERAL = 1 << 6
	// MESSAGE_FLAG_LOADING means this message is an Interaction Response and the bot is "thinking"
	MESSAGE_FLAG_LOADING = 1 << 7
)

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
// ResolvedData is a map (which is unordered), so First
// should only be used when ResolvedData has a single element.
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

type subRoute struct {
	Name    string     `json:"name"`
	Value   JsonRaw    `json:"value"`
	Type    OptionType `json:"type"`
	Options []subRoute `json:"options"`
}

// UnmarshalJSON implements json.Unmarshaler
func (o *OptionsInteractions) UnmarshalJSON(b []byte) error {
	var opts []subRoute
	if err := json.Unmarshal(b, &opts); err != nil {
		return err
	}

	// max is 3 deep, as per discord's docs
	m := make(map[string]JsonRaw)
	for _, opt := range opts {
		// enables us to route easily
		switch opt.Type {
		case OPTION_SUB_COMMAND_GROUP:
			opt.Value = []byte(opt.Name)
			opt.Name = "$group"
		case OPTION_SUB_COMMAND:
			opt.Value = []byte(opt.Name)
			opt.Name = "$command"
		}

		m[opt.Name] = opt.Value
		for _, opt2 := range opt.Options {
			// enables us to route easily
			if opt2.Type == OPTION_SUB_COMMAND {
				opt2.Value = []byte(opt2.Name)
				opt2.Name = "$command"
			}

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
	RoleIDs                    []Snowflake `json:"roles,omitempty"`
	Avatar                     Hash        `json:"avatar,omitempty"`
	Joined                     Timestamp   `json:"joined_at"`
	BoostedSince               Timestamp   `json:"premium_since,omitempty"`
	CommunicationDisabledUntil Timestamp   `json:"communication_disabled_until,omitempty"`
	Deaf                       bool        `json:"deaf,omitempty"`
	Mute                       bool        `json:"mute,omitempty"`
	IsPending                  bool        `json:"pending,omitempty"`
	Permissions                string      `json:"permissions,omitempty"`
}

// User is a Discord User
type User struct {
	ID            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        Hash      `json:"avatar,omitempty"`
	Bot           bool      `json:"bot,omitempty"`
	System        bool      `json:"system,omitempty"`
	MFAEnabled    bool      `json:"mfa_enabled,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         string    `json:"email,omitempty"`
	Flags         int       `json:"flags,omitempty"`
	Locale        string    `json:"locale,omitempty"`
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

// Sticker
// https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
	ID          Snowflake     `json:"id"`
	Name        string        `json:"name"`
	Tags        string        `json:"tags"`
	Asset       string        `json:"asset"`
	Type        int           `json:"type"`
	FormatType  StickerFormat `json:"format_type"`
	Description string        `json:"description,omitempty"`
	PackID      Snowflake     `json:"pack_id,omitempty"`
	Available   bool          `json:"available,omitempty"`
	GuildID     Snowflake     `json:"guild_id,omitempty"`
	User        User          `json:"user,omitempty"`
	SortValue   int           `json:"sort_value,omitempty"`
}

// StickerItem
type StickerItem struct {
	ID         Snowflake     `json:"id"`
	Name       string        `json:"name"`
	FormatType StickerFormat `json:"format_type"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
type StickerFormat int

const (
	STICKER_FORMAT_PNG StickerFormat = iota + 1
	STICKER_FORMAT_APNG
	STICKER_FORMAT_LOTTIE
)

// Channel
// https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	ID                   Snowflake   `json:"id"`
	Name                 string      `json:"name"`
	Type                 ChannelType `json:"type"`
	GuildID              Snowflake   `json:"guild_id"`
	Position             int         `json:"position"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites"`
	Topic                string      `json:"topic,omitempty"`
	NSFW                 bool        `json:"nsfw,omitempty"`
	LastMessageID        Snowflake   `json:"last_message_id,omitempty"`
	Bitrate              int         `json:"bitrate,omitempty"`
	UserLimit            int         `json:"user_limit,omitempty"`
	RateLimitPerUser     int         `json:"rate_limit_per_user,omitempty"`
	LastPinTimestamp     Timestamp   `json:"last_pin_timestamp,omitempty"`
	OwnerID              Snowflake   `json:"owner_id,omitempty"`
}

// Overwrite
// https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	ID Snowflake `json:"id"`
	// Type: 0 = @role, 1 = @user
	Type int `json:"type"`
	// Permission bit set
	Allow string `json:"allow"`
	// Permission bit set
	Deny string `json:"deny"`
}

// Reaction
// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

// Activity
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type Activity struct {
	Type    int    `json:"type,omitempty"`
	PartyID string `json:"party_id,omitempty"`
}

// MessageActivity
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
type MessageActivity int

// https://regex101.com/r/Sj1ZLk/1
const (
	MESSAGE_ACTIVITY_JOIN         = 1
	MESSAGE_ACTIVITY_SPECTATE     = 2
	MESSAGE_ACTIVITY_LISTEN       = 3
	MESSAGE_ACTIVITY_JOIN_REQUEST = 5
)

// Application
// https://discord.com/developers/docs/resources/application#application-object-application-structure
type Application struct {
	ID                  Snowflake `json:"id"`
	Name                string    `json:"name"`
	Icon                string    `json:"icon"`
	Description         string    `json:"description"`
	RPCOrigins          []string  `json:"rpc_origins,omitempty"`
	BotPublic           bool      `json:"bot_public"`
	BotRequireCodeGrant bool      `json:"bot_require_code_grant"`
	TermsOfServiceURL   string    `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string    `json:"privacy_policy_url,omitempty"`
	Owner               User      `json:"owner,omitempty"`
	Summary             string    `json:"summary"`
	VerifyKey           string    `json:"verify_key"`
	Team                Team      `json:"team,omitempty"`
	GuildID             Snowflake `json:"guild_id,omitempty"`
	PrimarySKUID        Snowflake `json:"primary_sku_id,omitempty"`
	Slug                string    `json:"slug,omitempty"`
	CoverImage          string    `json:"cover_image,omitempty"`
	Flags               int       `json:"flags,omitempty"`
}

// Team
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        string       `json:"icon,omitempty"`
	ID          Snowflake    `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID Snowflake    `json:"owner_user_id"`
}

// TeamMember
// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState int `json:"membership_state"`
	// 	will always be ["*"]
	Permissions []string  `json:"permissions"`
	TeamID      Snowflake `json:"team_id"`
	User        User      `json:"user"`
}
