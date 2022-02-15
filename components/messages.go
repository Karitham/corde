package components

import "github.com/Karitham/corde/snowflake"

// Message is a Discord Message
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                snowflake.Snowflake   `json:"id"`
	ChannelID         snowflake.Snowflake   `json:"channel_id"`
	GuildID           snowflake.Snowflake   `json:"guild_id,omitempty"`
	Author            User                  `json:"user,omitempty"`
	Member            Member                `json:"member,omitempty"`
	Content           string                `json:"content"`
	Timestamp         Timestamp             `json:"timestamp"`
	Edited            Timestamp             `json:"edited_timestamp,omitempty"`
	TTS               bool                  `json:"tts"`
	Mention           bool                  `json:"mention_everyone"`
	Mentions          []User                `json:"mentions,omitempty"`
	MentionRoles      []snowflake.Snowflake `json:"mention_roles,omitempty"`
	MentionChannels   []Channel             `json:"mention_channels,omitempty"`
	Attachments       []Attachment          `json:"attachments,omitempty"`
	Embeds            []Embed               `json:"embeds,omitempty"`
	Reactions         []Reaction            `json:"reactions,omitempty"`
	Nonce             string                `json:"nonce,omitempty"`
	Pinned            bool                  `json:"pinned,omitempty"`
	WebhookID         snowflake.Snowflake   `json:"webhook_id,omitempty"`
	Type              MessageType           `json:"type"`
	Activity          Activity              `json:"activity,omitempty"`
	Application       Application           `json:"application,omitempty"`
	ApplicationID     snowflake.Snowflake   `json:"application_id,omitempty"`
	MessageReference  MessageReference      `json:"message_reference,omitempty"`
	Flags             MessageFlag           `json:"flags,omitempty"`
	ReferencedMessage *Message              `json:"referenced_message,omitempty"`
	Interaction       *Interaction          `json:"interaction,omitempty"`
	Thread            Channel               `json:"thread,omitempty"`
	Components        []Component           `json:"components,omitempty"`
	StickerItems      []StickerItem         `json:"sticker_items,omitempty"`
	Stickers          []Sticker             `json:"stickers,omitempty"`
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
