package main

import (
	"log"
	"os"
	"sync"

	"github.com/Karitham/corde"
)

var commands = []corde.CreateCommander{
	corde.NewSlashCommand(
		"todo",
		"view edit and remove todos",
		corde.NewSubcommand(
			"list",
			"list todos",
			false,
			corde.NewStringOption("name", "todo name", true),
			corde.NewStringOption("value", "todo value", true),
		),
		corde.NewSubcommand("add", "add a todo", false),
		corde.NewSubcommand(
			"rm",
			"remove a todo",
			false,
			corde.NewStringOption("name", "todo name", true),
		),
	),
}

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
		list: make(map[string]string),
	}

	m := corde.NewMux(pk, appID, token)
	m.Route("todo", func(m *corde.Mux) {
		m.Command("add", t.addHandler)
		m.Command("rm", t.removeHandler)
		m.Command("list", t.listHandler)
	})

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	if err := m.BulkRegisterCommand(commands, g); err != nil {
		log.Fatalln(err)
	}

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}
