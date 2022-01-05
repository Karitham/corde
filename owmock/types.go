package owmock

import "github.com/Karitham/corde"

// id	snowflake	id of the interaction
// application_id	snowflake	id of the application this interaction is for
// type	interaction type	the type of interaction
// data?*	interaction data	the command data payload
// guild_id?	snowflake	the guild it was sent from
// channel_id?	snowflake	the channel it was sent from
// member?**	guild member object	guild member data for the invoking user, including permissions
// user?	user object	user object for the invoking user, if invoked in a DM
// token	string	a continuation token for responding to the interaction
// version	integer	read-only property, always 1
// message?	message object	for components, the message they were attached to
type Interaction struct {
	ID            corde.Snowflake `json:"id,omitempty"`
	ApplicationID corde.Snowflake `json:"application_id,omitempty"`
	Type          InteractionType `json:"type,omitempty"`
	Data          InteractionData `json:"data,omitempty"`
	GuildID       corde.Snowflake `json:"guild_id,omitempty"`
	ChannelID     corde.Snowflake `json:"channel_id,omitempty"`
	Member        Member          `json:"member,omitempty"`
	User          User            `json:"user,omitempty"`
	Token         string          `json:"token,omitempty"`
	Version       int             `json:"version,omitempty"`
	Message       Message         `json:"message,omitempty"`
}

// PING	1
// APPLICATION_COMMAND	2
// MESSAGE_COMPONENT	3
// APPLICATION_COMMAND_AUTOCOMPLETE	4
type InteractionType int

const (
	PING InteractionType = iota + 1
	APPLICATION_COMMAND
	MESSAGE_COMPONENT
	APPLICATION_COMMAND_AUTOCOMPLETE
)

