package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/format"
	"github.com/Karitham/corde/snowflake"
)

type todo struct {
	mu   sync.Mutex
	list map[string]todoItem
}

type todoItem struct {
	user  snowflake.Snowflake
	value string
}

func (t *todo) autoCompleteNames(w corde.ResponseWriter, _ *corde.InteractionRequest) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		w.Autocomplete(components.NewResp())
		return
	}

	resp := components.NewResp()
	for k := range t.list {
		resp.Choice(k, k)
	}

	w.Autocomplete(resp)
}

func (t *todo) addHandler(w corde.ResponseWriter, i *corde.InteractionRequest) {
	data, _ := components.GetInteractionData[components.SlashInteractionData](i.Interaction)
	value, _ := data.Options.String("value")
	name, _ := data.Options.String("name")

	user, err := data.OptionsUser("user")
	if err != nil {
		user = i.Member.User
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.list[name] = todoItem{
		user:  user.ID,
		value: value,
	}

	w.Respond(components.NewResp().Contentf("Successfully added %s", name).Ephemeral())
}

func (t *todo) listHandler(w corde.ResponseWriter, _ *corde.InteractionRequest) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		w.Respond(components.NewResp().Content("no todos").Ephemeral())
		return
	}

	i := 1
	s := &strings.Builder{}
	for k, v := range t.list {
		s.WriteString(fmt.Sprintf("%d. %s: %s - %s\n", i, k, v.value, format.User(v.user)))
		i++
	}

	w.Respond(components.NewEmbed().
		Title("Todo list").
		Description(s.String()).
		Color(0x69b00b),
	)
}

func (t *todo) removeHandler(w corde.ResponseWriter, i *corde.InteractionRequest) {
	data, _ := components.GetInteractionData[components.SlashInteractionData](i.Interaction)
	t.mu.Lock()
	defer t.mu.Unlock()

	name, _ := data.Options.String("name")
	if _, ok := t.list[name]; !ok {
		w.Respond(components.NewResp().Contentf("%s not found", name).Ephemeral())
		return
	}

	delete(t.list, name)
	w.Respond(components.NewResp().Contentf("deleted todo %s", name).Ephemeral())
}
