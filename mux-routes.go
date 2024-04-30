package corde

import "context"

// ButtonComponent mounts a button route on the mux
func (m *Mux) ButtonComponent(route string, handler func(context.Context, ResponseWriter, *Interaction[ButtonInteractionData])) {
	m.Mount(ButtonInteraction, route, handler)
}

// Autocomplete mounts an autocomplete route on the mux
func (m *Mux) Autocomplete(route string, handler func(context.Context, ResponseWriter, *Interaction[AutocompleteInteractionData])) {
	m.Mount(AutocompleteInteraction, route, handler)
}

// SlashCommand mounts a slash command route on the mux
func (m *Mux) SlashCommand(route string, handler func(context.Context, ResponseWriter, *Interaction[SlashCommandInteractionData])) {
	m.Mount(SlashCommandInteraction, route, handler)
}

// UserCommand mounts a user command on the mux
func (m *Mux) UserCommand(route string, handler func(context.Context, ResponseWriter, *Interaction[UserCommandInteractionData])) {
	m.Mount(UserCommandInteraction, route, handler)
}

// MessageCommand mounts a message command on the mux
func (m *Mux) MessageCommand(route string, handler func(context.Context, ResponseWriter, *Interaction[MessageCommandInteractionData])) {
	m.Mount(MessageCommandInteraction, route, handler)
}

// Modal mounts a modal interaction response on the mux
func (m *Mux) Modal(route string, handler func(context.Context, ResponseWriter, *Interaction[ModalInteractionData])) {
	m.Mount(ModalInteraction, route, handler)
}