// id	snowflake	the ID of the invoked command	Application Command
// name	string	the name of the invoked command	Application Command
// type	integer	the type of the invoked command	Application Command
// resolved?	resolved data	converted users + roles + channels	Application Command
// options?	array of application command interaction data option	the params + values from the user	Application Command
// custom_id?	string	the custom_id of the component	Component
// component_type?	integer	the type of the component	Component
// values?	array of select option values	the values the user selected	Component (Select)
// target_id?	snowflake	id the of user or message targetted by a user or message command	User Command, Message Command
type InteractionData struct {
	ID            corde.Snowflake        `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Type          ApplicationCommandType `json:"type,omitempty"`
	Resolved      Resolved               `json:"resolved,omitempty"`
	Options       []Option               `json:"options,omitempty"`
	CustomID      string                 `json:"custom_id,omitempty"`
	ComponentType int                    `json:"component_type,omitempty"`
	Values        []string               `json:"values,omitempty"`
	TargetID      corde.Snowflake        `json:"target_id,omitempty"`
}

// CHAT_INPUT	1	Slash commands; a text-based command that shows up when a user types
// USER	2	A UI-based command that shows up when you right click or tap on a user
// MESSAGE	3	A UI-based command that shows up when you right click or tap on a message
type ApplicationCommandType int

const (
	CHAT_INPUT ApplicationCommandType = iota + 1
	USER
	MESSAGE
)

// user?	user object	the user this guild member represents
// nick?	?string	this user's guild nickname
// avatar?	?string	the member's guild avatar hash
// roles	array of snowflakes	array of role object ids
// joined_at	ISO8601 timestamp	when the user joined the guild
// premium_since?	?ISO8601 timestamp	when the user started boosting the guild
// deaf	boolean	whether the user is deafened in voice channels
// mute	boolean	whether the user is muted in voice channels
// pending?	boolean	whether the user has not yet passed the guild's Membership Screening requirements
// permissions?	string	total permissions of the member in the channel, including overwrites, returned when in the interaction object
// communication_disabled_until?	?ISO8601 timestamp	when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
type Member struct {
	User          User              `json:"user,omitempty"`
	Nick          string            `json:"nick,omitempty"`
	Avatar        string            `json:"avatar,omitempty"`
	Roles         []corde.Snowflake `json:"roles,omitempty"`
	JoinedAt      string            `json:"joined_at,omitempty"`
	PremiumSince  corde.Timestamp   `json:"premium_since,omitempty"`
	Deaf          bool              `json:"deaf,omitempty"`
	Mute          bool              `json:"mute,omitempty"`
	Pending       bool              `json:"pending,omitempty"`
	Permissions   string            `json:"permissions,omitempty"`
	Communication corde.Timestamp   `json:"communication_disabled_until,omitempty"`
}

// id	snowflake	the user's id	identify
// username	string	the user's username, not unique across the platform	identify
// discriminator	string	the user's 4-digit discord-tag	identify
// avatar	?string	the user's avatar hash	identify
// bot?	boolean	whether the user belongs to an OAuth2 application	identify
// system?	boolean	whether the user is an Official Discord System user (part of the urgent message system)	identify
// mfa_enabled?	boolean	whether the user has two factor enabled on their account	identify
// banner?	?string	the user's banner hash	identify
// accent_color?	?integer	the user's banner color encoded as an integer representation of hexadecimal color code	identify
// locale?	string	the user's chosen language option	identify
// verified?	boolean	whether the email on this account has been verified	email
// email?	?string	the user's email	email
// flags?	integer	the flags on a user's account	identify
// premium_type?	integer	the type of Nitro subscription on a user's account	identify
// public_flags?	integer	the public flags on a user's account	identify
type User struct {
	ID            corde.Snowflake `json:"id,omitempty"`
	Username      string          `json:"username,omitempty"`
	Discriminator string          `json:"discriminator,omitempty"`
	Avatar        string          `json:"avatar,omitempty"`
	Bot           bool            `json:"bot,omitempty"`
	System        bool            `json:"system,omitempty"`
	MFAEnabled    bool            `json:"mfa_enabled,omitempty"`
	Banner        string          `json:"banner,omitempty"`
	AccentColor   int             `json:"accent_color,omitempty"`
	Locale        string          `json:"locale,omitempty"`
	Verified      bool            `json:"verified,omitempty"`
	Email         string          `json:"email,omitempty"`
	Flags         int             `json:"flags,omitempty"`
	PremiumType   int             `json:"premium_type,omitempty"`
	PublicFlags   int             `json:"public_flags,omitempty"`
}

// users?	Map of corde.Snowflakes to user objects	the ids and User objects
// members?*	Map of corde.Snowflakes to partial member objects	the ids and partial Member objects
// roles?	Map of corde.Snowflakes to role objects	the ids and Role objects
// channels?**	Map of corde.Snowflakes to partial channel objects	the ids and partial Channel objects
// messages?	Map of corde.Snowflakes to partial messages objects	the ids and partial Message objects
type Resolved struct {
	Users    map[corde.Snowflake]User    `json:"users,omitempty"`
	Members  map[corde.Snowflake]Member  `json:"members,omitempty"`
	Roles    map[corde.Snowflake]Role    `json:"roles,omitempty"`
	Channels map[corde.Snowflake]Channel `json:"channels,omitempty"`
	Messages map[corde.Snowflake]Message `json:"messages,omitempty"`
}

// id	snowflake	role id
// name	string	role name
// color	integer	integer representation of hexadecimal color code
// hoist	boolean	if this role is pinned in the user listing
// icon?	?string	role icon hash
// unicode_emoji?	?string	role unicode emoji
// position	integer	position of this role
// permissions	string	permission bit set
// managed	boolean	whether this role is managed by an integration
// mentionable	boolean	whether this role is mentionable
// tags?	role tags object	the tags this role has
type Role struct {
	ID           corde.Snowflake `json:"id,omitempty"`
	Name         string          `json:"name,omitempty"`
	Color        int             `json:"color,omitempty"`
	Hoist        bool            `json:"hoist,omitempty"`
	Icon         string          `json:"icon,omitempty"`
	UnicodeEmoji string          `json:"unicode_emoji,omitempty"`
	Position     int             `json:"position,omitempty"`
	Permissions  string          `json:"permissions,omitempty"`
	Managed      bool            `json:"managed,omitempty"`
	Mentionable  bool            `json:"mentionable,omitempty"`
	Tags         []RoleTag       `json:"tags,omitempty"`
}

// bot_id?	snowflake	the id of the bot this role belongs to
// integration_id?	snowflake	the id of the integration this role belongs to
// premium_subscriber?	null	whether this is the guild's premium subscriber role
type RoleTag struct {
	BotID             corde.Snowflake `json:"bot_id,omitempty"`
	IntegrationID     corde.Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber bool            `json:"premium_subscriber,omitempty"`
}

// id	snowflake	channel id
// name	string	channel name
// type	integer	channel type
// guild_id	snowflake	guild id
// position	integer	channel position
// permission_overwrites	array of permission overwrite objects
// topic?	?string	channel topic
// nsfw?	boolean	whether the channel is nsfw
// last_message_id?	?snowflake	last message id
// bitrate?	?integer	channel bitrate
// user_limit?	?integer	channel user limit
// rate_limit_per_user?	?integer	channel rate limit per user
// last_pin_timestamp?	?ISO8601 timestamp	last pin timestamp
// owner_id?	?snowflake	channel owner id
type Channel struct {
	ID                   corde.Snowflake `json:"id,omitempty"`
	Name                 string          `json:"name,omitempty"`
	Type                 int             `json:"type,omitempty"`
	GuildID              corde.Snowflake `json:"guild_id,omitempty"`
	Position             int             `json:"position,omitempty"`
	PermissionOverwrites []Overwrite     `json:"permission_overwrites,omitempty"`
	Topic                string          `json:"topic,omitempty"`
	NSFW                 bool            `json:"nsfw,omitempty"`
	LastMessageID        corde.Snowflake `json:"last_message_id,omitempty"`
	Bitrate              int             `json:"bitrate,omitempty"`
	UserLimit            int             `json:"user_limit,omitempty"`
	RateLimitPerUser     int             `json:"rate_limit_per_user,omitempty"`
	LastPinTimestamp     corde.Timestamp `json:"last_pin_timestamp,omitempty"`
	OwnerID              corde.Snowflake `json:"owner_id,omitempty"`
}

// id	snowflake	id of the message
// channel_id	snowflake	id of the channel the message was sent in
// guild_id?	snowflake	id of the guild the message was sent in
// author*	user object	the author of this message (not guaranteed to be a valid user, see below)
// member?**	partial guild member object	member properties for this message's author
// content	string	contents of the message
// timestamp	ISO8601 timestamp	when this message was sent
// edited_timestamp	?ISO8601 timestamp	when this message was edited (or null if never)
// tts	boolean	whether this was a TTS message
// mention_everyone	boolean	whether this message mentions everyone
// mentions***	array of user objects, with an additional partial member field	users specifically mentioned in the message
// mention_roles	array of role object ids	roles specifically mentioned in this message
// mention_channels?****	array of channel mention objects	channels specifically mentioned in this message
// attachments	array of attachment objects	any attached files
// embeds	array of embed objects	any embedded content
// reactions?	array of reaction objects	reactions to the message
// nonce?	integer or string	used for validating a message was sent
// pinned	boolean	whether this message is pinned
// webhook_id?	snowflake	if the message is generated by a webhook, this is the webhook's id
// type	integer	type of message
// activity?	message activity object	sent with Rich Presence-related chat embeds
// application?	partial application object	sent with Rich Presence-related chat embeds
// application_id?	snowflake	if the message is a response to an Interaction, this is the id of the interaction's application
// message_reference?	message reference object	data showing the source of a crosspost, channel follow add, pin, or reply message
// flags?	integer	message flags combined as a bitfield
// referenced_message?*****	?message object	the message associated with the message_reference
// interaction?	message interaction object	sent if the message is a response to an Interaction
// thread?	channel object	the thread that was started from this message, includes thread member object
// components?	Array of message components	sent if the message contains components like buttons, action rows, or other interactive components
// sticker_items?	array of message sticker item objects	sent if the message contains stickers
// stickers?	array of sticker objects	Deprecated the stickers sent with the message
type Message struct {
	ID                corde.Snowflake   `json:"id,omitempty"`
	ChannelID         corde.Snowflake   `json:"channel_id,omitempty"`
	GuildID           corde.Snowflake   `json:"guild_id,omitempty"`
	Author            User              `json:"user,omitempty"`
	Member            Member            `json:"member,omitempty"`
	Content           string            `json:"content,omitempty"`
	Timestamp         corde.Timestamp   `json:"timestamp,omitempty"`
	Edited            corde.Timestamp   `json:"edited_timestamp,omitempty"`
	TTS               bool              `json:"tts,omitempty"`
	Mention           bool              `json:"mention_everyone,omitempty"`
	Mentions          []User            `json:"mentions,omitempty"`
	MentionRoles      []corde.Snowflake `json:"mention_roles,omitempty"`
	MentionChannels   []Channel         `json:"mention_channels,omitempty"`
	Attachments       []Attachment      `json:"attachments,omitempty"`
	Embeds            []Embed           `json:"embeds,omitempty"`
	Reactions         []Reaction        `json:"reactions,omitempty"`
	Nonce             string            `json:"nonce,omitempty"`
	Pinned            bool              `json:"pinned,omitempty"`
	WebhookID         corde.Snowflake   `json:"webhook_id,omitempty"`
	Type              int               `json:"type,omitempty"`
	Activity          Activity          `json:"activity,omitempty"`
	Application       Application       `json:"application,omitempty"`
	ApplicationID     corde.Snowflake   `json:"application_id,omitempty"`
	MessageReference  MessageReference  `json:"message_reference,omitempty"`
	Flags             int               `json:"flags,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	Interaction       *Interaction      `json:"interaction,omitempty"`
	Thread            Channel           `json:"thread,omitempty"`
	Components        []Component       `json:"components,omitempty"`
	StickerItems      []StickerItem     `json:"sticker_items,omitempty"`
	Stickers          []Sticker         `json:"stickers,omitempty"`
}

