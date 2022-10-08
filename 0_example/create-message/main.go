package main

import (
	"log"
	"os"

	"github.com/Karitham/corde"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN not set")
	}

	chID := corde.SnowflakeFromString(os.Getenv("DISCORD_CHANNEL_ID"))
	if chID == 0 {
		log.Fatalln("DISCORD_CHANNEL_ID not set")
	}

	m := corde.NewMux("", 0, token)

	message := corde.Message{
		Embeds: []corde.Embed{
			{
				Title:       "hello corde!",
				URL:         "https://github.com/Karitham/corde",
				Description: "corde is awesome :knot:",
			},
		},
	}

	msg, err := m.CreateMessage(chID, message)
	if err != nil {
		log.Fatalln("error creating message: ", err)
	}

	log.Printf("message created: %s", msg.ID)
}
