package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	root := &ffcli.Command{
		Name:       "owmock",
		ShortUsage: "`owmock <command>` or `owmock <uri>`",
		ShortHelp:  "owmock is a mock server for outgoing webhooks",
		Subcommands: []*ffcli.Command{
			{Name: "public-key", Exec: publicKey},
			{Name: "private-key", Exec: privateKey},
		},
		Exec: REPL,
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}

// GenerateKeys generates a new keypair of hex-encoded keys
// They shouldn't change
// PubK: 2f8c6129d816cf51c374bc7f08c3e63ed156cf78aefb4a6550d97b87997977ee
// PrivK: 31323334353637383930313233343536373839303132333435363738393031322f8c6129d816cf51c374bc7f08c3e63ed156cf78aefb4a6550d97b87997977ee
func GenerateKeys() (string, string) {
	pubK, privK, _ := ed25519.GenerateKey(bytes.NewBufferString("12345678901234567890123456789012"))
	return hex.EncodeToString(pubK), hex.EncodeToString(privK)
}

func publicKey(_ context.Context, _ []string) error {
	pub, _ := GenerateKeys()
	fmt.Println(pub)
	return nil
}

func privateKey(_ context.Context, _ []string) error {
	_, priv := GenerateKeys()
	fmt.Println(priv)
	return nil
}

// Read Eval Print Loop
// But actually it's just to use as a mock server
func REPL(_ context.Context, args []string) error {
	if len(args) == 0 {
		return errors.New("no outgoing uri specified")
	}
	return nil
}
