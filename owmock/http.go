package owmock

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/matryer/is"
)

// req makes a defined request to the given uri
func req(c Doer, method string, url string, buf *bytes.Buffer, privK ed25519.PrivateKey) (json.RawMessage, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	req, _ := http.NewRequest(method, url, buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Signature-Timestamp", timestamp)
	sig := ed25519.Sign(privK, []byte(timestamp+buf.String()))
	req.Header.Set("X-Signature-Ed25519", hex.EncodeToString(sig))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.ContentLength == 0 {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code: " + strconv.Itoa(resp.StatusCode))
	}

	respBody := json.RawMessage{}
	if err = json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	log.Println(string(respBody))
	return respBody, nil
}

// Doer makes requests
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Requester makes requests to the mock endpoint
type Requester struct {
	Client     Doer
	URL        string
	PrivateKey ed25519.PrivateKey
}

// New returns a new mock requester
func New(endpointURL string) *Requester {
	_, priv := GenerateKeys()
	v, _ := hex.DecodeString(priv)

	return &Requester{
		Client:     &http.Client{Timeout: 10 * time.Second},
		URL:        endpointURL,
		PrivateKey: v,
	}
}

// NewWithClient returns a new mock requester with a default client
func NewWithClient(endpointURL string, c Doer) *Requester {
	_, priv := GenerateKeys()
	v, err := hex.DecodeString(priv)
	if err != nil {
		panic(err)
	}
	return &Requester{
		Client:     c,
		PrivateKey: v,
		URL:        endpointURL,
	}
}

// PostJSON makes a POST request to the endpoint with the given body marshalled
func (r *Requester) PostJSON(body any) (json.RawMessage, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}
	return req(r.Client, http.MethodPost, r.URL, buf, r.PrivateKey)
}

// PostJSON makes a POST request to the endpoint with the given body
func (r *Requester) Post(body string) (json.RawMessage, error) {
	buf := bytes.NewBufferString(body)
	return req(r.Client, http.MethodPost, r.URL, buf, r.PrivateKey)
}

// PostExpect posts a payload and expects a response with the given body
func (r *Requester) PostExpect(t is.T, body any, expectV any) error {
	is := is.New(t)
	resp, err := r.PostJSON(body)
	is.NoErr(err)

	typ := reflect.TypeOf(expectV)
	respV := reflect.New(typ).Interface()

	err = json.Unmarshal(resp, respV)
	is.NoErr(err)

	is.Equal(respV, expectV)
	return nil
}
