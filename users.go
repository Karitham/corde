package corde

import (
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/internal/rest"
	"github.com/Karitham/corde/snowflake"
)

// Me returns the current user
func (m *Mux) Me() (components.User, error) {
	var user components.User
	_, err := rest.DoJSON(m.Client, rest.Req("/users/@me").Get(m.authorize), &user)
	return user, err
}

// GetUser returns a user by id
func (m *Mux) GetUser(id snowflake.Snowflake) (components.User, error) {
	var user components.User
	_, err := rest.DoJSON(m.Client, rest.Req("/users/", id).Get(m.authorize), &user)
	return user, err
}
