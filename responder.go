package corde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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

var _ ResponseWriter = (*Responder)(nil)

type intResponse struct {
	Type int                  `json:"type"`
	Data *InteractionRespData `json:"data,omitempty"`
}

// Pong responds to pings on the gateway
func (r *Responder) Ack() {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(intResponse{Type: 1})
}

// Respond responds to the interaction directly
func (r *Responder) Respond(i InteractionResponder) {
	r.respond(intResponse{Type: 4, Data: i.InteractionRespData()})
}

// DeferedRespond responds in defered
func (r *Responder) DeferedRespond() {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(intResponse{Type: 5})
}

// Update updates the target message
func (r *Responder) Update(i InteractionResponder) {
	r.respond(intResponse{Type: 7, Data: i.InteractionRespData()})
}

// DeferedUpdate updates the target message in defered
func (r *Responder) DeferedUpdate() {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(intResponse{Type: 6})
}

// Autocomplete responds to the interaction with autocomplete data
func (r *Responder) Autocomplete(i InteractionResponder) {
	r.respond(intResponse{Type: 8, Data: i.InteractionRespData()})
}

// Modal responds to the interaction with modal data
func (r *Responder) Modal(m Modal) {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(
		struct {
			Type int   `json:"type"`
			Data Modal `json:"data"`
		}{
			Type: 9,
			Data: m,
		},
	)
}

func (r *Responder) respond(i intResponse) {
	payloadJSON := &bytes.Buffer{}
	err := json.NewEncoder(payloadJSON).Encode(i)
	if err != nil {
		return
	}

	if len(i.Data.Attachments) < 1 {
		r.w.Header().Set("content-type", "application/json")
		payloadJSON.WriteTo(r.w)
		return
	}

	mw := multipart.NewWriter(r.w)
	defer mw.Close()

	contentType := mw.FormDataContentType()
	r.w.Header().Set("content-type", contentType)
	mw.WriteField("payload_json", payloadJSON.String())

	for i, f := range i.Data.Attachments {
		if f.ID == 0 {
			f.ID = Snowflake(i)
		}

		ff, CFerr := mw.CreateFormFile(fmt.Sprintf("files[%d]", i), f.Filename)
		if CFerr != nil {
			return
		}

		if _, CopyErr := io.Copy(ff, f.Body); CopyErr != nil {
			return
		}
	}
}
