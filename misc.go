package chesscom

import (
	"context"
	"fmt"
)

// GetStreamers returns the list of Chess.com streamers and their details.
// This endpoint refreshes every 5 minutes.
//
// Endpoint: GET /pub/streamers
func (c *Client) GetStreamers(ctx context.Context) (*StreamersList, error) {
	var streamers StreamersList

	if err := c.getAndDecode(ctx, "/streamers", &streamers); err != nil {
		return nil, fmt.Errorf("get streamers: %w", err)
	}

	return &streamers, nil
}

// GetLeaderboards returns the top 50 players for each game type, tactics,
// and lessons categories.
//
// Endpoint: GET /pub/leaderboards
func (c *Client) GetLeaderboards(ctx context.Context) (*Leaderboards, error) {
	var leaderboards Leaderboards

	if err := c.getAndDecode(ctx, "/leaderboards", &leaderboards); err != nil {
		return nil, fmt.Errorf("get leaderboards: %w", err)
	}

	return &leaderboards, nil
}
