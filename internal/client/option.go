package client

import (
	"net/http"
	"time"
)

type Option func(client *Client)

// Timeout

func WithTimeout(duration time.Duration) Option {
	return func(client *Client) {
		client.Timeout = duration
	}
}

// Headers

type Headers = map[string][]string

func WithHeaders(headers Headers) Option {
	return func(client *Client) {
		client.Transport = &headerRoundTripper{client.Transport, headers}
	}
}

type headerRoundTripper struct {
	base    http.RoundTripper
	headers Headers
}

func (rt *headerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, vs := range rt.headers {
		for _, v := range vs {
			r.Header.Add(k, v)
		}
	}
	return rt.base.RoundTrip(r)
}
