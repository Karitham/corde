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

// Mux is a discord gateway muxer, which handles the routing
type Mux struct {
	rMu        *sync.RWMutex
	routes     *radix.Tree[Handlers]
	PublicKey  string // the hex public key provided by discord
	BasePath   string // base route path, default is "/"
	OnNotFound func(ResponseWriter, *InteractionRequest[components.JsonRaw])
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
func (m *Mux) Mount(typ components.InnerInteractionType, route string, handler any) {
	m.rMu.Lock()
	defer m.rMu.Unlock()

	if r, ok := m.routes.Get(route); ok {
		(*r)[typ] = handler
	} else {
		m.routes.Insert(route, &Handlers{typ: handler})
	}
}

// Button mounts a button route on the mux
func (m *Mux) Button(route string, handler func(ResponseWriter, *InteractionRequest[components.ButtonInteractionData])) {
	m.Mount(components.ButtonInteraction, route, handler)
}

// Autocomplete mounts an autocomplete route on the mux
func (m *Mux) Autocomplete(route string, handler func(ResponseWriter, *InteractionRequest[components.AutocompleteInteractionData])) {
	m.Mount(components.AutocompleteInteraction, route, handler)
}

// Command mounts a slash command route on the mux
func (m *Mux) Command(route string, handler func(ResponseWriter, *InteractionRequest[components.SlashCommandInteractionData])) {
	m.Mount(components.SlashCommandInteraction, route, handler)
}

// Route routes common parts along a pattern
func (m *Mux) Route(pattern string, fn func(m *Mux)) {
	if fn == nil {
		panic(fmt.Sprintf("corde: attempting to Route() a nil subrouter on %q", pattern))
	}

	r := NewMux(m.PublicKey, m.AppID, m.BotToken)
	fn(r)

	pattern = strings.TrimLeft(pattern, "/")
	for route, handler := range r.routes.ToMap() {
		m.routes.Insert(path.Join(pattern, route), handler)
	}
}

// NewMux returns a new mux for routing slash commands
//
// When you mount a command on the mux, it's prefix based routed,
// which means you can route to a button like `/list/next/456132153` having mounted `/list/next`
func NewMux(publicKey string, appID snowflake.Snowflake, botToken string) *Mux {
	m := &Mux{
		rMu:       &sync.RWMutex{},
		routes:    radix.New[Handlers](),
		PublicKey: publicKey,
		BasePath:  "/",
		OnNotFound: func(_ ResponseWriter, i *InteractionRequest[components.JsonRaw]) {
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

// Handlers handles incoming requests
type Handlers map[components.InnerInteractionType]any

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
type InteractionRequest[T components.InteractionDataConstraint] struct {
	components.Interaction[T]
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
	i := &InteractionRequest[components.JsonRaw]{
		Context: r.Context(),
	}

	if err := json.NewDecoder(r.Body).Decode(&i.Data); err != nil {
		log.Println("Errors unmarshalling json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m.routeReq(&Responder{w: w}, i)
}

// routeReq is a recursive implementation to route requests
func (m *Mux) routeReq(r ResponseWriter, i *InteractionRequest[components.JsonRaw]) {
	m.rMu.RLock()
	defer m.rMu.RUnlock()
	if i.Type == components.INTERACTION_TYPE_PING {
		r.Pong()
		return
	}

	var err error
	if _, handler, ok := m.routes.LongestPrefix(i.Route); ok {
		switch i.InnerInteractionType {
		// Component
		case components.ButtonInteraction:
			err = routeRequest[components.ButtonInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.ModalInteraction:
			err = routeRequest[components.ModalInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.SelectInteraction:
			err = routeRequest[components.SelectInteractionData](*handler, i.InnerInteractionType, r, i)
		// Autocomplete
		case components.AutocompleteInteraction:
			err = routeRequest[components.AutocompleteInteractionData](*handler, i.InnerInteractionType, r, i)

		// Slash
		case components.SlashCommandInteraction:
			err = routeRequest[components.SlashCommandInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.MessageCommandInteraction:
			err = routeRequest[components.MessageCommandInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.UserCommandInteraction:
			err = routeRequest[components.UserCommandInteractionData](*handler, i.InnerInteractionType, r, i)

		}
	}
	if err != nil {
		m.OnNotFound(r, i)
	}
}

// authorize adds the Authorization header to the request
func (m *Mux) authorize(req *http.Request) {
	req.Header.Add("Authorization", "Bot "+m.BotToken)
}

// Finds the handler for the route
func routeRequest[IntReqData components.InteractionDataConstraint](
	routes map[components.InnerInteractionType]any,
	it components.InnerInteractionType,
	r ResponseWriter,
	rawI *InteractionRequest[components.JsonRaw],
) error {
	if h, ok := routes[it].(func(ResponseWriter, *InteractionRequest[IntReqData])); ok {
		var intValues components.Interaction[IntReqData]

		json.Unmarshal(rawI.Data, &intValues)
		h(r, &InteractionRequest[IntReqData]{Context: rawI.Context, Interaction: intValues})
		return nil
	}

	return fmt.Errorf("No handler for interaction type: %d", it)
}
