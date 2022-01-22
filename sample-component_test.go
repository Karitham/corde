package corde_test

import (
	"net/http/httptest"
	"testing"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/owmock"
	"github.com/matryer/is"
)

func TestComponentInteraction(t *testing.T) {
	assert := is.New(t)
	pub, _ := owmock.GenerateKeys()
	mux := corde.NewMux(pub, 0, "")

	mux.Button("click_one", func(w corde.ResponseWriter, _ *corde.InteractionRequest) {
		w.Respond(&corde.InteractionRespData{
			Content: "Hello World!",
		})
	})

	expect := &owmock.InteractionResponse{
		Type: 4,
		Data: corde.InteractionRespData{
			Content: "Hello World!",
		},
	}

	s := httptest.NewServer(mux)
	err := owmock.NewWithClient(s.URL, s.Client()).PostExpect(t, SampleComponent, expect)
	assert.NoErr(err)
}

const SampleComponent = `{
	"version": 1,
    "type": 3,
    "token": "unique_interaction_token",
    "message": {
        "type": 0,
        "tts": false,
        "timestamp": "2021-05-19T02:12:51.710000+00:00",
        "pinned": false,
        "mentions": [],
        "mention_roles": [],
        "mention_everyone": false,
        "id": "844397162624450620",
        "flags": 0,
        "embeds": [],
        "edited_timestamp": null,
        "content": "This is a message with components.",
        "components": [
            {
                "type": 1,
                "components": [
                    {
                        "type": 2,
                        "label": "Click me!",
                        "style": 1,
                        "custom_id": "click_one"
                    }
                ]
            }
        ],
        "channel_id": "345626669114982402",
        "author": {
            "username": "Mason",
            "public_flags": 131141,
            "id": "53908232506183680",
            "discriminator": "1337",
            "avatar": "a_d5efa99b3eeaa7dd43acca82f5692432"
        },
        "attachments": []
    },
    "member": {
        "user": {
            "username": "Mason",
            "public_flags": 131141,
            "id": "53908232506183680",
            "discriminator": "1337",
            "avatar": "a_d5efa99b3eeaa7dd43acca82f5692432"
        },
        "roles": [
            "290926798626357999"
        ],
        "premium_since": null,
        "permissions": "17179869183",
        "pending": false,
        "nick": null,
        "mute": false,
        "joined_at": "2017-03-13T19:19:14.040000+00:00",
        "is_pending": false,
        "deaf": false,
        "avatar": null
    },
    "id": "846462639134605312",
    "guild_id": "290926798626357999",
    "data": {
        "custom_id": "click_one",
        "component_type": 2
    },
    "channel_id": "345626669114982999",
    "application_id": "290926444748734465"
}`
