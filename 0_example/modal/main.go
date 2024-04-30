package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Karitham/corde"
)

var command = corde.NewSlashCommand("modal", "send a modal")

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

	m := corde.NewMux(pk, appID, token)
	m.SlashCommand("modal", respondModal)
	m.Modal("pog-modal", func(ctx context.Context, w corde.ResponseWriter, r *corde.Interaction[corde.ModalInteractionData]) {
		json.NewEncoder(os.Stderr).Encode(r)
		w.DeferedUpdate()
	})

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func respondModal(ctx context.Context, w corde.ResponseWriter, r *corde.Interaction[corde.SlashCommandInteractionData]) {
	w.Modal(corde.Modal{
		Title:    "xoxo",
		CustomID: "pog-modal",
		Components: []corde.Component{
			corde.TextInputComponent{
				CustomID:    "pog-component",
				Style:       corde.TEXT_PARAGRAPH,
				Label:       "label",
				Required:    false,
				Placeholder: "placeholder",
			}.Component(),
		},
	})
}
