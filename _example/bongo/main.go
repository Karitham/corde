package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Karitham/corde"
)

var command = corde.Command{
	Name:        "bongo",
	Description: "send a big bongo",
	Type:        corde.COMMAND_CHAT_INPUT,
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

	m := corde.NewMux(pk, appID, token)
	m.Command("bongo", bongoHandler)

	g := corde.GuildOpt(corde.SnowflakeFromString(os.Getenv("DISCORD_GUILD_ID")))
	if err := m.RegisterCommand(command, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func bongoHandler(w corde.ResponseWriter, i *corde.Interaction) {
	resp, err := http.Get("https://cdn.discordapp.com/emojis/745709799890747434.gif?size=128")
	if err != nil {
		w.Respond(corde.NewResp().Content("couldn't retrieve bongo").Ephemeral().B())
		return
	}
	defer resp.Body.Close()
	w.Respond(corde.NewResp().
		Attachments(corde.Attachment{
			Body:     resp.Body,
			ID:       corde.Snowflake(0),
			Filename: "bongo.gif",
		}).B(),
	)
}
