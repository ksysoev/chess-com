package chesscom

import (
	"context"
	"fmt"
)

// GetTournament returns details about the tournament identified by the given
// url-ID. The url-ID is the identifier from the tournament's web URL on
// chess.com (e.g. "-33rd-chesscom-quick-knockouts-1401-1600").
//
// Endpoint: GET /pub/tournament/{url-ID}
func (c *Client) GetTournament(ctx context.Context, urlID string) (*Tournament, error) {
	var tournament Tournament

	if err := c.getAndDecode(ctx, "/tournament/"+urlID, &tournament); err != nil {
		return nil, fmt.Errorf("get tournament %q: %w", urlID, err)
	}

	return &tournament, nil
}

// GetTournamentRound returns details about the specified round of the given
// tournament.
//
// Endpoint: GET /pub/tournament/{url-ID}/{round}
func (c *Client) GetTournamentRound(ctx context.Context, urlID string, round int) (*TournamentRound, error) {
	var tournamentRound TournamentRound

	path := fmt.Sprintf("/tournament/%s/%d", urlID, round)

	if err := c.getAndDecode(ctx, path, &tournamentRound); err != nil {
		return nil, fmt.Errorf("get tournament round %q/%d: %w", urlID, round, err)
	}

	return &tournamentRound, nil
}

// GetTournamentRoundGroup returns details about a specific group within a
// tournament round, including games and player standings.
//
// Endpoint: GET /pub/tournament/{url-ID}/{round}/{group}
func (c *Client) GetTournamentRoundGroup(ctx context.Context, urlID string, round, group int) (*TournamentRoundGroup, error) {
	var roundGroup TournamentRoundGroup

	path := fmt.Sprintf("/tournament/%s/%d/%d", urlID, round, group)

	if err := c.getAndDecode(ctx, path, &roundGroup); err != nil {
		return nil, fmt.Errorf("get tournament round group %q/%d/%d: %w", urlID, round, group, err)
	}

	return &roundGroup, nil
}
