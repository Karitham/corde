package corde

import (
	"github.com/Karitham/corde/internal/rest"
)

// Me returns the current user
func (m *Mux) Me() (User, error) {
	var user User
	_, err := rest.DoJson(m.Client, rest.Req("/users/@me").Get(m.authorize), &user)
	return user, err
}

// GetUser returns a user by id
func (m *Mux) GetUser(id Snowflake) (User, error) {
	var user User
	_, err := rest.DoJson(m.Client, rest.Req("/users/", id).Get(m.authorize), &user)
	return user, err
}
