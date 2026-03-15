package chesscom

import (
	"context"
	"fmt"
)

// GetClub returns the public profile of the club identified by the given
// url-ID. The url-ID is the identifier from the club's web URL on chess.com
// (e.g. "chess-com-developer-community").
//
// Endpoint: GET /pub/club/{url-ID}
func (c *Client) GetClub(ctx context.Context, urlID string) (*Club, error) {
	var club Club

	if err := c.getAndDecode(ctx, "/club/"+urlID, &club); err != nil {
		return nil, fmt.Errorf("get club %q: %w", urlID, err)
	}

	return &club, nil
}

// GetClubMembers returns the members of the given club grouped by activity
// level: weekly, monthly, and all-time.
//
// Endpoint: GET /pub/club/{url-ID}/members
func (c *Client) GetClubMembers(ctx context.Context, urlID string) (*ClubMembers, error) {
	var members ClubMembers

	if err := c.getAndDecode(ctx, "/club/"+urlID+"/members", &members); err != nil {
		return nil, fmt.Errorf("get club members %q: %w", urlID, err)
	}

	return &members, nil
}

// GetClubMatches returns the list of team matches for the given club, grouped
// by status: finished, in-progress, and registered.
//
// Endpoint: GET /pub/club/{url-ID}/matches
func (c *Client) GetClubMatches(ctx context.Context, urlID string) (*ClubMatches, error) {
	var matches ClubMatches

	if err := c.getAndDecode(ctx, "/club/"+urlID+"/matches", &matches); err != nil {
		return nil, fmt.Errorf("get club matches %q: %w", urlID, err)
	}

	return &matches, nil
}
