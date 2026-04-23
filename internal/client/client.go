package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nheggoe/gober/internal/util"
)

type Client struct {
	*http.Client
	baseURL string
}

func NewClient(baseUrl string, options ...Option) Client {
	client := Client{
		Client:  http.DefaultClient,
		baseURL: baseUrl,
	}
	for _, opt := range options {
		opt(&client)
	}
	return client
}

// Get sends an HTTP GET request to the specified path with optional query parameters and returns the HTTP response.
func (c *Client) Get(ctx context.Context, path string, parameters url.Values) (_ *http.Response, err error) {
	defer util.WrapError(&err, "Client.Get")

	req, err := c.newRequest(ctx, http.MethodGet, path, parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.Do(req)
}

// newRequest constructs a new HTTP request with the given context, method, path, and query parameters.
func (c *Client) newRequest(ctx context.Context, method string, path string, parameters url.Values) (_ *http.Request, err error) {
	defer util.WrapError(&err, "Client.NewRequest")

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %w", err)
	}

	rel, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url path: %w", err)
	}

	u = u.ResolveReference(rel)
	if parameters != nil {
		u.RawQuery = parameters.Encode()
	}

	return http.NewRequestWithContext(ctx, method, u.String(), nil)
}
