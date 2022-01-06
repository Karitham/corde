package main

import (
	"log"
	"os"
	"sync"

	"github.com/Karitham/corde"
)

var commands = corde.NewSlashCommand("todo", "view edit and remove todos",
	corde.NewSubcommand("list", "list todos"),
	corde.NewSubcommand("add", "add a todo",
		corde.NewStringOption("name", "todo name", true),
		corde.NewStringOption("value", "todo value", true),
		corde.NewUserOption("user", "assign it to a user", false),
	),
	corde.NewSubcommand("rm", "remove a todo",
		corde.NewStringOption("name", "todo name", true).CanAutocomplete(),
	),
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN not set")
	}
	appID := corde.SnowflakeFromString(os.Getenv("DISCORD_APP_ID"))
	if appID == 0 {
		log.Fatalln("DISCORD_APP_ID not set")
	}
	pk := os.Getenv("DISCORD_PUBLIC_KEY")
	if pk == "" {
		log.Fatalln("DISCORD_PUBLIC_KEY not set")
	}

	t := todo{
		mu:   sync.Mutex{},
		list: make(map[string]todoItem),
	}

	m := corde.NewMux(pk, appID, token)
	m.Route("todo", func(m *corde.Mux) {
		m.Command("add", t.addHandler)
		m.Command("list", t.listHandler)
		m.Route("rm", func(m *corde.Mux) {
			m.Command("", t.removeHandler)
			m.Autocomplete("", t.autoCompleteNames)
		}
	})

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	if err := m.RegisterCommand(commands, g); err != nil {
		log.Fatalln(err)
	}

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}
