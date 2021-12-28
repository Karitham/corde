package corde

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"net/http"
)

// Validate is a middleware to validate Interaction payloads
func Validate(publicKey string) func(http.Handler) http.Handler {
	pk, err := hex.DecodeString(publicKey)
	if err != nil {
		panic("invalid public key")
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sig, _ := hex.DecodeString(r.Header.Get("X-Signature-Ed25519"))
			timestamp := r.Header.Get("X-Signature-Timestamp")

			b, oldBody := &bytes.Buffer{}, r.Body
			if _, err := b.ReadFrom(r.Body); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			oldBody.Close()

			r.Body = io.NopCloser(b)
			if !ed25519.Verify(pk, append([]byte(timestamp), b.Bytes()...), sig) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
