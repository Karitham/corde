package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func CodeBetween(resp *http.Response, BigOrEq int, LessOrEq int) error {
	if BigOrEq != 0 && resp.StatusCode < BigOrEq {
		return body(resp, fmt.Errorf("received status code %d", resp.StatusCode))
	}
	if LessOrEq != 0 && resp.StatusCode > LessOrEq {
		return body(resp, fmt.Errorf("received status code %d", resp.StatusCode))
	}

	return nil
}

func ExpectCode(resp *http.Response, expect int) error {
	if resp.StatusCode != expect {
		return body(resp, fmt.Errorf("received status code %d, expected %d", resp.StatusCode, expect))
	}

	return nil
}

func body(resp *http.Response, err error) error {
	if resp.Body != nil {
		b := &bytes.Buffer{}
		b.ReadFrom(resp.Body)
		resp.Body = io.NopCloser(b)
		return fmt.Errorf("error: %w, body: %s", err, b.String())
	}

	return err
}
