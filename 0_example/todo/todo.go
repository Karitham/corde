package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/format"
)

type todo struct {
	mu   sync.Mutex
	list map[string]todoItem
}

type todoItem struct {
	user  corde.Snowflake
	value string
}

func (t *todo) addHandler(w corde.ResponseWriter, i *corde.Interaction) {
	value := i.Data.Options.String("value")
	name := i.Data.Options.String("name")

	user := i.Data.Options.Snowflake("user")
	if user == 0 {
		user = i.User.ID
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.list[name] = todoItem{
		user:  user,
		value: value,
	}

	w.Respond(corde.NewResp().Contentf("Sucessfully added %s", name).Ephemeral().B())
}

func (t *todo) listHandler(w corde.ResponseWriter, _ *corde.Interaction) {
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
				s, i := &strings.Builder{}, 1
				for k, v := range t.list {
					s.WriteString(fmt.Sprintf("%d. %s: %s - %s\n", i, k, v.value, format.User(v.user.String())))
					i++
				}
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