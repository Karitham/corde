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
	message := corde.NewEmbed().
		Title("Hello corde!").
		URL("https://github.com/Karitham/corde").
		Color(0xffffff).
		Description("corde is awesome :knot:").
		Message()

	msg, err := m.CreateMessage(chID, message)
	if err != nil {
		log.Fatalln("error creating message: ", err)
	}

	log.Printf("message created: %s", msg.ID)
}
