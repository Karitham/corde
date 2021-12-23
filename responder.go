package corde

import (
	"encoding/json"
	"net/http"
)

type Responder struct {
	w http.ResponseWriter
}

var _ ResponseWriter = &Responder{}

type InteractionResponse struct {
	Type int                      `json:"type"`
	Data *InteractionResponseData `json:"data,omitempty"`
}

func (r Responder) Pong() {
	r.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r.w).Encode(InteractionResponse{Type: 1})
}

// TODO(@Karitham): Handle sending files
func (r Responder) ChannelMessageWithSource(i InteractionResponseData) {
	r.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r.w).Encode(InteractionResponse{Type: 4, Data: &i})
}
