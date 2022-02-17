package main

import (
	"log"
	"os"
	"sync"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/snowflake"
)

var command = corde.NewSlashCommand(
	"cmd",
	"edit and view existing slash commands",
	corde.NewSubcommand("list", "list existing slash commands"),
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN not set")
	}
	appID := snowflake.SnowflakeFromString(os.Getenv("DISCORD_APP_ID"))
	if appID == 0 {
		log.Fatalln("DISCORD_APP_ID not set")
	}
	pk := os.Getenv("DISCORD_PUBLIC_KEY")
	if pk == "" {
		log.Fatalln("DISCORD_PUBLIC_KEY not set")
	}

	g := corde.GuildOpt(snowflake.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))

	mu := &sync.Mutex{}
	selectedID := 0

	m := corde.NewMux(pk, appID, token)
	m.Route("cmd", func(m *corde.Mux) {
		m.Route("list", func(m *corde.Mux) {
			m.SlashCommand("", list(m, g))
			m.ButtonComponent("next", btnNext(m, g, mu, &selectedID))
			m.ButtonComponent("remove", btnRemove(m, g, mu, &selectedID))
		})
	})

	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

var nextBtn = components.Component{
	Type:     components.COMPONENT_BUTTON,
	CustomID: "cmd/list/next",
	Style:    components.BUTTON_SECONDARY,
	Label:    "Next",
	Emoji:    &components.Emoji{Name: "➡️"},
}

var delBtn = components.Component{
	Type:     components.COMPONENT_BUTTON,
	CustomID: "cmd/list/remove",
	Style:    components.BUTTON_DANGER,
	Label:    "Delete",
	Emoji:    &components.Emoji{Name: "🗑️"},
}

func list(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Request[components.SlashCommandInteractionData]) {
	return func(w corde.ResponseWriter, _ *corde.Request[components.SlashCommandInteractionData]) {
		w.Respond(components.NewResp().
			ActionRow(nextBtn).
			Ephemeral().
			Content("Click on the buttons to move between existing commands and or delete them."),
		)
	}
}

func btnNext(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Request[components.ButtonInteractionData]) {
	return func(w corde.ResponseWriter, _ *corde.Request[components.ButtonInteractionData]) {
		mu.Lock()
		defer mu.Unlock()
		commands, err := m.GetCommands(g)
		if err != nil {
			w.Update(components.NewResp().Contentf("Error getting commands: %s", err.Error()).Ephemeral())
			return
		}
		if len(commands) == 0 {
			w.Update(components.NewResp().Content("No commands found.").Ephemeral())
			return
		}

		*selectedID = (*selectedID + 1) % len(commands)

		w.Update(components.NewResp().
			Contentf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description).
			ActionRow(nextBtn, delBtn).
			Ephemeral(),
		)
	}
}

func btnRemove(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Request[components.ButtonInteractionData]) {
	return func(w corde.ResponseWriter, _ *corde.Request[components.ButtonInteractionData]) {
		mu.Lock()
		defer mu.Unlock()
		commands, err := m.GetCommands(g)
		if err != nil {
			w.Update(components.NewResp().Contentf("Error getting commands: %s", err.Error()).Ephemeral())
			return
		}
		c := commands[*selectedID]

		m.DeleteCommand(c.ID, g)

		commands, _ = m.GetCommands(g)
		if len(commands) == 0 {
			w.Update(components.NewResp().Content("No commands found.").Ephemeral())
			return
		}

		*selectedID = (*selectedID + 1) % len(commands)

		w.Update(components.NewResp().
			Contentf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description).
			ActionRow(nextBtn, delBtn).
			Ephemeral(),
		)
	}
}
