package corde

import (
	"encoding/json"
	"net/http"
)

// Responder loosely maps to the discord gateway response
// https://discord.com/developers/docs/interactions/receiving-and-responding#responding-to-an-interaction
type Responder struct {
	w http.ResponseWriter
}

// InteractionResponder returns InteractionRespData
type InteractionResponder interface {
	InteractionRespData() *InteractionRespData
}

var _ ResponseWriter = &Responder{}

type intResponse struct {
	Type int                  `json:"type"`
	Data *InteractionRespData `json:"data,omitempty"`
}

// Pong responds to pings on the gateway
func (r *Responder) Pong() {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(intResponse{Type: 1})
}

// Respond responds to the interaction directly
func (r *Responder) Respond(i InteractionResponder) {
	r.respond(intResponse{Type: 4, Data: i.InteractionRespData()})
}

// DeferedRespond responds in defered
func (r *Responder) DeferedRespond(i InteractionResponder) {
	r.respond(intResponse{Type: 5, Data: i.InteractionRespData()})
}

// Update updates the target message
func (r *Responder) Update(i InteractionResponder) {
	r.respond(intResponse{Type: 7, Data: i.InteractionRespData()})
}

// DeferedUpdate updates the target message in defered
func (r *Responder) DeferedUpdate(i InteractionResponder) {
	r.respond(intResponse{Type: 6, Data: i.InteractionRespData()})
}

// Autocomplete responds to the interaction with autocomplete data
func (r *Responder) Autocomplete(i InteractionResponder) {
	r.respond(intResponse{Type: 8, Data: i.InteractionRespData()})
}

func (r *Responder) respond(i intResponse) {
	body, contentType := toBody(i.Data)
	r.w.Header().Set("content-type", contentType)
	body.WriteTo(r.w)
}
