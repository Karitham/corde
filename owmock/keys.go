package owmock

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
)

// GenerateKeys generates a new keypair of hex-encoded keys
// They shouldn't change between runs, but there's a generator just in case
//
// PubK:
// 2f8c6129d816cf51c374bc7f08c3e63ed156cf78aefb4a6550d97b87997977ee
//
// PrivK:
// 31323334353637383930313233343536373839303132333435363738393031322f8c6129d816cf51c374bc7f08c3e63ed156cf78aefb4a6550d97b87997977ee
func GenerateKeys() (pub string, priv string) {
	pubK, privK, _ := ed25519.GenerateKey(bytes.NewBufferString("12345678901234567890123456789012"))
	return hex.EncodeToString(pubK), hex.EncodeToString(privK)
}