// name	string	the name of the parameter
// type	integer	value of application command option type
// value?	string, integer, or double	the value of the option resulting from user input
// options?	array of application command interaction data option	present if this option is a group or subcommand
// focused?	boolean	true if this option is the currently focused option for autocomplete
type Option struct {
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Options []Option
	Focused bool `json:"focused,omitempty"`
}

// id	snowflake	attachment id
// filename	string	name of file attached
// description?	string	description for the file
// content_type?	string	the attachment's media type
// size	integer	size of file in bytes
// url	string	source url of file
// proxy_url	string	a proxied url of file
// height?	?integer	height of file (if image)
// width?	?integer	width of file (if image)
// ephemeral? *	boolean	whether this attachment is ephemeral
type Attachment struct {
	ID          corde.Snowflake `json:"id,omitempty"`
	Filename    string          `json:"filename,omitempty"`
	Description string          `json:"description,omitempty"`
	ContentType string          `json:"content_type,omitempty"`
	Size        int             `json:"size,omitempty"`
	URL         string          `json:"url,omitempty"`
	ProxyURL    string          `json:"proxy_url,omitempty"`
	Height      int             `json:"height,omitempty"`
	Width       int             `json:"width,omitempty"`
	Ephemeral   bool            `json:"ephemeral,omitempty"`
}

