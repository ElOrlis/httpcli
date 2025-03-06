package qo

import (
	"io"
	"net/http"
	"strings"
)

func Body(b []byte) func(*http.Request) error {
	return func(r *http.Request) error {
		if b != nil {
			r.Body = io.NopCloser(strings.NewReader(string(b)))
		}
		return nil
	}
}

func Header(h http.Header) func(*http.Request) error {
	return func(r *http.Request) error {
		if h != nil {
			r.Header = h
		}
		return nil
	}
}

func QueryValues(q Query) func(*http.Request) error {
	return func(r *http.Request) error {
		q.compile(r)
		return nil
	}
}

func UrlParams(p Params) func(*http.Request) error {
	return func(r *http.Request) error {
		p.formatUrl(r)
		return nil
	}
}
