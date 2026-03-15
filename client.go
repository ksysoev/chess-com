// Package chess provides a client for the Chess.com public API (PubAPI).
//
// The PubAPI is a read-only REST API that exposes public data from Chess.com
// such as player profiles, game histories, club information, tournaments,
// leaderboards, and more.
//
// Usage:
//
//	client := chess.New()
//	profile, err := client.GetPlayer(ctx, "hikaru")
//
// The client supports functional options for customisation:
//
//	client := chess.New(
//	    chess.WithTimeout(10 * time.Second),
//	    chess.WithUserAgent("myapp/1.0 (contact@example.com)"),
//	)
package chess

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	defaultBaseURL   = "https://api.chess.com/pub"
	defaultUserAgent = "github.com/ksysoev/chess-com"
	defaultTimeout   = 30 * time.Second
)

// cacheEntry holds an ETag and the raw response body for a cached response.
type cacheEntry struct {
	etag string
	body []byte
}

// Client is the Chess.com public API client.
// The zero value is not usable; use New to create a client.
type Client struct {
	httpClient *http.Client
	cache      map[string]cacheEntry
	baseURL    string
	userAgent  string
	mu         sync.RWMutex
}

// Option is a functional option for configuring a Client.
type Option func(*Client)

// WithHTTPClient sets a custom *http.Client to use for requests.
// This allows callers to configure transport-level settings such as
// TLS configuration, proxies, or custom round-trippers.
func WithHTTPClient(c *http.Client) Option {
	return func(cl *Client) {
		cl.httpClient = c
	}
}

// WithBaseURL overrides the base URL of the Chess.com API.
// This is primarily useful for testing against a mock server.
func WithBaseURL(url string) Option {
	return func(cl *Client) {
		cl.baseURL = url
	}
}

// WithUserAgent sets the User-Agent header sent with every request.
// The Chess.com API recommends including contact information so they can
// reach out if your application is blocked.
func WithUserAgent(ua string) Option {
	return func(cl *Client) {
		cl.userAgent = ua
	}
}

// WithTimeout sets the HTTP request timeout on the default HTTP client.
// This option has no effect if WithHTTPClient is also used.
func WithTimeout(d time.Duration) Option {
	return func(cl *Client) {
		cl.httpClient.Timeout = d
	}
}

// New creates and returns a new Chess.com API client.
// Options are applied in order; later options override earlier ones.
func New(opts ...Option) *Client {
	c := &Client{
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
		cache:     make(map[string]cacheEntry),
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// get performs an HTTP GET request to the given path (relative to the base URL),
// honouring ETag-based caching. On a 304 Not Modified response it returns the
// previously cached body. On success it updates the cache and returns the body.
func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	c.mu.RLock()
	entry, cached := c.cache[url]
	c.mu.RUnlock()

	if cached {
		req.Header.Set("If-None-Match", entry.etag)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotModified:
		return entry.body, nil
	case http.StatusOK:
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("read response body: %w", readErr)
		}

		if etag := resp.Header.Get("ETag"); etag != "" {
			c.mu.Lock()
			c.cache[url] = cacheEntry{etag: etag, body: body}
			c.mu.Unlock()
		}

		return body, nil
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusGone:
		return nil, ErrGone
	case http.StatusTooManyRequests:
		return nil, ErrRateLimited
	default:
		return nil, newAPIError(resp.StatusCode, resp.Status)
	}
}

// getAndDecode performs a GET request and JSON-decodes the response into dst.
func (c *Client) getAndDecode(ctx context.Context, path string, dst any) error {
	body, err := c.get(ctx, path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, dst); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}
