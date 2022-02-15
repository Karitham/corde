package format

import (
	"fmt"
	"testing"
	"time"

	"github.com/Karitham/corde/snowflake"
	"github.com/matryer/is"
)

const testIDString = "1234567890"

var testIDSnowflake = snowflake.SnowflakeFromString(testIDString)

type testcase[T ID] struct {
	Name     string
	Fmt      func(T) string
	Expected string
}

func TestFormatString(t *testing.T) {
	tt := []testcase[string]{
		{
			Name:     "User",
			Fmt:      User[string],
			Expected: fmt.Sprintf("<@%s>", testIDString),
		},
		{
			Name:     "UserNick",
			Fmt:      UserNick[string],
			Expected: fmt.Sprintf("<@!%s>", testIDString),
		},
		{
			Name:     "Channel",
			Fmt:      Channel[string],
			Expected: fmt.Sprintf("<#%s>", testIDString),
		},
		{
			Name:     "Role",
			Fmt:      Role[string],
			Expected: fmt.Sprintf("<@&%s>", testIDString),
		},
		{
			Name: "Emoji",
			Fmt: func(s string) string {
				return Emoji("test", s)
			},
			Expected: fmt.Sprintf("<:test:%s>", testIDString),
		},
		{
			Name: "AnimatedEmoji",
			Fmt: func(s string) string {
				return AnimatedEmoji("test", s)
			},
			Expected: fmt.Sprintf("<a:test:%s>", testIDString),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert := is.New(t)
			out := tc.Fmt(testIDString)
			assert.Equal(out, tc.Expected) // Formatted should match expected
		})
	}
}

func TestFormatsSnowflake(t *testing.T) {
	tt := []testcase[snowflake.Snowflake]{
		{
			Name:     "User",
			Fmt:      User[snowflake.Snowflake],
			Expected: fmt.Sprintf("<@%s>", testIDString),
		},
		{
			Name:     "UserNick",
			Fmt:      UserNick[snowflake.Snowflake],
			Expected: fmt.Sprintf("<@!%s>", testIDString),
		},
		{
			Name:     "Channel",
			Fmt:      Channel[snowflake.Snowflake],
			Expected: fmt.Sprintf("<#%s>", testIDString),
		},
		{
			Name:     "Role",
			Fmt:      Role[snowflake.Snowflake],
			Expected: fmt.Sprintf("<@&%s>", testIDString),
		},
		{
			Name: "Emoji",
			Fmt: func(s snowflake.Snowflake) string {
				return Emoji("test", s)
			},
			Expected: fmt.Sprintf("<:test:%s>", testIDString),
		},
		{
			Name: "AnimatedEmoji",
			Fmt: func(s snowflake.Snowflake) string {
				return AnimatedEmoji("test", s)
			},
			Expected: fmt.Sprintf("<a:test:%s>", testIDString),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert := is.New(t)
			out := tc.Fmt(testIDSnowflake)
			assert.Equal(out, tc.Expected) // Formatted should match expected
		})
	}
}

func TestTimestamp(t *testing.T) {
	assert := is.New(t)

	now := time.Now()
	expected := fmt.Sprintf("<t:%d:f>", now.Unix())
	got := Timestamp(now)
	assert.Equal(got, expected) // Formatted should match expected
}