// title?	string	title of embed
// type?	string	type of embed (always "rich" for webhook embeds)
// description?	string	description of embed
// url?	string	url of embed
// timestamp?	ISO8601 timestamp	timestamp of embed content
// color?	integer	color code of the embed
// footer?	embed footer object	footer information
// image?	embed image object	image information
// thumbnail?	embed thumbnail object	thumbnail information
// video?	embed video object	video information
// provider?	embed provider object	provider information
// author?	embed author object	author information
// fields?	array of embed field objects	fields information
type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   corde.Timestamp `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      EmbedFooter     `json:"footer,omitempty"`
	Image       EmbedImage      `json:"image,omitempty"`
	Thumbnail   EmbedThumbnail  `json:"thumbnail,omitempty"`
	Video       EmbedVideo      `json:"video,omitempty"`
	Provider    EmbedProvider   `json:"provider,omitempty"`
	Author      EmbedAuthor     `json:"author,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
}

// text	string	footer text
// icon_url?	string	url of footer icon (only supports http(s) and attachments)
// proxy_icon_url?	string	a proxied url of footer icon
type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// name	string	name of the field
// value	string	value of the field
// inline?	boolean	whether or not this field should display inline
type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// url	string	source url of image (only supports http(s) and attachments)
// proxy_url?	string	a proxied url of the image
// height?	integer	height of image
// width?	integer	width of image
type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// name?	string	name of provider
// url?	string	url of provider
type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// name	string	name of author
// url?	string	url of author
// icon_url?	string	url of author icon (only supports http(s) and attachments)
// proxy_icon_url?	string	a proxied url of author icon
type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// url?	string	source url of video
// proxy_url?	string	a proxied url of the video
// height?	integer	height of video
// width?	integer	width of video
type EmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// url	string	source url of thumbnail (only supports http(s) and attachments)
// proxy_url?	string	a proxied url of the thumbnail
// height?	integer	height of thumbnail
// width?	integer	width of thumbnail
type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// count	integer	times this emoji has been used to react
// me	boolean	whether the current user reacted using this emoji
// emoji	partial emoji object	emoji information
type Reaction struct {
	Count int   `json:"count,omitempty"`
	Me    bool  `json:"me,omitempty"`
	Emoji Emoji `json:"emoji,omitempty"`
}

