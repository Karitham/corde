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
	value := i.Data.Options.String("value")
	name := i.Data.Options.String("name")

	t.mu.Lock()
	defer t.mu.Unlock()
	t.list[name] = value

	w.Respond(corde.NewResp().Contentf("Sucessfully added %s", name).Ephemeral().B())
}

func (t *todo) listHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		w.Respond(corde.NewResp().Content("no todos").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().
		Embeds(corde.NewEmbed().
			Title("Todo list").
			// build todo list description
			Description(func() string {
				s := &strings.Builder{}
				s.WriteString("```todo\n")
				i := 0
				for k, v := range t.list {
					i++
					s.WriteString(fmt.Sprintf("%d. %s: %s\n", i, k, v))
				}
				s.WriteString("```")
				return s.String()
			}()).
			Color(0x69b00b).
			B(),
		).B(),
	)
}

func (t *todo) removeHandler(w corde.ResponseWriter, i *corde.Interaction) {
	t.mu.Lock()
	defer t.mu.Unlock()

	name := i.Data.Options.String("name")

	delete(t.list, name)
	w.Respond(corde.NewResp().Content("deleted todo").Ephemeral().B())
}
