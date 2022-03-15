package owmock

import (
	"testing"

	"github.com/Karitham/corde"
)

// ResponseWriterMock mocks corde's ResponseWriter interface
type ResponseWriterMock struct {
	RespondHook        func(corde.InteractionResponder)
	DeferedRespondHook func(corde.InteractionResponder)
	UpdateHook         func(corde.InteractionResponder)
	DeferedUpdateHook  func(corde.InteractionResponder)
	AutocompleteHook   func(corde.InteractionResponder)

	T *testing.T
}

// NewRWMock returns a new ResponseWriterMock with the given testing.T
func NewRWMock(t *testing.T) ResponseWriterMock {
	return ResponseWriterMock{T: t}
}

// Pong implements ResponseWriter interface
func (r ResponseWriterMock) Pong() {}

// Response implements ResponseWriter interface
func (r ResponseWriterMock) Respond(i corde.InteractionResponder) {
	if r.RespondHook != nil {
		r.RespondHook(i)
		return
	}

	r.T.Error("unexpected respond hook called")
}

// DeferedRespond implements ResponseWriter interface
func (r ResponseWriterMock) DeferedRespond(i corde.InteractionResponder) {
	if r.DeferedRespondHook != nil {
		r.DeferedRespondHook(i)
		return
	}

	r.T.Error("unexpected defered respond hook called")
}

// Update implements ResponseWriter interface
func (r ResponseWriterMock) Update(i corde.InteractionResponder) {
	if r.UpdateHook != nil {
		r.UpdateHook(i)
		return
	}

	r.T.Error("unexpected update hook called")
}

// DeferedUpdate implements ResponseWriter interface
func (r ResponseWriterMock) DeferedUpdate(i corde.InteractionResponder) {
	if r.DeferedUpdateHook != nil {
		r.DeferedUpdateHook(i)
		return
	}

	r.T.Error("unexpected defered update hook called")
}

// Autocomplete implements ResponseWriter interface
func (r ResponseWriterMock) Autocomplete(i corde.InteractionResponder) {
	if r.AutocompleteHook != nil {
		r.AutocompleteHook(i)
		return
	}

	r.T.Error("unexpected autocomplete hook called")
}

// type	interaction callback type	the type of response
// data?	interaction callback data	an optional response message
type InteractionResponse struct {
	Type corde.InteractionType     `json:"type"`
	Data corde.InteractionRespData `json:"data,omitempty"`
}
