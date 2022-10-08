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

func (t *todo) autoCompleteNames(w corde.ResponseWriter, _ *corde.Request[corde.AutocompleteInteractionData]) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		w.Autocomplete(corde.NewResp())
		return
	}

	resp := corde.NewResp()
	for k := range t.list {
		resp.Choice(k, k)
	}

	w.Autocomplete(resp)
}

func (t *todo) addHandler(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	value, _ := i.Data.Options.String("value")
	name, _ := i.Data.Options.String("name")

	user, err := i.Data.OptionsUser("user")
	if err != nil {
		user = i.Member.User
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.list[name] = todoItem{
		user:  user.ID,
		value: value,
	}

	w.Respond(corde.NewResp().Contentf("Successfully added %s", name).Ephemeral())
}

func (t *todo) listHandler(w corde.ResponseWriter, _ *corde.Request[corde.SlashCommandInteractionData]) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.list) == 0 {
		w.Respond(corde.NewResp().Content("no todos").Ephemeral())
		return
	}

	i := 1
	s := &strings.Builder{}
	for k, v := range t.list {
		s.WriteString(fmt.Sprintf("%d. %s: %s - %s\n", i, k, v.value, format.User(v.user)))
		i++
	}

	w.Respond(corde.NewEmbed().
		Title("Todo list").
		Description(s.String()).
		Color(0x69b00b),
	)
}

func (t *todo) removeHandler(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	t.mu.Lock()
	defer t.mu.Unlock()

	name, _ := i.Data.Options.String("name")
	if _, ok := t.list[name]; !ok {
		w.Respond(corde.NewResp().Contentf("%s not found", name).Ephemeral())
		return
	}

	delete(t.list, name)
	w.Respond(corde.NewResp().Contentf("deleted todo %s", name).Ephemeral())
}
