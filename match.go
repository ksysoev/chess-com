package chesscom

import (
	"context"
	"fmt"
)

// GetTeamMatch returns details about the daily team match with the given
// numeric ID. The ID is the same as found in the match's web URL on chess.com.
//
// Endpoint: GET /pub/match/{ID}
func (c *Client) GetTeamMatch(ctx context.Context, id int) (*TeamMatch, error) {
	var match TeamMatch

	if err := c.getAndDecode(ctx, fmt.Sprintf("/match/%d", id), &match); err != nil {
		return nil, fmt.Errorf("get team match %d: %w", id, err)
	}

	return &match, nil
}

// GetTeamMatchBoard returns the game details for a specific board in the
// given daily team match. Only in-progress or finished games are included.
//
// Endpoint: GET /pub/match/{ID}/{board}
func (c *Client) GetTeamMatchBoard(ctx context.Context, id, board int) (*MatchBoard, error) {
	var matchBoard MatchBoard

	if err := c.getAndDecode(ctx, fmt.Sprintf("/match/%d/%d", id, board), &matchBoard); err != nil {
		return nil, fmt.Errorf("get team match board %d/%d: %w", id, board, err)
	}

	return &matchBoard, nil
}

// GetLiveTeamMatch returns details about the live team match with the given
// numeric ID.
//
// Endpoint: GET /pub/match/live/{ID}
func (c *Client) GetLiveTeamMatch(ctx context.Context, id int) (*LiveTeamMatch, error) {
	var match LiveTeamMatch

	if err := c.getAndDecode(ctx, fmt.Sprintf("/match/live/%d", id), &match); err != nil {
		return nil, fmt.Errorf("get live team match %d: %w", id, err)
	}

	return &match, nil
}

// GetLiveTeamMatchBoard returns the game details for a specific board in the
// given live team match.
//
// Endpoint: GET /pub/match/live/{ID}/{board}
func (c *Client) GetLiveTeamMatchBoard(ctx context.Context, id, board int) (*LiveMatchBoard, error) {
	var matchBoard LiveMatchBoard

	if err := c.getAndDecode(ctx, fmt.Sprintf("/match/live/%d/%d", id, board), &matchBoard); err != nil {
		return nil, fmt.Errorf("get live team match board %d/%d: %w", id, board, err)
	}

	return &matchBoard, nil
}
