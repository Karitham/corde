package corde

import (
	"github.com/matryer/is"
	"testing"
)

func TestRoute(t *testing.T) {
	assert := is.New(t)

	m := NewMux("", Snowflake(0), "")

	m.Route("foo", func(m *Mux) {
		m.Route("bar", func(m *Mux) {
			m.Command("baz", nil)
		})
	})
	m.Route("/foo", func(m *Mux) {
		m.Route("/bar", func(m *Mux) {
			m.Command("/baz", nil)
		})
	})
	m.Route("foo/", func(m *Mux) {
		m.Route("bar/", func(m *Mux) {
			m.Command("baz/", nil)
		})
	})
	m.Route("/foo/", func(m *Mux) {
		m.Route("/bar/", func(m *Mux) {
			m.Command("/baz/", nil)
		})
	})
	m.Command("foo/bar/baz", nil)

	var commands []string
	for cmd := range m.routes.command.ToMap() {
		commands = append(commands, cmd)
	}

	assert.Equal(commands, []string{"foo/bar/baz"}) // There should only be a single command on the router
}
