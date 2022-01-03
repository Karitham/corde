package format

import (
	"fmt"
	"testing"
	"time"
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
			out := tc.Fmt(testID)
			if out != tc.Expected {
				t.Logf("Incorrect Format\nExpected: %s\n     Got: %s\n", tc.Expected, out)
				t.Fail()
			}
		})
	}
}

func TestTimestamp(t *testing.T) {
	now := time.Now()
	expected := fmt.Sprintf("<t:%d:f>", now.Unix())
	got := Timestamp(now)
	if got != expected {
		t.Logf("Expected: %s\n     Got: %s\n", expected, got)
		t.Fail()
	}
}
