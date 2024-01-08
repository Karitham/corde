package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
)

type Request struct {
	root string
	path string
	body io.Reader
}

var API = "https://discord.com/api/v10"

func Req(paths ...any) *Request {
	r := &Request{
		root: API,
	}
	r.Append(paths...)

	return r
}

func (r *Request) URL() string {
	u, _ := url.Parse(r.root)
	u.Path = path.Join(u.Path, r.path)
	return u.String()
}

func (r *Request) JSONBody(v any) *Request {
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(v)

	r.body = b
	return r
}

func (r *Request) AnyBody(body io.Reader) *Request {
	r.body = body
	return r
}

func (r *Request) Append(v ...any) *Request {
	for _, val := range v {
		r.path = path.Join(r.path, fmt.Sprint(val))
	}
	return r
}

func (r *Request) Post(opts ...func(*http.Request)) *http.Request {
	return r.new(http.MethodPost, r.body, opts...)
}

func (r *Request) Put(opts ...func(*http.Request)) *http.Request {
	return r.new(http.MethodPut, r.body, opts...)
}

func (r *Request) Get(opts ...func(*http.Request)) *http.Request {
	return r.new(http.MethodGet, r.body, opts...)
}

func (r *Request) Delete(opts ...func(*http.Request)) *http.Request {
	return r.new(http.MethodDelete, r.body, opts...)
}

func (r *Request) Patch(opts ...func(*http.Request)) *http.Request {
	return r.new(http.MethodPatch, r.body, opts...)
}

func JSON(r *http.Request) {
	r.Header.Set("content-type", "application/json")
}

func Multipart(r *http.Request) {
	r.Header.Set("content-type", "multipart/form-data")
}

func Authorization(tok string) func(*http.Request) {
	return func(r *http.Request) {
		r.Header.Set("authorization", "Bot "+tok)
	}
}

func (r *Request) new(method string, body io.Reader, opts ...func(*http.Request)) *http.Request {
	req, err := http.NewRequest(method, r.URL(), body)
	if err != nil {
		return nil
	}

	for _, o := range opts {
		o(req)
	}

	if os.Getenv("CORDE_DUMP_HTTP_REQUEST") != "" {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(dump))
	}

	return req
}
