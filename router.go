package corde

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type InteractionCommand struct {
	Type  InteractionType `json:"type"`
	Route string          `json:"name"`
}

func SlashCommand(route string) InteractionCommand {
	return InteractionCommand{Type: APPLICATION_COMMAND, Route: route}
}

func ButtonInteraction(customID string) InteractionCommand {
	return InteractionCommand{Type: MESSAGE_COMPONENT, Route: customID}
}

// Mux is a discord gateway muxer, which handles the routing
type Mux struct {
	rMu        *sync.RWMutex
	routes     map[InteractionCommand]Handler
	PublicKey  string // the hex public key provided by discord
	BasePath   string // base route path, default is "/"
	OnNotFound Handler
	Client     *http.Client
	AppID      Snowflake
	BotToken   string
}

// Routes return the discord routes mounted on the mux
// DO NOT EDIT THOSE, IN RISK OF HAVING ROUTING ISSUES
func (m *Mux) Routes() map[InteractionCommand]Handler {
	return m.routes
}

// Lock the mux, to be able to mount or unmount routes
func (m *Mux) Lock() {
	m.rMu.Lock()
}

// Unlock the mux, so it can route again
func (m *Mux) Unlock() {
	m.rMu.Unlock()
}

func (m *Mux) SetRoute(command InteractionCommand, handler Handler) {
	m.rMu.Lock()
	defer m.rMu.Unlock()
	m.routes[command] = handler
}

// NewMux returns a new mux for routing slash commands
func NewMux(publicKey string, appID Snowflake, botToken string) *Mux {
	return &Mux{
		rMu:       &sync.RWMutex{},
		routes:    make(map[InteractionCommand]Handler),
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
	WithSource(i *InteractionRespData)
	DeferedWithSource(i *InteractionRespData)
	UpdateMessage(i *InteractionRespData)
	DeferedUpdateMessage(i *InteractionRespData)
	AutocompleteResult(i *InteractionRespData)
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
		if h, ok := m.routes[InteractionCommand{Type: i.Type, Route: i.Data.CustomID}]; ok {
			h(r, i)
			return
		}

		for optName := range i.Data.Options {
			nr := InteractionCommand{Type: i.Type, Route: i.Data.Name + "/" + i.Data.CustomID}

			if handler, ok := m.routes[nr]; ok {
				i.Data.Name += "/" + optName
				handler(r, i)
				return
			}
		}

		m.OnNotFound(r, i)
	case APPLICATION_COMMAND:
		fallthrough
	case APPLICATION_COMMAND_AUTOCOMPLETE:
		if h, ok := m.routes[InteractionCommand{Type: i.Type, Route: i.Data.Name}]; ok {
			h(r, i)
			return
		}

		for optName := range i.Data.Options {
			nr := InteractionCommand{Type: i.Type, Route: i.Data.Name + "/" + optName}

			if handler, ok := m.routes[nr]; ok {
				i.Data.Name += "/" + optName
				handler(r, i)
				return
			}
		}

		m.OnNotFound(r, i)
	}
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