// id	?snowflake	emoji id
// name	?string (can be null only in reaction emoji objects)	emoji name
// roles?	array of role object ids	roles allowed to use this emoji
// user?	user object	user that created this emoji
// require_colons?	boolean	whether this emoji must be wrapped in colons
// managed?	boolean	whether this emoji is managed
// animated?	boolean	whether this emoji is animated
// available?	boolean	whether this emoji can be used, may be false due to loss of Server Boosts
type Emoji struct {
	ID            corde.Snowflake   `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Roles         []corde.Snowflake `json:"roles,omitempty"`
	User          User              `json:"user,omitempty"`
	RequireColons bool              `json:"require_colons,omitempty"`
	Managed       bool              `json:"managed,omitempty"`
	Animated      bool              `json:"animated,omitempty"`
	Available     bool              `json:"available,omitempty"`
}

// id	snowflake	role or user id
// type	int	either 0 (role) or 1 (member)
// allow	string	permission bit set
// deny	string	permission bit set
type Overwrite struct {
	ID    corde.Snowflake `json:"id,omitempty"`
	Type  int             `json:"type,omitempty"`
	Allow string          `json:"allow,omitempty"`
	Deny  string          `json:"deny,omitempty"`
}

// type	integer	type of message activity
// party_id?	string	party_id from a Rich Presence event
type Activity struct {
	Type    int    `json:"type,omitempty"`
	PartyID string `json:"party_id,omitempty"`
}

// id	snowflake	the id of the app
// name	string	the name of the app
// icon	?string	the icon hash of the app
// description	string	the description of the app
// rpc_origins?	array of strings	an array of rpc origin urls, if rpc is enabled
// bot_public	boolean	when false only app owner can join the app's bot to guilds
// bot_require_code_grant	boolean	when true the app's bot will only join upon completion of the full oauth2 code grant flow
// terms_of_service_url?	string	the url of the app's terms of service
// privacy_policy_url?	string	the url of the app's privacy policy
// owner?	partial user object	partial user object containing info on the owner of the application
// summary	string	if this application is a game sold on Discord, this field will be the summary field for the store page of its primary sku
// verify_key	string	the hex encoded key for verification in interactions and the GameSDK's GetTicket
// team	?team object	if the application belongs to a team, this will be a list of the members of that team
// guild_id?	snowflake	if this application is a game sold on Discord, this field will be the guild to which it has been linked
// primary_sku_id?	snowflake	if this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
// slug?	string	if this application is a game sold on Discord, this field will be the URL slug that links to the store page
// cover_image?	string	the application's default rich presence invite cover image hash
// flags?	integer	the application's public flags
type Application struct {
	ID                  corde.Snowflake `json:"id,omitempty"`
	Name                string          `json:"name,omitempty"`
	Icon                string          `json:"icon,omitempty"`
	Description         string          `json:"description,omitempty"`
	RPCOrigins          []string        `json:"rpc_origins,omitempty"`
	BotPublic           bool            `json:"bot_public,omitempty"`
	BotRequireCodeGrant bool            `json:"bot_require_code_grant,omitempty"`
	TermsOfServiceURL   string          `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string          `json:"privacy_policy_url,omitempty"`
	Owner               User            `json:"owner,omitempty"`
	Summary             string          `json:"summary,omitempty"`
	VerifyKey           string          `json:"verify_key,omitempty"`
	Team                Team            `json:"team,omitempty"`
	GuildID             corde.Snowflake `json:"guild_id,omitempty"`
	PrimarySKUID        corde.Snowflake `json:"primary_sku_id,omitempty"`
	Slug                string          `json:"slug,omitempty"`
	CoverImage          string          `json:"cover_image,omitempty"`
	Flags               int             `json:"flags,omitempty"`
}

