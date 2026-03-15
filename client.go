// Package chesscom provides a client for the Chess.com public API (PubAPI).
//
// The PubAPI is a read-only REST API that exposes public data from Chess.com
// such as player profiles, game histories, club information, tournaments,
// leaderboards, and more.
//
// Usage:
//
//	client := chesscom.New()
//	profile, err := client.GetPlayer(ctx, "hikaru")
//
// The client supports functional options for customisation:
//
//	client := chesscom.New(
//	    chesscom.WithTimeout(10 * time.Second),
//	    chesscom.WithUserAgent("myapp/1.0 (contact@example.com)"),
//	)
package chesscom

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL   = "https://api.chess.com/pub"
	defaultUserAgent = "github.com/ksysoev/chess-com"
	defaultTimeout   = 30 * time.Second
)

// Client is the Chess.com public API client.
// The zero value is not usable; use New to create a client.
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

// config holds transient build-time configuration collected during New().
type config struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
	timeout    time.Duration
}

// Option is a functional option for configuring a Client.
type Option func(*config)

// WithHTTPClient sets a custom *http.Client to use for requests.
// This allows callers to configure transport-level settings such as
// TLS configuration, proxies, or custom round-trippers.
// When this option is used, WithTimeout has no effect.
func WithHTTPClient(c *http.Client) Option {
	return func(cfg *config) {
		cfg.httpClient = c
	}
}

// WithBaseURL overrides the base URL of the Chess.com API.
// This is primarily useful for testing against a mock server.
func WithBaseURL(url string) Option {
	return func(cfg *config) {
		cfg.baseURL = url
	}
}

// WithUserAgent sets the User-Agent header sent with every request.
// The Chess.com API recommends including contact information so they can
// reach out if your application is blocked.
func WithUserAgent(ua string) Option {
	return func(cfg *config) {
		cfg.userAgent = ua
	}
}

// WithTimeout sets the HTTP request timeout for the default HTTP client.
// This option has no effect when WithHTTPClient is also provided, regardless
// of the order in which the options are applied.
func WithTimeout(d time.Duration) Option {
	return func(cfg *config) {
		cfg.timeout = d
	}
}

// New creates and returns a new Chess.com API client.
// Options are applied in order; later options override earlier ones.
// If WithTimeout is provided without WithHTTPClient, the timeout is applied
// to the default HTTP client. If both are provided, the custom client is used
// as-is and WithTimeout has no effect.
func New(opts ...Option) *Client {
	cfg := &config{
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	c := &Client{
		baseURL:   cfg.baseURL,
		userAgent: cfg.userAgent,
	}

	switch {
	case cfg.httpClient != nil:
		c.httpClient = cfg.httpClient
	case cfg.timeout != 0:
		c.httpClient = &http.Client{Timeout: cfg.timeout}
	default:
		c.httpClient = &http.Client{Timeout: defaultTimeout}
	}

	return c
}

// get performs an HTTP GET request to the given path (relative to the base URL)
// and returns the response body bytes.
func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("read response body: %w", readErr)
		}

		return body, nil
	case http.StatusNotFound:
		_, _ = io.Copy(io.Discard, resp.Body)

		return nil, ErrNotFound
	case http.StatusGone:
		_, _ = io.Copy(io.Discard, resp.Body)

		return nil, ErrGone
	case http.StatusTooManyRequests:
		_, _ = io.Copy(io.Discard, resp.Body)

		return nil, ErrRateLimited
	default:
		_, _ = io.Copy(io.Discard, resp.Body)

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
