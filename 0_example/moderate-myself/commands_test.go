package main

import (
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
		interaction *corde.InteractionRequest
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
			interaction: &corde.InteractionRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			list(tt.mock, tt.interaction)
		})
	}
}

func Test_btnNext(t *testing.T) {
	selectedID := 0
	tests := []struct {
		name        string
		mock        owmock.ResponseWriterMock
		interaction *corde.InteractionRequest
		fn          func(corde.ResponseWriter, *corde.InteractionRequest)
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
			interaction: &corde.InteractionRequest{},
			fn:          btnNext(&corde.Mux{Client: http.DefaultClient}, corde.GuildOpt(0), &sync.Mutex{}, &selectedID),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			tt.fn(tt.mock, tt.interaction)
		})
	}
}