// icon	?string	a hash of the image of the team's icon
// id	snowflake	the unique id of the team
// members	array of team member objects	the members of the team
// name	string	the name of the team
// owner_user_id	snowflake	the user id of the current team owner
type Team struct {
	Icon        string          `json:"icon,omitempty"`
	ID          corde.Snowflake `json:"id,omitempty"`
	Members     []TeamMember    `json:"members,omitempty"`
	Name        string          `json:"name,omitempty"`
	OwnerUserID corde.Snowflake `json:"owner_user_id,omitempty"`
}

// membership_state	integer	the user's membership state on the team
// permissions	array of strings	will always be ["*"]
// team_id	snowflake	the id of the parent team of which they are a member
// user	partial user object	the avatar, discriminator, id, and username of the user
type TeamMember struct {
	MembershipState int             `json:"membership_state,omitempty"`
	Permissions     []string        `json:"permissions,omitempty"`
	TeamID          corde.Snowflake `json:"team_id,omitempty"`
	User            User            `json:"user,omitempty"`
}

// message_id?	snowflake	id of the originating message
// channel_id? *	snowflake	id of the originating message's channel
// guild_id?	snowflake	id of the originating message's guild
// fail_if_not_exists?	boolean	when sending, whether to error if the referenced message doesn't exist instead of sending as a normal (non-reply) message, default true
type MessageReference struct {
	MessageID       corde.Snowflake `json:"message_id,omitempty"`
	ChannelID       corde.Snowflake `json:"channel_id,omitempty"`
	GuildID         corde.Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool            `json:"fail_if_not_exists,omitempty"`
}

