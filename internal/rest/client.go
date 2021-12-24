package rest

import (
	"encoding/json"
	"net/http"
)

func DoJson(c *http.Client, r *http.Request, v any) (*http.Response, error) {
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
