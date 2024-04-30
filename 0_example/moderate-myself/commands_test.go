package main

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/Karitham/corde"

	"github.com/Karitham/corde/owmock"
)

func Test_list(t *testing.T) {
	list := list(nil, nil)

	tests := []struct {
		name        string
		mock        owmock.ResponseWriterMock
		interaction *corde.Interaction[corde.SlashCommandInteractionData]
	}{
		{
			name: "list",
			mock: owmock.ResponseWriterMock{
				RespondHook: func(i corde.InteractionResponder) {
					data := i.InteractionRespData()
					if data.Content != "Click on the buttons to move between existing commands and or delete them." {
						t.Errorf("expected 'no todos' got %s", data.Content)
					}
					if data.Flags != corde.RESPONSE_FLAGS_EPHEMERAL {
						t.Errorf("expected ephemeral flag got %d", data.Flags)
					}
				},
			},
			interaction: &corde.Interaction[corde.SlashCommandInteractionData]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			list(context.Background(), tt.mock, tt.interaction)
		})
	}
}

func Test_btnNext(t *testing.T) {
	selectedID := 0
	tests := []struct {
		name        string
		mock        owmock.ResponseWriterMock
		interaction *corde.Interaction[corde.ButtonInteractionData]
		fn          func(context.Context, corde.ResponseWriter, *corde.Interaction[corde.ButtonInteractionData])
	}{
		{
			name: "btn next",
			mock: owmock.ResponseWriterMock{
				T: t,
				UpdateHook: func(i corde.InteractionResponder) {
					data := i.InteractionRespData()
					if data.Content == "" {
						t.Error("expected some sort of response")
					}

					if data.Flags != corde.RESPONSE_FLAGS_EPHEMERAL {
						t.Errorf("expected ephemeral flag got %d", data.Flags)
					}
				},
			},
			interaction: &corde.Interaction[corde.ButtonInteractionData]{},
			fn:          btnNext(&corde.Mux{Client: http.DefaultClient}, corde.GuildOpt(0), &sync.Mutex{}, &selectedID),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.fn(context.Background(), tt.mock, tt.interaction)
		})
	}
}
