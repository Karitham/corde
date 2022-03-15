package corde

import (
	"encoding/json"
	"errors"
)

// JsonRaw is a raw json value
type JsonRaw []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m JsonRaw) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *JsonRaw) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

// UnmarshalTo unmarshals the JSON-encoded data and stores the result in the value pointed to by v.
func (m JsonRaw) UnmarshalTo(v any) error {
	return json.Unmarshal(m, v)
}

// String returns the raw json as a string
func (m JsonRaw) String() string {
	return string(m)
}
