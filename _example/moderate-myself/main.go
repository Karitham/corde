package main

import (
	"log"
	"os"
	"sync"

	"github.com/Karitham/corde"
)

var command = corde.Command{
	Name:        "cmd",
	Description: "edit and view existing slash commands",
	Type:        corde.COMMAND_CHAT_INPUT,
	Options: []corde.Option{
		{
			Name:        "list",
			Type:        corde.OPTION_SUB_COMMAND,
			Description: "list existing slash commands",
		},
		{
			Name:        "remove",
			Type:        corde.OPTION_SUB_COMMAND,
			Description: "remove slash commands",
			Options: []corde.Option{
				{
					Name:        "name",
					Type:        corde.OPTION_STRING,
					Description: "name of the slash command you wish to remove",
					Required:    true,
				},
			},
		},
	},
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

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))

	mu := &sync.Mutex{}
	selectedID := 0

	m := corde.NewMux(pk, appID, token)
	m.Mount(corde.SlashCommand("cmd/list"), list(m, g))
	m.Mount(corde.SlashCommand("cmd/remove"), rm(m, g))
	m.Mount(corde.ButtonInteraction("btn-cmd/list/next"), btnNext(m, g, mu, &selectedID))
	m.Mount(corde.ButtonInteraction("btn-cmd/list/remove"), btnRemove(m, g, mu, &selectedID))

	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

var nextBtn = corde.Component{
	Type:     corde.COMPONENT_BUTTON,
	CustomID: "btn-cmd/list/next",
	Style:    corde.BUTTON_SECONDARY,
	Label:    "next",
	Emoji:    &corde.Emoji{Name: "➡️"},
}

var delBtn = corde.Component{
	Type:     corde.COMPONENT_BUTTON,
	CustomID: "btn-cmd/list/remove",
	Style:    corde.BUTTON_DANGER,
	Label:    "remove",
	Emoji:    &corde.Emoji{Name: "🗑️"},
}

func list(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		w.Respond(corde.NewResp().
			ActionRow(nextBtn).
			Ephemeral().
			Content("Click on the buttons to move between existing commands and or delete them.").
			B(),
		)
	}
}

func rm(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		n := i.Data.Options.String("name")

		c, _ := m.GetCommands(g)
		for _, c := range c {
			if c.Name == n {
				m.DeleteCommand(c.ID, g)
				w.Respond(corde.NewResp().Content("No command named %s found.", n).Ephemeral().B())
				return
			}
		}

		w.Respond(corde.NewResp().Contentf("No command named %s found.", n).Ephemeral().B())
	}
}

func btnNext(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		mu.Lock()
		defer mu.Unlock()
		commands, _ := m.GetCommands(g)
		if len(commands) == 0 {
			w.Update(corde.NewResp().Content("No commands found.").Ephemeral().B())
			return
		}

		*selectedID = (*selectedID + 1) % len(commands)

		w.Update(corde.NewResp().
			Contentf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description).
			ActionRow(nextBtn, delBtn).
			Ephemeral().
			B(),
		)
	}
}

func btnRemove(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		mu.Lock()
		defer mu.Unlock()
		commands, _ := m.GetCommands(g)
		c := commands[*selectedID]

		m.DeleteCommand(c.ID, g)

		commands, _ = m.GetCommands(g)
		*selectedID = (*selectedID + 1) % len(commands)

		w.Update(corde.NewResp().
			Contentf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description).
			ActionRow(nextBtn, delBtn).
			Ephemeral().
			B(),
		)
	}
}
