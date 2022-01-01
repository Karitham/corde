package corde

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	radix "github.com/akrennmair/go-radix"
)

type routes struct {
	command      *radix.Tree[Handler]
	autocomplete *radix.Tree[Handler]
	component    *radix.Tree[Handler]
}

// Mux is a discord gateway muxer, which handles the routing
type Mux struct {
	rMu        *sync.RWMutex
	routes     routes
	PublicKey  string // the hex public key provided by discord
	BasePath   string // base route path, default is "/"
	OnNotFound Handler
	Client     *http.Client
	AppID      Snowflake
	BotToken   string
}

// Lock the mux, to be able to mount or unmount routes
func (m *Mux) Lock() {
	m.rMu.Lock()
}

// Unlock the mux, so it can route again
func (m *Mux) Unlock() {
	m.rMu.Unlock()
}

func (m *Mux) Mount(typ InteractionType, route string, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()

	switch typ {
	case APPLICATION_COMMAND:
		m.routes.command.Insert(route, &handler)
	case APPLICATION_COMMAND_AUTOCOMPLETE:
		m.routes.autocomplete.Insert(route, &handler)
	case MESSAGE_COMPONENT:
		m.routes.component.Insert(route, &handler)
	}
}

// Button mounts a button route on the mux
func (m *Mux) Button(route string, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()
	m.routes.component.Insert(route, &handler)
}

// Autocomplete mounts an autocomplete route on the mux
func (m *Mux) Autocomplete(route string, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()
	m.routes.autocomplete.Insert(route, &handler)
}

// Command mounts a slash command route on the mux
func (m *Mux) Command(route string, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()
	m.routes.command.Insert(route, &handler)
}

// Route routes common parts along a pattern
func (m *Mux) Route(pattern string, fn func(m *Mux)) {
	if fn == nil {
		panic(fmt.Sprintf("corde: attempting to Route() a nil subrouter on %q", pattern))
	}

	r := NewMux(m.PublicKey, m.AppID, m.BotToken)
	fn(r)

	pattern = strings.TrimLeft(pattern, "/")
	for route, handler := range r.routes.command.ToMap() {
		m.routes.command.Insert(path.Join(pattern, route), handler)
	}
	for route, handler := range r.routes.autocomplete.ToMap() {
		m.routes.autocomplete.Insert(path.Join(pattern, route), handler)
	}
	for route, handler := range r.routes.component.ToMap() {
		m.routes.component.Insert(path.Join(pattern, route), handler)
	}
}

// NewMux returns a new mux for routing slash commands
//
// When you mount a command on the mux, it's prefix based routed,
// which means you can route to a button like `/list/next/456132153` having mounted `/list/next`
func NewMux(publicKey string, appID Snowflake, botToken string) *Mux {
	return &Mux{
		rMu: &sync.RWMutex{},
		routes: routes{
			command:      radix.New[Handler](),
			autocomplete: radix.New[Handler](),
			component:    radix.New[Handler](),
		},
		PublicKey: publicKey,
		BasePath:  "/",
		OnNotFound: func(_ ResponseWriter, i *Interaction) {
			log.Printf("No handler for registered command: %s\n", i.Data.Name)
		},
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		AppID:    appID,
		BotToken: botToken,
	}
}

// Handler handles incoming requests
type Handler func(ResponseWriter, *Interaction)

// ResponseWriter handles responding to interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type ResponseWriter interface {
	pong()
	Respond(i *InteractionRespData)
	DeferedRespond(i *InteractionRespData)
	Update(i *InteractionRespData)
	DeferedUpdate(i *InteractionRespData)
	Autocomplete(i *InteractionRespData)
}

// ListenAndServe starts the gateway listening to events
func (m *Mux) ListenAndServe(addr string) error {
	validator := Validate(m.PublicKey)
	r := http.NewServeMux()
	r.Handle(m.BasePath, validator(http.HandlerFunc(m.route)))

	return http.ListenAndServe(addr, r)
}

// route handles routing the requests
func (m *Mux) route(w http.ResponseWriter, r *http.Request) {
	i := &Interaction{}
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		log.Println("Errors unmarshalling json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m.routeReq(&Responder{w: w}, i)
}

// routeReq is a recursive implementation to route requests
func (m *Mux) routeReq(r ResponseWriter, i *Interaction) {
	m.rMu.RLock()
	defer m.rMu.RUnlock()
	switch i.Type {
	case PING:
		r.pong()
	case MESSAGE_COMPONENT:
		if _, h, ok := m.routes.component.LongestPrefix(i.Data.CustomID); ok {
			(*h)(r, i)
			return
		}
	case APPLICATION_COMMAND:
		if _, h, ok := m.routes.command.LongestPrefix(i.Data.Name); ok {
			(*h)(r, i)
			return
		}
		for optName := range i.Data.Options {
			nr := i.Data.Name + "/" + optName
			if _, h, ok := m.routes.command.LongestPrefix(nr); ok {
				i.Data.Name = nr
				(*h)(r, i)
				return
			}
		}
	case APPLICATION_COMMAND_AUTOCOMPLETE:
		log.Println("unimplemented autocomplete")
		r.Respond(NewResp().Ephemeral().Content("unimplemented autocomplete").B())
	}
	m.OnNotFound(r, i)
}

// reqOpts applies functions on an http request.
// useful for setting headers
func reqOpts(req *http.Request, h ...func(*http.Request)) {
	for _, option := range h {
		option(req)
	}
}

func (m *Mux) authorize(req *http.Request) {
	req.Header.Add("Authorization", "Bot "+m.BotToken)
}
