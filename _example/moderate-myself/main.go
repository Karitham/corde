package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	m := corde.NewMux(pk, appID, token)
	m.SetRoute(corde.SlashCommand("cmd/list"), list(m, g))
	m.SetRoute(corde.SlashCommand("cmd/remove"), rm(m, g))

	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func list(m *corde.Mux, g func(*corde.CommandsOpt)) func(corde.ResponseWriter, *corde.Interaction) {
	return func(w corde.ResponseWriter, i *corde.Interaction) {
		c, _ := m.GetCommands(g)

		w.WithSource(&corde.InteractionRespData{
			Content: func() string {
				s := &strings.Builder{}
				s.WriteString("```\n")
				for _, c := range c {
					s.WriteString(fmt.Sprintf("%s: %s\n", c.Name, c.Description))
				}
				s.WriteString("```")
				return s.String()
			}(),
			Flags: corde.RESPONSE_FLAGS_EPHEMERAL,
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