// type	integer	component type	all types
// custom_id?	string	a developer-defined identifier for the component, max 100 characters	Buttons, Select Menus
// disabled?	boolean	whether the component is disabled, default false	Buttons, Select Menus
// style?	integer	one of button styles	Buttons
// label?	string	text that appears on the button, max 80 characters	Buttons
// emoji?	partial emoji	name, id, and animated	Buttons
// url?	string	a url for link-style buttons	Buttons
// options?	array of select options	the choices in the select, max 25	Select Menus
// placeholder?	string	custom placeholder text if nothing is selected, max 100 characters	Select Menus
// min_values?	integer	the minimum number of items that must be chosen; default 1, min 0, max 25	Select Menus
// max_values?	integer	the maximum number of items that can be chosen; default 1, max 25	Select Menus
// components?	array of components	a list of child components	Action Rows
type Component struct {
	Type        int         `json:"type,omitempty"`
	CustomID    string      `json:"custom_id,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Style       int         `json:"style,omitempty"`
	Label       string      `json:"label,omitempty"`
	Emoji       Emoji       `json:"emoji,omitempty"`
	URL         string      `json:"url,omitempty"`
	Options     []Option    `json:"options,omitempty"`
	Placeholder string      `json:"placeholder,omitempty"`
	MinValues   int         `json:"min_values,omitempty"`
	MaxValues   int         `json:"max_values,omitempty"`
	Components  []Component `json:"components,omitempty"`
}

// id	snowflake	id of the sticker
// name	string	name of the sticker
// format_type	integer	type of sticker format
type StickerItem struct {
	ID       corde.Snowflake `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	FormatID int             `json:"format_type,omitempty"`
}

// id	snowflake	id of the sticker
// pack_id?	snowflake	for standard stickers, id of the pack the sticker is from
// name	string	name of the sticker
// description	?string	description of the sticker
// tags*	string	autocomplete/suggestion tags for the sticker (max 200 characters)
// asset	string	Deprecated previously the sticker asset hash, now an empty string
// type	integer	type of sticker
// format_type	integer	type of sticker format
// available?	boolean	whether this guild sticker can be used, may be false due to loss of Server Boosts
// guild_id?	snowflake	id of the guild that owns this sticker
// user?	user object	the user that uploaded the guild sticker
// sort_value?	integer	the standard sticker's sort order within its pack
type Sticker struct {
	ID          corde.Snowflake `json:"id,omitempty"`
	PackID      corde.Snowflake `json:"pack_id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	Asset       string          `json:"asset,omitempty"`
	Type        int             `json:"type,omitempty"`
	FormatType  int             `json:"format_type,omitempty"`
	Available   bool            `json:"available,omitempty"`
	GuildID     corde.Snowflake `json:"guild_id,omitempty"`
	User        User            `json:"user,omitempty"`
	SortValue   int             `json:"sort_value,omitempty"`
}

// type	interaction callback type	the type of response
// data?	interaction callback data	an optional response message
type InteractionResponse struct {
	Type int                     `json:"type,omitempty"`
	Data InteractionResponseData `json:"data,omitempty"`
}

// tts?	boolean	is the response TTS
// content?	string	message content
// embeds?	array of embeds	supports up to 10 embeds
// allowed_mentions?	allowed mentions	allowed mentions object
// flags?	integer	interaction callback data flags
// components?	array of components	message components
// attachments? *	array of partial attachment objects	attachment objects with filename and description
type InteractionResponseData struct {
	TTS             bool            `json:"tts,omitempty"`
	Content         string          `json:"content,omitempty"`
	Embeds          []Embed         `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           int             `json:"flags,omitempty"`
	Components      []Component     `json:"components,omitempty"`
	Attachments     []Attachment    `json:"attachments,omitempty"`
}

type AllowedMentions string

// Role Mentions	"roles"	Controls role mentions
// User Mentions	"users"	Controls user mentions
// Everyone Mentions	"everyone"	Controls @everyone and @here mentions
const (
	RoleMentions     AllowedMentions = "roles"
	UserMentions     AllowedMentions = "users"
	EveryoneMentions AllowedMentions = "everyone"
)
