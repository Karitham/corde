package corde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Karitham/corde/internal/rest"
)

// returns the content-type
func toBodyMessage(w io.Writer, m Message) (string, error) {
	contentType := "application/json"

	payloadJSON := &bytes.Buffer{}
	err := json.NewEncoder(payloadJSON).Encode(m)
	if err != nil {
		return "", err
	}

	if len(m.Attachments) < 1 {
		payloadJSON.WriteTo(w)
		return contentType, nil
	}

	mw := multipart.NewWriter(w)
	defer mw.Close()

	contentType = mw.FormDataContentType()
	mw.WriteField("payload_json", payloadJSON.String())

	if err := writeAttachments(mw, m.Attachments); err != nil {
		return contentType, err
	}

	return contentType, nil
}

// CreateMessage creates a new message in a channel
//
// https://discord.com/developers/docs/resources/channel#create-message
func (m *Mux) CreateMessage(channelID Snowflake, data Message) (*Message, error) {
	body := &bytes.Buffer{}
	contentType, err := toBodyMessage(body, data)
	if err != nil {
		return nil, err
	}

	resp, err := m.Client.Do(
		rest.Req("/channels", channelID, "messages").
			AnyBody(body).Post(m.authorize, rest.ContentType(contentType)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		anyerr := map[string]any{}
		json.NewDecoder(resp.Body).Decode(&anyerr)
		return nil, fmt.Errorf("failed to create message: %+v", anyerr)
	}

	msg := &Message{}
	json.NewDecoder(resp.Body).Decode(msg)
	return msg, nil
}
