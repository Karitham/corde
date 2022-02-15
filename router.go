package corde

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/internal/rest"
	"github.com/Karitham/corde/snowflake"
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
	AppID      snowflake.Snowflake
	BotToken   string

	handler http.Handler
}

// Lock the mux, to be able to mount or unmount routes
func (m *Mux) Lock() {
	m.rMu.Lock()
}

// Unlock the mux, so it can route again
func (m *Mux) Unlock() {
	m.rMu.Unlock()
}

// Mount is for mounting a Handler on the Mux
func (m *Mux) Mount(typ components.InteractionType, route string, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()

	switch typ {
	case components.INTERACTION_TYPE_APPLICATION_COMMAND:
		m.routes.command.Insert(route, &handler)
	case components.INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE:
		m.routes.autocomplete.Insert(route, &handler)
	case components.INTERACTION_TYPE_MESSAGE_COMPONENT:
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
func NewMux(publicKey string, appID snowflake.Snowflake, botToken string) *Mux {
	m := &Mux{
		rMu: &sync.RWMutex{},
		routes: routes{
			command:      radix.New[Handler](),
			autocomplete: radix.New[Handler](),
			component:    radix.New[Handler](),
		},
		PublicKey: publicKey,
		BasePath:  "/",
		OnNotFound: func(_ ResponseWriter, i *InteractionRequest) {
			log.Printf("No handler for registered command: %s\n", i.Route)
		},
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		AppID:    appID,
		BotToken: botToken,
	}

	m.handler = rest.Verify(publicKey)(http.HandlerFunc(m.route))
	return m
}

// Handler handles incoming requests
type Handler func(ResponseWriter, *InteractionRequest)

// ResponseWriter handles responding to interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type ResponseWriter interface {
	Pong()
	Respond(InteractionResponder)
	DeferedRespond(InteractionResponder)
	Update(InteractionResponder)
	DeferedUpdate(InteractionResponder)
	Autocomplete(InteractionResponder)
}

// InteractionRequest is an incoming request Interaction
type InteractionRequest struct {
	components.Interaction
	Context context.Context
}

// ListenAndServe starts the gateway listening to events
func (m *Mux) ListenAndServe(addr string) error {
	r := http.NewServeMux()
	r.Handle(m.BasePath, m)

	return http.ListenAndServe(addr, r)
}

// ServeHTTP will serve HTTP requests with discord public key validation
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r)
}

// route handles routing the requests
func (m *Mux) route(w http.ResponseWriter, r *http.Request) {
	i := &InteractionRequest{
		Context: r.Context(),
	}
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		log.Println("Errors unmarshalling json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m.routeReq(&Responder{w: w}, i)
}

// routeReq is a recursive implementation to route requests
func (m *Mux) routeReq(r ResponseWriter, i *InteractionRequest) {
	m.rMu.RLock()
	defer m.rMu.RUnlock()
	switch i.Type {
	case components.INTERACTION_TYPE_PING:
		r.Pong()
		return
	case components.INTERACTION_TYPE_MESSAGE_COMPONENT:
		data, _ := components.GetInteractionData[components.SelectInteractionData](i.Interaction)
		i.Route = data.CustomID
		if _, h, ok := m.routes.component.LongestPrefix(i.Route); ok {
			(*h)(r, i)
			return
		}
	case components.INTERACTION_TYPE_APPLICATION_COMMAND:
		// for menu & app commands, which can have spaces
		data, _ := components.GetInteractionData[components.SlashInteractionData](i.Interaction)
		// not slash command
		if data.Type != 1 {
			data.Name = path.Join(strings.Fields(data.Name)...)
		}

		group := data.Options[components.RouteInteractionSubcommandGroup]
		cmd := data.Options[components.RouteInteractionSubcommand]
		i.Route = path.Join(data.Name, group.String(), cmd.String())

		if _, h, ok := m.routes.command.LongestPrefix(i.Route); ok {
			(*h)(r, i)
			return
		}
	case components.INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE:
		data, _ := components.GetInteractionData[components.AutocompleteInteractionData](i.Interaction)

		group := data.Options[components.RouteInteractionSubcommandGroup]
		cmd := data.Options[components.RouteInteractionSubcommand]
		focused := data.Options[components.RouteInteractionFocused]
		i.Route = path.Join(data.Name, group.String(), cmd.String(), focused.String())

		if _, h, ok := m.routes.autocomplete.LongestPrefix(i.Route); ok {
			(*h)(r, i)
			return
		}
	}
	m.OnNotFound(r, i)
}

// authorize adds the Authorization header to the request
func (m *Mux) authorize(req *http.Request) {
	req.Header.Add("Authorization", "Bot "+m.BotToken)
}
