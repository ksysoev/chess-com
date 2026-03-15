package chess

import (
	"context"
	"fmt"
)

// GetDailyPuzzle returns information about today's daily puzzle on chess.com.
//
// Endpoint: GET /pub/puzzle
func (c *Client) GetDailyPuzzle(ctx context.Context) (*DailyPuzzle, error) {
	var puzzle DailyPuzzle

	if err := c.getAndDecode(ctx, "/puzzle", &puzzle); err != nil {
		return nil, fmt.Errorf("get daily puzzle: %w", err)
	}

	return &puzzle, nil
}

// GetRandomPuzzle returns information about a randomly selected daily puzzle.
// Note that this endpoint has a caching latency of approximately 15 seconds,
// so it does not return a unique puzzle on every call.
//
// Endpoint: GET /pub/puzzle/random
func (c *Client) GetRandomPuzzle(ctx context.Context) (*DailyPuzzle, error) {
	var puzzle DailyPuzzle

	if err := c.getAndDecode(ctx, "/puzzle/random", &puzzle); err != nil {
		return nil, fmt.Errorf("get random puzzle: %w", err)
	}

	return &puzzle, nil
}
