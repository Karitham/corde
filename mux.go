package corde

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/Karitham/corde/internal/rest"
	"github.com/akrennmair/go-radix"
)

// Mux is a discord gateway muxer, which handles the routing
type Mux struct {
	rMu        *sync.RWMutex
	routes     *radix.Tree[Handlers]
	PublicKey  string // the hex public key provided by discord
	BasePath   string // base route path, default is "/"
	OnNotFound func(ResponseWriter, *Request[JsonRaw])
	Client     *http.Client
	AppID      Snowflake
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

// NewMux returns a new mux for routing slash commands
//
// When you mount a command on the mux, it's prefix based routed,
// which means you can route to a button like `/list/next/456132153` having mounted `/list/next`
func NewMux(publicKey string, appID Snowflake, botToken string) *Mux {
	m := &Mux{
		rMu:       &sync.RWMutex{},
		routes:    radix.New[Handlers](),
		PublicKey: publicKey,
		BasePath:  "/",
		OnNotFound: func(_ ResponseWriter, i *Request[JsonRaw]) {
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
type Handlers map[InnerInteractionType]any

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

// Mount is for mounting a Handler on the Mux
func (m *Mux) Mount(typ InnerInteractionType, route string, handler any) {
	m.rMu.Lock()
	defer m.rMu.Unlock()

	if r, ok := m.routes.Get(route); ok {
		(*r)[typ] = handler
		return
	}

	m.routes.Insert(route, &Handlers{typ: handler})
}
