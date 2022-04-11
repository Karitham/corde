// Package format https://discord.com/developers/docs/reference#message-formatting
package format

import (
	"fmt"
	"time"

	"github.com/Karitham/corde"
)

// ID is a Discord ID, either in string form or Snowflake form
type ID interface {
	~string | corde.Snowflake
}

// User returns a user mention
func User[T ID](id T) string {
	return fmt.Sprintf("<@%s>", id)
}

// UserNick returns a user (nickname) mention
func UserNick[T ID](id T) string {
	return fmt.Sprintf("<@!%s>", id)
}

// Channel returns a channel mention
func Channel[T ID](id T) string {
	return fmt.Sprintf("<#%s>", id)
}

// Role returns a role mention
func Role[T ID](id T) string {
	return fmt.Sprintf("<@&%s>", id)
}

// Emoji returns a custom emoji
func Emoji[T ID](name string, id T) string {
	return fmt.Sprintf("<:%s:%s>", name, id)
}

// AnimatedEmoji returns a custom animated emoji
func AnimatedEmoji[T ID](name string, id T) string {
	return fmt.Sprintf("<a:%s:%s>", name, id)
}

// Timestamp returns a timestamp
func Timestamp(ts time.Time) string {
	return TimestampStyled(ts, TimestampShortDateTime)
}

// TimestampStyle https://discord.com/developers/docs/reference#message-formatting-timestamp-styles
type TimestampStyle string

const (
	TimestampShortTime     TimestampStyle = "t"
	TimestampLongTime      TimestampStyle = "T"
	TimestampShortDate     TimestampStyle = "d"
	TimestampLongDate      TimestampStyle = "D"
	TimestampShortDateTime TimestampStyle = "f"
	TimestampLongDateTime  TimestampStyle = "F"
	TimestampRelative      TimestampStyle = "R"
)

// TimestampStyled returns a styled timestamp
func TimestampStyled(ts time.Time, style TimestampStyle) string {
	return fmt.Sprintf("<t:%d:%s>", ts.Unix(), style)
}
