package owmock

import (
	"testing"

	"github.com/Karitham/corde"
)

// ResponseWriterMock mocks corde's ResponseWriter interface
type ResponseWriterMock struct {
	RespondHook        func(corde.InteractionResponseDataBuilder)
	DeferedRespondHook func(corde.InteractionResponseDataBuilder)
	UpdateHook         func(corde.InteractionResponseDataBuilder)
	DeferedUpdateHook  func(corde.InteractionResponseDataBuilder)
	AutocompleteHook   func(corde.InteractionResponseDataBuilder)

	T *testing.T
}

// NewRWMock returns a new ResponseWriterMock with the given testing.T
func NewRWMock(t *testing.T) ResponseWriterMock {
	return ResponseWriterMock{T: t}
}

// Pong implements ResponseWriter interface
func (r ResponseWriterMock) Pong() {}

// Response implements ResponseWriter interface
func (r ResponseWriterMock) Respond(i corde.InteractionResponseDataBuilder) {
	if r.RespondHook != nil {
		r.RespondHook(i)
		return
	}

	r.T.Error("unexpected respond hook called")
}

// DeferedRespond implements ResponseWriter interface
func (r ResponseWriterMock) DeferedRespond(i corde.InteractionResponseDataBuilder) {
	if r.DeferedRespondHook != nil {
		r.DeferedRespondHook(i)
		return
	}

	r.T.Error("unexpected defered respond hook called")
}

// Update implements ResponseWriter interface
func (r ResponseWriterMock) Update(i corde.InteractionResponseDataBuilder) {
	if r.UpdateHook != nil {
		r.UpdateHook(i)
		return
	}

	r.T.Error("unexpected update hook called")
}

// DeferedUpdate implements ResponseWriter interface
func (r ResponseWriterMock) DeferedUpdate(i corde.InteractionResponseDataBuilder) {
	if r.DeferedUpdateHook != nil {
		r.DeferedUpdateHook(i)
		return
	}

	r.T.Error("unexpected defered update hook called")
}

// Autocomplete implements ResponseWriter interface
func (r ResponseWriterMock) Autocomplete(i corde.InteractionResponseDataBuilder) {
	if r.AutocompleteHook != nil {
		r.AutocompleteHook(i)
		return
	}

	r.T.Error("unexpected autocomplete hook called")
}
