package corde

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/Karitham/corde/components"
)

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

// Request is an incoming request Interaction
type Request[T components.InteractionDataConstraint] struct {
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
	i := &Request[components.JsonRaw]{
		Context: r.Context(),
	}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Println("Errors unmarshalling json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data components.PartialRoutingType
	i.Data.UnmarshalTo(&data)

	// build route
	group := data.Options[components.RouteInteractionSubcommandGroup]
	cmd := data.Options[components.RouteInteractionSubcommand]
	focused := data.Options[components.RouteInteractionFocused]
	data.Name = path.Join(strings.Fields(data.Name)...)
	i.Route = path.Join(data.Name, data.CustomID, group.String(), cmd.String(), focused.String())

	// Find inner type
	switch i.Type {
	case components.INTERACTION_TYPE_PING:
		i.Type = components.INTERACTION_TYPE_PING
	case components.INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE:
		i.InnerInteractionType = components.AutocompleteInteraction
	case components.INTERACTION_TYPE_APPLICATION_COMMAND:
		i.Type = components.INTERACTION_TYPE_APPLICATION_COMMAND
		switch data.Type {
		case 1:
			i.InnerInteractionType = components.SlashCommandInteraction
		case 2:
			i.InnerInteractionType = components.UserCommandInteraction
		case 3:
			i.InnerInteractionType = components.MessageCommandInteraction
		default:
			i.InnerInteractionType = components.SlashCommandInteraction
		}
	case components.INTERACTION_TYPE_MESSAGE_COMPONENT:
		i.Type = components.INTERACTION_TYPE_MESSAGE_COMPONENT
		switch data.ComponentType {
		case 1:
			i.InnerInteractionType = components.ActionRowInteraction
		case 2:
			i.InnerInteractionType = components.ButtonInteraction
		case 3:
			i.InnerInteractionType = components.SelectMenuInteraction
		case 4:
			i.InnerInteractionType = components.TextInputInteraction
		}
	}

	m.routeReq(&Responder{w: w}, i)
}

// routeReq is a recursive implementation to route requests
func (m *Mux) routeReq(r ResponseWriter, i *Request[components.JsonRaw]) {
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
		case components.ButtonInteraction: // works & tested
			err = routeRequest[components.ButtonInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.SelectMenuInteraction:
			err = routeRequest[components.ModalInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.ActionRowInteraction:
			err = routeRequest[components.SelectInteractionData](*handler, i.InnerInteractionType, r, i)
		case components.TextInputInteraction:
			err = routeRequest[components.TextInputInteractionData](*handler, i.InnerInteractionType, r, i)

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
	rawI *Request[components.JsonRaw],
) error {
	if h, ok := routes[it].(func(ResponseWriter, *Request[IntReqData])); ok {
		var intValues components.Interaction[IntReqData]
		v, _ := json.Marshal(rawI) // Better than mapping by hand, but I hate it
		if err := json.Unmarshal(v, &intValues); err != nil {
			return err
		}

		h(r, &Request[IntReqData]{Context: rawI.Context, Interaction: intValues})
		return nil
	}

	return fmt.Errorf("No handler for interaction type: %d", it)
}
