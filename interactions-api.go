package corde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Karitham/corde/internal/rest"
)

// returns the body and its content-type
func toBody(i *InteractionRespData) (*bytes.Buffer, string) {
	body := new(bytes.Buffer)
	contentType := "application/json"

	payloadJSON := &bytes.Buffer{}
	err := json.NewEncoder(payloadJSON).Encode(i)
	if err != nil {
		return nil, ""
	}

	if len(i.Attachments) < 1 {
		payloadJSON.WriteTo(body)
		return body, contentType
	}

	mw := multipart.NewWriter(body)
	defer mw.Close()

	contentType = mw.FormDataContentType()
	mw.WriteField("payload_json", payloadJSON.String())

	for i, f := range i.Attachments {
		if f.ID == 0 {
			f.ID = Snowflake(i)
		}

		ff, CFerr := mw.CreateFormFile(fmt.Sprintf("files[%d]", i), f.Filename)
		if CFerr != nil {
			return body, contentType
		}

		if _, CopyErr := io.Copy(ff, f.Body); CopyErr != nil {
			return body, contentType
		}
	}
	return body, contentType
}

// GetOriginalInteraction returns the original response to an Interaction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-original-interaction-response
func (m *Mux) GetOriginalInteraction(token string) (*InteractionRespData, error) {
	data := &InteractionRespData{}
	_, err := rest.DoJSON(m.Client, rest.Req("/webhooks", m.AppID, token, "messages/@original").Get(m.authorize), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// EditOriginalInteraction to edit your initial response to an Interaction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-original-interaction-response
func (m *Mux) EditOriginalInteraction(token string, data InteractionResponder) error {
	body, contentType := toBody(data.InteractionRespData())

	_, err := m.Client.Do(
		rest.Req("/webhooks", m.AppID, token, "messages/@original").
			AnyBody(body).Patch(m.authorize, rest.ContentType(contentType)),
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOriginalInteraction to delete your initial response to an Interaction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-original-interaction-response
func (m *Mux) DeleteOriginalInteraction(token string) error {
	_, err := m.Client.Do(
		rest.Req("/webhooks", m.AppID, token, "messages/@original").
			Delete(m.authorize),
	)
	if err != nil {
		return err
	}

	return nil
}

// FollowUpInteraction follows up a response to an Interaction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#followup-messages
func (m *Mux) FollowUpInteraction(token string, data InteractionResponder) error {
	body, contentType := toBody(data.InteractionRespData())

	_, err := m.Client.Do(
		rest.Req("/webhooks", m.AppID, token).
			AnyBody(body).Post(m.authorize, rest.ContentType(contentType)),
	)
	if err != nil {
		return err
	}
	return nil
}

// GetFollowUpInteraction returns the response to a FollowUpInteraction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-followup-message
func (m *Mux) GetFollowUpInteraction(token string, messageID Snowflake) (*InteractionRespData, error) {
	data := &InteractionRespData{}
	_, err := rest.DoJSON(m.Client, rest.Req("/webhooks", m.AppID, token, "messages", messageID).Get(m.authorize), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// EditFollowUpInteraction to edit a response to a FollowUpInteraction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-followup-message
func (m *Mux) EditFollowUpInteraction(token string, messageID Snowflake, data InteractionResponder) error {
	body, contentType := toBody(data.InteractionRespData())

	_, err := m.Client.Do(
		rest.Req("/webhooks", m.AppID, token, "messages", messageID).
			AnyBody(body).Patch(m.authorize, rest.ContentType(contentType)),
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFollowUpInteraction to delete a response to a FollowUpInteraction
//
// https://discord.com/developers/docs/interactions/receiving-and-responding#delete-followup-message
func (m *Mux) DeleteFollowUpInteraction(token string, messageID Snowflake) error {
	_, err := m.Client.Do(
		rest.Req("/webhooks", m.AppID, token, "messages", messageID).
			Delete(m.authorize),
	)
	if err != nil {
		return err
	}

	return nil
}
