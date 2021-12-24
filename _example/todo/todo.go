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
	value, ok := i.Data.Options["value"].(string)
	if !ok {
		ephemeral(w, "no value provided")
		return
	}

	name, ok := i.Data.Options["name"].(string)
	if !ok {
		ephemeral(w, "no name provided")
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	t.list[name] = value

	ephemeral(w, fmt.Sprintf("Sucessfully added %s", name))
}

func (t *todo) listHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		ephemeral(w, "no todos")
		return
	}

	// build todo list
	s := &strings.Builder{}
	s.WriteString("```todo\n")
	for k, v := range t.list {
		s.WriteString(fmt.Sprintf("- %s: %s\n", k, v))
	}
	s.WriteString("```")

	w.WithSource(&corde.InteractionRespData{
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

	name, ok := i.Data.Options["name"].(string)
	if !ok {
		ephemeral(w, "no name provided")
		return
	}

	delete(t.list, name)
	ephemeral(w, "deleted todo")
}

// ephemeral returns an ephemeral response
func ephemeral(w corde.ResponseWriter, message string) {
	w.WithSource(&corde.InteractionRespData{
		Content: message,
		Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
	})
}
