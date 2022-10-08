package corde

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

// ResponseWriter handles responding to interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type ResponseWriter interface {
	Ack()
	Respond(InteractionResponder)
	DeferedRespond()
	Update(InteractionResponder)
	DeferedUpdate()
	Autocomplete(InteractionResponder)
	Modal(Modal)
}

// Request is an incoming request Interaction
type Request[T InteractionDataConstraint] struct {
	Interaction[T]
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
	i := &Request[JsonRaw]{
		Context: r.Context(),
	}

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Println("Errors unmarshalling json: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data PartialRoutingType
	i.Data.UnmarshalTo(&data)

	// build route
	group := data.Options[RouteInteractionSubcommandGroup]
	cmd := data.Options[RouteInteractionSubcommand]
	focused := data.Options[RouteInteractionFocused]
	data.Name = path.Join(strings.Fields(data.Name)...)
	i.Route = path.Join(data.Name, data.CustomID, group.String(), cmd.String(), focused.String())

	// Find inner type
	switch i.Type {
	case INTERACTION_TYPE_PING:
		i.Type = INTERACTION_TYPE_PING
	case INTERACTION_TYPE_APPLICATION_COMMAND_AUTOCOMPLETE:
		i.InnerInteractionType = AutocompleteInteraction
	case INTERACTION_TYPE_APPLICATION_COMMAND:
		i.Type = INTERACTION_TYPE_APPLICATION_COMMAND
		switch data.Type {
		case 1:
			i.InnerInteractionType = SlashCommandInteraction
		case 2:
			i.InnerInteractionType = UserCommandInteraction
		case 3:
			i.InnerInteractionType = MessageCommandInteraction
		default:
			i.InnerInteractionType = SlashCommandInteraction
		}
	case INTERACTION_TYPE_MESSAGE_COMPONENT:
		i.Type = INTERACTION_TYPE_MESSAGE_COMPONENT
		switch data.ComponentType {
		case 1:
			i.InnerInteractionType = ActionRowInteraction
		case 2:
			i.InnerInteractionType = ButtonInteraction
		case 3:
			i.InnerInteractionType = SelectMenuInteraction
		case 4:
			i.InnerInteractionType = TextInputInteraction
		}
	case INTERACTION_TYPE_MODAL:
		i.Type = INTERACTION_TYPE_MODAL
		i.InnerInteractionType = ModalInteraction
	}

	m.routeReq(&Responder{w: w}, i)
}

// routeReq is a recursive implementation to route requests
func (m *Mux) routeReq(r ResponseWriter, i *Request[JsonRaw]) {
	m.rMu.RLock()
	defer m.rMu.RUnlock()
	if i.Type == INTERACTION_TYPE_PING {
		r.Ack()
		return
	}

	var err error
	if _, handler, ok := m.routes.LongestPrefix(i.Route); ok {
		switch i.InnerInteractionType {
		// Component
		case ButtonInteraction: // works & tested
			err = routeRequest[ButtonInteractionData](*handler, i.InnerInteractionType, r, i)
		case SelectMenuInteraction:
			err = routeRequest[ModalInteractionData](*handler, i.InnerInteractionType, r, i)
		case ActionRowInteraction:
			err = routeRequest[SelectInteractionData](*handler, i.InnerInteractionType, r, i)
		case TextInputInteraction:
			err = routeRequest[TextInputInteractionData](*handler, i.InnerInteractionType, r, i)

		// Autocomplete
		case AutocompleteInteraction:
			err = routeRequest[AutocompleteInteractionData](*handler, i.InnerInteractionType, r, i)

		// Slash
		case SlashCommandInteraction:
			err = routeRequest[SlashCommandInteractionData](*handler, i.InnerInteractionType, r, i)
		case MessageCommandInteraction:
			err = routeRequest[MessageCommandInteractionData](*handler, i.InnerInteractionType, r, i)
		case UserCommandInteraction:
			err = routeRequest[UserCommandInteractionData](*handler, i.InnerInteractionType, r, i)

		// Modal
		case ModalInteraction:
			err = routeRequest[ModalInteractionData](*handler, i.InnerInteractionType, r, i)
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
func routeRequest[IntReqData InteractionDataConstraint](
	routes map[InnerInteractionType]any,
	it InnerInteractionType,
	r ResponseWriter,
	rawI *Request[JsonRaw],
) error {
	if h, ok := routes[it].(func(ResponseWriter, *Request[IntReqData])); ok {
		var intValues Interaction[IntReqData]
		v, _ := json.Marshal(rawI) // Better than mapping by hand, but I hate it
		if err := json.Unmarshal(v, &intValues); err != nil {
			return err
		}
		intValues.Route = rawI.Route

		h(r, &Request[IntReqData]{Context: rawI.Context, Interaction: intValues})
		return nil
	}

	return fmt.Errorf("no handler for interaction type: %d", it)
}
