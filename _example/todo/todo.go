package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Karitham/corde"
)

type todo struct {
	mu   sync.Mutex
	list map[string]string
}

func (t *todo) addHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	value := ""
	name := ""
	for _, opt := range i.Data.Options {
		switch opt.Name {
		case "name":
			name = opt.Value.(string)
		case "value":
			value = opt.Value.(string)
		}
	}

	t.list[name] = value

	w.ChannelMessageWithSource(corde.InteractionResponseData{
		Content: corde.Opt(fmt.Sprintf("Sucessfully added %s", name)),
		Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
	})
}

func (t *todo) listHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	s := &strings.Builder{}

	if len(t.list) == 0 {
		s.WriteString("No todos")
	} else {
		s.WriteString("```todo\n")
		for k, v := range t.list {
			s.WriteString(fmt.Sprintf("- %s: %s\n", k, v))
		}
		s.WriteString("```")
	}

	w.ChannelMessageWithSource(corde.InteractionResponseData{
		Embeds: []corde.Embed{
			{
				Title:       "Todo list",
				Description: s.String(),
			},
		},
	})
}

func (t *todo) removeHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	name := ""
	for _, opt := range i.Data.Options {
		if opt.Name == "name" {
			name = opt.Value.(string)
		}
	}

	delete(t.list, name)

	w.ChannelMessageWithSource(corde.InteractionResponseData{
		Content: corde.Opt("deleted todo"),
		Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
	})
}
