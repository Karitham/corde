package corde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Responder struct {
	w http.ResponseWriter
}

var _ ResponseWriter = &Responder{}

type intResponse struct {
	Type int                  `json:"type"`
	Data *InteractionRespData `json:"data,omitempty"`
}

// pong responds to pings on the gateway
func (r *Responder) pong() {
	r.w.Header().Set("content-type", "application/json")
	json.NewEncoder(r.w).Encode(intResponse{Type: 1})
}

// WithSource responds to the interaction directly
func (r *Responder) WithSource(i *InteractionRespData) {
	r.respond(intResponse{Type: 4, Data: i})
}

// DeferedWithSource responds in defered
func (r *Responder) DeferedWithSource(i *InteractionRespData) {
	r.respond(intResponse{Type: 5, Data: i})
}

// UpdateMessage updates the target message
func (r *Responder) UpdateMessage(i *InteractionRespData) {
	r.respond(intResponse{Type: 7, Data: i})
}

// DeferedUpdateMessage updates the target message in defered
func (r *Responder) DeferedUpdateMessage(i *InteractionRespData) {
	r.respond(intResponse{Type: 6, Data: i})
}

// AutocompleteResult responds to the interaction with autocomplete data
func (r *Responder) AutocompleteResult(i *InteractionRespData) {
	r.respond(intResponse{Type: 8, Data: i})
}

func (r *Responder) respond(i intResponse) {
	payloadJSON := &bytes.Buffer{}
	err := json.NewEncoder(payloadJSON).Encode(i)
	if err != nil {
		return
	}

	if len(i.Data.Attachements) < 1 {
		r.w.Header().Set("content-type", "application/json")
		payloadJSON.WriteTo(r.w)
		return
	}

	mw := multipart.NewWriter(r.w)
	defer mw.Close()

	r.w.Header().Set("Content-Type", mw.FormDataContentType())
	mw.WriteField("payload_json", payloadJSON.String())

	for i, f := range i.Data.Attachements {
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
