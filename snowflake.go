package corde

import (
	"encoding/json"
	"strconv"
)

// Snowflake is a DiscordSnowflake ID
type Snowflake uint64

// corde.SnowflakeFromString returns aSnowflake from a string
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
