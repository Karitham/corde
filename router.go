package corde

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type InteractionCommand struct {
	Type InteractionType `json:"type"`
	Name string          `json:"name"`
}

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

func (m *Mux) Routes() map[InteractionCommand]Handler {
	return m.routes
}

func (m *Mux) AddRoute(command InteractionCommand, handler Handler) {
	m.rMu.Lock()
	m.routes[command] = handler
	m.rMu.Unlock()
}

func NewMux(publicKey string, appID Snowflake, botToken string) *Mux {
	return &Mux{
		rMu:       &sync.RWMutex{},
		routes:    make(map[InteractionCommand]Handler),
		PublicKey: publicKey,
		BasePath:  "/",
		OnNotFound: func(_ ResponseWriter, i *Interaction) {
			log.Println("no handler for", i.Type, i.Data.Name)
		},
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		AppID:    appID,
		BotToken: botToken,
	}
}

type Handler func(ResponseWriter, *Interaction)

type ResponseWriter interface {
	Pong()
	ChannelMessageWithSource(i InteractionResponseData)
}

func (m *Mux) ListenAndServe(addr string) error {
	validator := Validate(m.PublicKey)
	r := http.NewServeMux()
	r.Handle(m.BasePath, validator(http.HandlerFunc(m.Route)))

	return http.ListenAndServe(addr, r)
}

func (m *Mux) Route(w http.ResponseWriter, r *http.Request) {
	i := &Interaction{}
	if err := json.NewDecoder(r.Body).Decode(i); err != nil {
		log.Println("errors unmarshalling json:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := &Responder{w: w}
	switch i.Type {
	case PING:
		rsp.Pong()
	case APPLICATION_COMMAND:
		fallthrough
	case APPLICATION_COMMAND_AUTOCOMPLETE:
		m.rMu.RLock()
		defer m.rMu.RUnlock()
		if h, ok := m.routes[InteractionCommand{Type: i.Type, Name: i.Data.Name}]; ok {
			h(rsp, i)
			return
		}

		m.OnNotFound(rsp, i)
	}
}

func (m *Mux) applyOpt(req *http.Request, h ...func(*http.Request)) {
	for _, option := range h {
		option(req)
	}
}

func (m *Mux) authorize(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", m.BotToken))
}

func (m *Mux) contentTypeJSON(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
}
