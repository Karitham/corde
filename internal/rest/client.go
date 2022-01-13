package rest

import (
	"encoding/json"
	"net/http"
)

// DoJSON executes a request and decodes the response into the given interface
// It already calls `Close()` on the body
func DoJSON(c *http.Client, r *http.Request, v any) (*http.Response, error) {
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, err
	}

	return resp, nil
}

func ContentType(contentType string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("content-type", contentType)
	}
}
