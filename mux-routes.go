package corde

// ButtonComponent mounts a button route on the mux
func (m *Mux) ButtonComponent(route string, handler func(ResponseWriter, *Request[ButtonInteractionData])) {
	m.Mount(ButtonInteraction, route, handler)
}

// Autocomplete mounts an autocomplete route on the mux
func (m *Mux) Autocomplete(route string, handler func(ResponseWriter, *Request[AutocompleteInteractionData])) {
	m.Mount(AutocompleteInteraction, route, handler)
}

// SlashCommand mounts a slash command route on the mux
func (m *Mux) SlashCommand(route string, handler func(ResponseWriter, *Request[SlashCommandInteractionData])) {
	m.Mount(SlashCommandInteraction, route, handler)
}

// UserCommand mounts a user command on the mux
func (m *Mux) UserCommand(route string, handler func(ResponseWriter, *Request[UserCommandInteractionData])) {
	m.Mount(UserCommandInteraction, route, handler)
}

// MessageCommand mounts a message command on the mux
func (m *Mux) MessageCommand(route string, handler func(ResponseWriter, *Request[MessageCommandInteractionData])) {
	m.Mount(MessageCommandInteraction, route, handler)
}
