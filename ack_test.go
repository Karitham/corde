package corde

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, bad, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	timestamp := fmt.Sprint(time.Now().Unix())

	tt := []struct {
		Name   string
		Body   string
		Key    ed25519.PrivateKey
		Status int
	}{
		{
			Name:   "Correct",
			Key:    priv,
			Status: http.StatusOK,
		},
		{
			Name:   "Correct With Body",
			Body:   "foo",
			Key:    priv,
			Status: http.StatusOK,
		},
		{
			Name:   "Bad Key",
			Key:    bad,
			Status: http.StatusUnauthorized,
		},
		{
			Name:   "Bad Key With Body",
			Body:   "bar",
			Key:    bad,
			Status: http.StatusUnauthorized,
		},
	}

	s := httptest.NewServer(Validate(hex.EncodeToString(pub))(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(nil)
	})))
	defer s.Close()

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, s.URL, strings.NewReader(tc.Body))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			req.Header.Set("X-Signature-Timestamp", timestamp)
			sig := ed25519.Sign(tc.Key, []byte(timestamp+tc.Body))
			req.Header.Set("X-Signature-Ed25519", hex.EncodeToString(sig))

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if resp.StatusCode != tc.Status {
				t.Logf("incorrect response, expected status %d but got %d\n", tc.Status, resp.StatusCode)
				t.FailNow()
			}
		})
	}
}
