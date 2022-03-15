package corde

import (
	"fmt"
)

// OptionsUser returns the resolved User for an Option
func (i resolvedInteractionWithOptions) OptionsUser(k string) (User, error) {
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
func (i resolvedInteractionWithOptions) OptionsMember(k string) (Member, error) {
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
func (i resolvedInteractionWithOptions) OptionsRole(k string) (Role, error) {
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
func (i resolvedInteractionWithOptions) OptionsMessage(k string) (Message, error) {
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
