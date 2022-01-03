package format

import (
	"fmt"
	"testing"
	"time"

	"github.com/matryer/is"
)

const testID = "1234567890"

func TestFormat(t *testing.T) {
	tt := []struct {
		Name     string
		Fmt      func(string) string
		Expected string
	}{
		{
			Name:     "User",
			Fmt:      User,
			Expected: fmt.Sprintf("<@%s>", testID),
		},
		{
			Name:     "UserNick",
			Fmt:      UserNick,
			Expected: fmt.Sprintf("<@!%s>", testID),
		},
		{
			Name:     "Channel",
			Fmt:      Channel,
			Expected: fmt.Sprintf("<#%s>", testID),
		},
		{
			Name:     "Role",
			Fmt:      Role,
			Expected: fmt.Sprintf("<@&%s>", testID),
		},
		{
			Name: "Emoji",
			Fmt: func(s string) string {
				return Emoji("test", s)
			},
			Expected: fmt.Sprintf("<:test:%s>", testID),
		},
		{
			Name: "AnimatedEmoji",
			Fmt: func(s string) string {
				return AnimatedEmoji("test", s)
			},
			Expected: fmt.Sprintf("<a:test:%s>", testID),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			assert := is.New(t)
			out := tc.Fmt(testID)
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
