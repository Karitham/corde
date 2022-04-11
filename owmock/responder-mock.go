package owmock

import (
	"testing"

	"github.com/Karitham/corde"
)

// ResponseWriterMock mocks corde's ResponseWriter interface
type ResponseWriterMock struct {
	RespondHook      func(corde.InteractionResponder)
	UpdateHook       func(corde.InteractionResponder)
	AutocompleteHook func(corde.InteractionResponder)
	ModalHook        func(corde.Modal)

	T *testing.T
}

// NewRWMock returns a new ResponseWriterMock with the given testing.T
func NewRWMock(t *testing.T) ResponseWriterMock {
	return ResponseWriterMock{T: t}
}

// Pong implements ResponseWriter interface
func (r ResponseWriterMock) Ack() {}

// Response implements ResponseWriter interface
func (r ResponseWriterMock) Respond(i corde.InteractionResponder) {
	if r.RespondHook != nil {
		r.RespondHook(i)
		return
	}

	r.T.Error("unexpected respond hook called")
}

// DeferedRespond implements ResponseWriter interface
func (r ResponseWriterMock) DeferedRespond() {}

// Update implements ResponseWriter interface
func (r ResponseWriterMock) Update(i corde.InteractionResponder) {
	if r.UpdateHook != nil {
		r.UpdateHook(i)
		return
	}

	r.T.Error("unexpected update hook called")
}

// DeferedUpdate implements ResponseWriter interface
func (r ResponseWriterMock) DeferedUpdate() {}

// Autocomplete implements ResponseWriter interface
func (r ResponseWriterMock) Autocomplete(i corde.InteractionResponder) {
	if r.AutocompleteHook != nil {
		r.AutocompleteHook(i)
		return
	}

	r.T.Error("unexpected autocomplete hook called")
}

// Modal implements ResponseWriter interface
func (r ResponseWriterMock) Modal(m corde.Modal) {
	if r.ModalHook != nil {
		r.ModalHook(m)
		return
	}

	r.T.Error("unexpected modal hook called")
}

// type	interaction callback type	the type of response
// data?	interaction callback data	an optional response message
type InteractionResponse struct {
	Type corde.InteractionType     `json:"type"`
	Data corde.InteractionRespData `json:"data,omitempty"`
}
