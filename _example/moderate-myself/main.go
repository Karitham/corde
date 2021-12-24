package main

import (
	"fmt"
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
	m.SetRoute(corde.SlashCommand("cmd/list"), list(m, g))
	m.SetRoute(corde.SlashCommand("cmd/remove"), rm(m, g))
	m.SetRoute(corde.ButtonInteraction("btn-cmd/list/next"), btnNext(m, g, mu, &selectedID))
	m.SetRoute(corde.ButtonInteraction("btn-cmd/list/remove"), btnRemove(m, g, mu, &selectedID))

	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func list(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		w.WithSource(&corde.InteractionRespData{
			Components: []corde.Component{
				{
					Type: corde.COMPONENT_ACTION_ROW,
					Components: []corde.Component{
						{
							Type:     corde.COMPONENT_BUTTON,
							CustomID: "btn-cmd/list/next",
							Style:    corde.BUTTON_SECONDARY,
							Label:    "next",
							Emoji:    &corde.Emoji{Name: "➡️"},
						},
					},
				},
			},
			Content: fmt.Sprintf("Click on the buttons to move between existing commands and or delete them."),
			Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
		})
	}
}

func rm(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		n, ok := i.Data.Options["name"]
		if !ok {
			w.WithSource(&corde.InteractionRespData{
				Content: "Please enter an actual command name.",
				Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
			})
			return
		}

		c, _ := m.GetCommands(g)
		for _, c := range c {
			if c.Name == n {
				m.DeleteCommand(c.ID, g)
				w.WithSource(&corde.InteractionRespData{
					Content: fmt.Sprintf("Removed command %s.", n),
					Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
				})

				return
			}
		}

		w.WithSource(&corde.InteractionRespData{
			Content: fmt.Sprintf("No command named %s found.", n),
			Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
		})
	}
}

func btnNext(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		mu.Lock()
		defer mu.Unlock()
		commands, _ := m.GetCommands(g)
		if len(commands) == 0 {
			w.UpdateMessage(&corde.InteractionRespData{
				Content: "No commands found.",
				Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
			})
			return
		}
		*selectedID = *selectedID + 1%(len(commands))

		w.UpdateMessage(&corde.InteractionRespData{
			Content: fmt.Sprintf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description),
			Components: []corde.Component{
				{
					Type: corde.COMPONENT_ACTION_ROW,
					Components: []corde.Component{
						{
							Type:     corde.COMPONENT_BUTTON,
							CustomID: "btn-cmd/list/next",
							Style:    corde.BUTTON_SECONDARY,
							Label:    "next",
							Emoji:    &corde.Emoji{Name: "➡️"},
						},
						{
							Type:     corde.COMPONENT_BUTTON,
							CustomID: "btn-cmd/list/remove",
							Style:    corde.BUTTON_DANGER,
							Label:    "remove",
							Emoji:    &corde.Emoji{Name: "🗑️"},
						},
					},
				},
			},
			Flags: corde.RESPONSE_FLAGS_EPHEMERAL,
		})
	}
}

func btnRemove(m *corde.Mux, g func(*corde.CommandsOpt), mu *sync.Mutex, selectedID *int) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		mu.Lock()
		defer mu.Unlock()
		commands, _ := m.GetCommands(g)
		c := commands[*selectedID]

		m.DeleteCommand(c.ID, g)

		w.UpdateMessage(&corde.InteractionRespData{
			Content: fmt.Sprintf("%s - %s", commands[*selectedID].Name, commands[*selectedID].Description),
			Components: []corde.Component{
				{
					Type: corde.COMPONENT_ACTION_ROW,
					Components: []corde.Component{
						{
							Type:     corde.COMPONENT_BUTTON,
							CustomID: "btn-cmd/list/next",
							Style:    corde.BUTTON_SECONDARY,
							Label:    "next",
							Emoji:    &corde.Emoji{Name: "➡️"},
						},
						{
							Type:     corde.COMPONENT_BUTTON,
							CustomID: "btn-cmd/list/remove",
							Style:    corde.BUTTON_DANGER,
							Label:    "remove",
							Emoji:    &corde.Emoji{Name: "🗑️"},
						},
					},
				},
			},
			Flags: corde.RESPONSE_FLAGS_EPHEMERAL,
		})
	}
}
