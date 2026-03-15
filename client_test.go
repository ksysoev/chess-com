package chess

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := New()

	require.NotNil(t, c, "New() should not return nil")
	assert.Equal(t, defaultBaseURL, c.baseURL)
	assert.Equal(t, defaultUserAgent, c.userAgent)
	assert.NotNil(t, c.httpClient)
}

func TestWithBaseURL(t *testing.T) {
	t.Parallel()

	c := New(WithBaseURL("https://example.com"))

	assert.Equal(t, "https://example.com", c.baseURL)
}

func TestWithUserAgent(t *testing.T) {
	t.Parallel()

	c := New(WithUserAgent("test-agent/1.0"))

	assert.Equal(t, "test-agent/1.0", c.userAgent)
}

func TestWithTimeout(t *testing.T) {
	t.Parallel()

	c := New(WithTimeout(5 * time.Second))

	assert.Equal(t, 5*time.Second, c.httpClient.Timeout)
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	custom := &http.Client{Timeout: 99 * time.Second}
	c := New(WithHTTPClient(custom))

	assert.Equal(t, custom, c.httpClient)
}

func TestGet_Success(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("ETag", `"abc123"`)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"username":"hikaru"}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	body, err := c.get(t.Context(), "/player/hikaru")

	require.NoError(t, err)
	assert.Equal(t, `{"username":"hikaru"}`, string(body))
}

func TestGet_ETagCaching(t *testing.T) {
	t.Parallel()

	callCount := 0

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		if r.Header.Get("If-None-Match") == `"etag1"` {
			w.WriteHeader(http.StatusNotModified)

			return
		}

		w.Header().Set("ETag", `"etag1"`)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"online":true}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	body1, err := c.get(t.Context(), "/player/erik/is-online")

	require.NoError(t, err)
	assert.Equal(t, `{"online":true}`, string(body1))

	// Second call should get a 304 and return the cached body.
	body2, err := c.get(t.Context(), "/player/erik/is-online")

	require.NoError(t, err)
	assert.Equal(t, `{"online":true}`, string(body2))
	assert.Equal(t, 2, callCount, "server should have been called twice")
}

func TestGet_NotFound(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.get(t.Context(), "/player/doesnotexist")

	assert.ErrorIs(t, err, ErrNotFound)
}

func TestGet_Gone(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusGone)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.get(t.Context(), "/player/gone")

	assert.ErrorIs(t, err, ErrGone)
}

func TestGet_RateLimited(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.get(t.Context(), "/player/hikaru")

	assert.ErrorIs(t, err, ErrRateLimited)
}

func TestGet_APIError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.get(t.Context(), "/player/hikaru")

	require.Error(t, err)

	var apiErr *APIError

	assert.ErrorAs(t, err, &apiErr)
	assert.Equal(t, http.StatusInternalServerError, apiErr.StatusCode)
}

func TestGet_UserAgentHeader(t *testing.T) {
	t.Parallel()

	var receivedUA string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedUA = r.Header.Get("User-Agent")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL), WithUserAgent("myapp/2.0"))

	_, err := c.get(t.Context(), "/player/erik")

	require.NoError(t, err)
	assert.Equal(t, "myapp/2.0", receivedUA)
}
