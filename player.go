package chess

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// GetPlayer returns the public profile of the player with the given username.
//
// Endpoint: GET /pub/player/{username}
func (c *Client) GetPlayer(ctx context.Context, username string) (*PlayerProfile, error) {
	var profile PlayerProfile

	if err := c.getAndDecode(ctx, "/player/"+username, &profile); err != nil {
		return nil, fmt.Errorf("get player %q: %w", username, err)
	}

	return &profile, nil
}

// GetTitledPlayers returns the list of usernames for players holding the given
// chess title. Valid title abbreviations are: GM, WGM, IM, WIM, FM, WFM, NM,
// WNM, CM, WCM.
//
// Endpoint: GET /pub/titled/{title-abbrev}
func (c *Client) GetTitledPlayers(ctx context.Context, title string) (*TitledPlayers, error) {
	var players TitledPlayers

	if err := c.getAndDecode(ctx, "/titled/"+title, &players); err != nil {
		return nil, fmt.Errorf("get titled players %q: %w", title, err)
	}

	return &players, nil
}

// GetPlayerStats returns ratings, win/loss records, and other stats for the
// given player across all game types they have played.
//
// Endpoint: GET /pub/player/{username}/stats
func (c *Client) GetPlayerStats(ctx context.Context, username string) (*PlayerStats, error) {
	var stats PlayerStats

	if err := c.getAndDecode(ctx, "/player/"+username+"/stats", &stats); err != nil {
		return nil, fmt.Errorf("get player stats %q: %w", username, err)
	}

	return &stats, nil
}

// IsPlayerOnline reports whether the player with the given username has been
// online in the last five minutes.
//
// Endpoint: GET /pub/player/{username}/is-online
func (c *Client) IsPlayerOnline(ctx context.Context, username string) (bool, error) {
	var status PlayerOnlineStatus

	if err := c.getAndDecode(ctx, "/player/"+username+"/is-online", &status); err != nil {
		return false, fmt.Errorf("get player online status %q: %w", username, err)
	}

	return status.Online, nil
}

// GetCurrentGames returns the daily chess games that the given player is
// currently playing.
//
// Endpoint: GET /pub/player/{username}/games
func (c *Client) GetCurrentGames(ctx context.Context, username string) (*CurrentGames, error) {
	var games CurrentGames

	if err := c.getAndDecode(ctx, "/player/"+username+"/games", &games); err != nil {
		return nil, fmt.Errorf("get current games %q: %w", username, err)
	}

	return &games, nil
}

// GetGamesToMove returns the daily chess games where it is the given player's
// turn to act.
//
// Endpoint: GET /pub/player/{username}/games/to-move
func (c *Client) GetGamesToMove(ctx context.Context, username string) (*ToMoveGames, error) {
	var games ToMoveGames

	if err := c.getAndDecode(ctx, "/player/"+username+"/games/to-move", &games); err != nil {
		return nil, fmt.Errorf("get games to move %q: %w", username, err)
	}

	return &games, nil
}

// GetGameArchives returns the list of monthly archive URLs available for the
// given player, in ascending chronological order.
//
// Endpoint: GET /pub/player/{username}/games/archives
func (c *Client) GetGameArchives(ctx context.Context, username string) (*GameArchives, error) {
	var archives GameArchives

	if err := c.getAndDecode(ctx, "/player/"+username+"/games/archives", &archives); err != nil {
		return nil, fmt.Errorf("get game archives %q: %w", username, err)
	}

	return &archives, nil
}

// GetMonthlyArchive returns all completed games (both live and daily) for the
// given player in the specified month. year must be a four-digit year and month
// must be a two-digit month (e.g. "2024", "01").
//
// Endpoint: GET /pub/player/{username}/games/{YYYY}/{MM}
func (c *Client) GetMonthlyArchive(ctx context.Context, username, year, month string) (*MonthlyGames, error) {
	var games MonthlyGames

	path := fmt.Sprintf("/player/%s/games/%s/%s", username, year, month)

	if err := c.getAndDecode(ctx, path, &games); err != nil {
		return nil, fmt.Errorf("get monthly archive %q %s/%s: %w", username, year, month, err)
	}

	return &games, nil
}

// GetMonthlyArchivePGN downloads the multi-game PGN file containing all games
// for the given player in the specified month. year must be a four-digit year
// and month must be a two-digit month (e.g. "2024", "01").
//
// Endpoint: GET /pub/player/{username}/games/{YYYY}/{MM}/pgn
func (c *Client) GetMonthlyArchivePGN(ctx context.Context, username, year, month string) (string, error) {
	url := fmt.Sprintf("%s/player/%s/games/%s/%s/pgn", c.baseURL, username, year, month)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", fmt.Errorf("create pgn request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("execute pgn request: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// intentionally fall through to read body
	case http.StatusNotFound:
		return "", ErrNotFound
	case http.StatusGone:
		return "", ErrGone
	case http.StatusTooManyRequests:
		return "", ErrRateLimited
	default:
		return "", newAPIError(resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read pgn response: %w", err)
	}

	return string(body), nil
}

// GetPlayerClubs returns the list of clubs that the given player is a member of,
// including the join date and last activity timestamp for each club.
//
// Endpoint: GET /pub/player/{username}/clubs
func (c *Client) GetPlayerClubs(ctx context.Context, username string) (*PlayerClubs, error) {
	var clubs PlayerClubs

	if err := c.getAndDecode(ctx, "/player/"+username+"/clubs", &clubs); err != nil {
		return nil, fmt.Errorf("get player clubs %q: %w", username, err)
	}

	return &clubs, nil
}

// GetPlayerMatches returns the list of team matches the given player has
// participated in, is currently playing, or is registered for.
//
// Endpoint: GET /pub/player/{username}/matches
func (c *Client) GetPlayerMatches(ctx context.Context, username string) (*PlayerMatches, error) {
	var matches PlayerMatches

	if err := c.getAndDecode(ctx, "/player/"+username+"/matches", &matches); err != nil {
		return nil, fmt.Errorf("get player matches %q: %w", username, err)
	}

	return &matches, nil
}

// GetPlayerTournaments returns the list of tournaments the given player has
// participated in, is currently playing, or is registered for.
//
// Endpoint: GET /pub/player/{username}/tournaments
func (c *Client) GetPlayerTournaments(ctx context.Context, username string) (*PlayerTournaments, error) {
	var tournaments PlayerTournaments

	if err := c.getAndDecode(ctx, "/player/"+username+"/tournaments", &tournaments); err != nil {
		return nil, fmt.Errorf("get player tournaments %q: %w", username, err)
	}

	return &tournaments, nil
}
