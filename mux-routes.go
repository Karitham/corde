package corde

import (
	"github.com/Karitham/corde/components"
)

// ButtonComponent mounts a button route on the mux
func (m *Mux) ButtonComponent(route string, handler func(ResponseWriter, *Request[components.ButtonInteractionData])) {
	m.Mount(components.ButtonInteraction, route, handler)
}

// Autocomplete mounts an autocomplete route on the mux
func (m *Mux) Autocomplete(route string, handler func(ResponseWriter, *Request[components.AutocompleteInteractionData])) {
	m.Mount(components.AutocompleteInteraction, route, handler)
}

// SlashCommand mounts a slash command route on the mux
func (m *Mux) SlashCommand(route string, handler func(ResponseWriter, *Request[components.SlashCommandInteractionData])) {
	m.Mount(components.SlashCommandInteraction, route, handler)
}

// UserCommand mounts a user command on the mux
func (m *Mux) UserCommand(route string, handler func(ResponseWriter, *Request[components.UserCommandInteractionData])) {
	m.Mount(components.UserCommandInteraction, route, handler)
}

// MessageCommand mounts a message command on the mux
func (m *Mux) MessageCommand(route string, handler func(ResponseWriter, *Request[components.MessageCommandInteractionData])) {
	m.Mount(components.MessageCommandInteraction, route, handler)
}
