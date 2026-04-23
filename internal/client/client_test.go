package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/shoenig/test"
	"github.com/shoenig/test/must"
)

func TestHttpClient_CreateRequest(t *testing.T) {
	t.Parallel()

	t.Skip("todo")
}

func TestHttpClient_Get(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		test.Eq(t, http.MethodGet, r.Method)
		test.Eq(t, "/test", r.URL.Path)
		test.Eq(t, "1", r.URL.Query().Get("a"))
		test.Eq(t, "2", r.URL.Query().Get("b"))
	}))
	defer srv.Close()

	client := NewClient(srv.URL)
	resp, err := client.Get(t.Context(), "/test", url.Values{"a": {"1"}, "b": {"2"}})
	must.NoError(t, err)
	must.NotNil(t, resp)
}
