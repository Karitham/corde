package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/snowflake"
)

// spaces are routed as `/`
// User & Messages are considered commands, so you can mount them using `m.Command`
var commands = []corde.CreateCommander{
	corde.NewMessageCommand("nft message"),
	corde.NewUserCommand("nft user"),
}

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
	m := corde.NewMux(pk, appID, token)

	// user
	if err := m.BulkRegisterCommand(commands, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	m.Route("nft", func(m *corde.Mux) {
		m.Command("user", NFTuser)
		m.Command("message", NFTmessage)
	})

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func NFTuser(w corde.ResponseWriter, i *corde.InteractionRequest) {
	data, _ := components.GetInteractionData[components.UserCommandInteractionData](i.Interaction)
	user := data.Resolved.Users.First()
	url := user.AvatarURL()

	if url == "" {
		w.Respond(components.NewResp().Contentf("error getting %s's profile pic", user.Username).Ephemeral())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		w.Respond(components.NewResp().Contentf("error getting %s's profile pic", user.Username).Ephemeral())
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	w.Respond(components.NewResp().
		Contentf("Good job %s, you just minted %s's profile picture", i.Member.User.Username, user.Username).
		Attachment(resp.Body, filename),
	)
}

func NFTmessage(w corde.ResponseWriter, i *corde.InteractionRequest) {
	data, _ := components.GetInteractionData[components.MessageCommandInteractionData](i.Interaction)

	msg := data.Resolved.Messages.First()
	chanID := msg.ChannelID
	msgID := msg.ID

	message := fmt.Sprintf("https://discordapp.com/channels/%d/%d/%d", i.GuildID, chanID, msgID)
	w.Respond(components.NewResp().
		Contentf("Good job %s, you just minted this message, here's the link %s", i.Member.User.Username, message),
	)
}
