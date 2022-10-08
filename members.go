package corde

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
