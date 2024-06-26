package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Karitham/corde"
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

	// user
	if err := m.BulkRegisterCommand(commands, g); err != nil {
		log.Fatalln("error registering command: ", err)
	}

	m.Route("nft", func(m *corde.Mux) {
		m.UserCommand("user", NFTuser)
		m.MessageCommand("message", NFTmessage)
	})

	log.Println("serving on :8070")
	if err := m.ListenAndServe(":8070"); err != nil {
		log.Fatalln(err)
	}
}

func NFTuser(ctx context.Context, w corde.ResponseWriter, i *corde.Interaction[corde.UserCommandInteractionData]) {
	user := i.Data.Resolved.Users.First()
	url := user.AvatarURL()

	if url == "" {
		w.Respond(corde.NewResp().Contentf("error getting %s's profile pic", user.Username).Ephemeral())
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		w.Respond(corde.NewResp().Contentf("error getting %s's profile pic", user.Username).Ephemeral())
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	w.Respond(corde.NewResp().
		Contentf("Good job %s, you just minted %s's profile picture", i.Member.User.Username, user.Username).
		Attachment(resp.Body, filename),
	)
}

func NFTmessage(ctx context.Context, w corde.ResponseWriter, i *corde.Interaction[corde.MessageCommandInteractionData]) {
	msg := i.Data.Resolved.Messages.First()
	chanID := msg.ChannelID
	msgID := msg.ID

	message := fmt.Sprintf("https://discordapp.com/channels/%d/%d/%d", i.GuildID, chanID, msgID)
	w.Respond(corde.NewResp().
		Contentf("Good job %s, you just minted this message, here's the link %s", i.Member.User.Username, message),
	)
}
