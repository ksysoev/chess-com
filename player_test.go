package chess

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, path string, statusCode int, body any) *httptest.Server {
	t.Helper()

	var payload []byte

	switch v := body.(type) {
	case string:
		payload = []byte(v)
	default:
		var err error

		payload, err = json.Marshal(body)
		require.NoError(t, err)
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path, r.URL.Path)
		w.WriteHeader(statusCode)
		_, _ = w.Write(payload)
	}))
}

func TestGetPlayer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body        any
		wantErr     error
		wantProfile *PlayerProfile
		name        string
		username    string
		statusCode  int
	}{
		{
			name:       "success",
			username:   "hikaru",
			statusCode: http.StatusOK,
			body: PlayerProfile{
				Username: "hikaru",
				PlayerID: 15448422,
				Status:   "premium",
			},
			wantProfile: &PlayerProfile{
				Username: "hikaru",
				PlayerID: 15448422,
				Status:   "premium",
			},
		},
		{
			name:       "not found",
			username:   "doesnotexist",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
		{
			name:       "rate limited",
			username:   "hikaru",
			statusCode: http.StatusTooManyRequests,
			body:       "",
			wantErr:    ErrRateLimited,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/player/"+tt.username, tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetPlayer(t.Context(), tt.username)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantProfile, got)
		})
	}
}

func TestGetTitledPlayers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *TitledPlayers
		name       string
		title      string
		statusCode int
	}{
		{
			name:       "success",
			title:      "GM",
			statusCode: http.StatusOK,
			body:       TitledPlayers{Players: []string{"hikaru", "magnuscarlsen"}},
			wantResult: &TitledPlayers{Players: []string{"hikaru", "magnuscarlsen"}},
		},
		{
			name:       "not found",
			title:      "INVALID",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/titled/"+tt.title, tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetTitledPlayers(t.Context(), tt.title)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestGetPlayerStats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *PlayerStats
		name       string
		username   string
		statusCode int
	}{
		{
			name:       "success",
			username:   "erik",
			statusCode: http.StatusOK,
			body: PlayerStats{
				ChessBlitz: &RatingStats{
					Last: &RatingPoint{Rating: 1500, RD: 50},
				},
			},
			wantResult: &PlayerStats{
				ChessBlitz: &RatingStats{
					Last: &RatingPoint{Rating: 1500, RD: 50},
				},
			},
		},
		{
			name:       "not found",
			username:   "ghost",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/player/"+tt.username+"/stats", tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetPlayerStats(t.Context(), tt.username)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestIsPlayerOnline(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		name       string
		username   string
		statusCode int
		wantOnline bool
	}{
		{
			name:       "online",
			username:   "erik",
			statusCode: http.StatusOK,
			body:       PlayerOnlineStatus{Online: true},
			wantOnline: true,
		},
		{
			name:       "offline",
			username:   "erik",
			statusCode: http.StatusOK,
			body:       PlayerOnlineStatus{Online: false},
			wantOnline: false,
		},
		{
			name:       "not found",
			username:   "ghost",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/player/"+tt.username+"/is-online", tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.IsPlayerOnline(t.Context(), tt.username)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantOnline, got)
		})
	}
}

func TestGetCurrentGames(t *testing.T) {
	t.Parallel()

	expected := &CurrentGames{
		Games: []CurrentGame{
			{URL: "https://www.chess.com/game/daily/1234", Turn: "white"},
		},
	}

	srv := newTestServer(t, "/player/erik/games", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetCurrentGames(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetGamesToMove(t *testing.T) {
	t.Parallel()

	expected := &ToMoveGames{
		Games: []ToMoveGame{
			{URL: "https://www.chess.com/game/daily/9999", MoveBy: 1700000000},
		},
	}

	srv := newTestServer(t, "/player/erik/games/to-move", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetGamesToMove(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetGameArchives(t *testing.T) {
	t.Parallel()

	expected := &GameArchives{
		Archives: []string{
			"https://api.chess.com/pub/player/erik/games/2009/10",
			"https://api.chess.com/pub/player/erik/games/2009/11",
		},
	}

	srv := newTestServer(t, "/player/erik/games/archives", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetGameArchives(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetMonthlyArchive(t *testing.T) {
	t.Parallel()

	expected := &MonthlyGames{
		Games: []Game{
			{URL: "https://www.chess.com/game/live/1111", Rules: "chess"},
		},
	}

	srv := newTestServer(t, "/player/erik/games/2009/10", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetMonthlyArchive(t.Context(), "erik", "2009", "10")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetMonthlyArchivePGN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		wantErr    error
		name       string
		statusCode int
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
		},
		{
			name:       "not found",
			statusCode: http.StatusNotFound,
			wantErr:    ErrNotFound,
		},
		{
			name:       "gone",
			statusCode: http.StatusGone,
			wantErr:    ErrGone,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			wantErr:    ErrRateLimited,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pgnData := "[Event \"Chess.com\"]\n1. e4 e5 *\n"

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)

				if tt.statusCode == http.StatusOK {
					_, _ = w.Write([]byte(pgnData))
				}
			}))
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetMonthlyArchivePGN(t.Context(), "erik", "2009", "10")

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, pgnData, got)
		})
	}
}

func TestGetPlayerClubs(t *testing.T) {
	t.Parallel()

	expected := &PlayerClubs{
		Clubs: []PlayerClubEntry{
			{Name: "Chess.com Developer Community", Joined: 1600000000},
		},
	}

	srv := newTestServer(t, "/player/erik/clubs", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetPlayerClubs(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetPlayerMatches(t *testing.T) {
	t.Parallel()

	expected := &PlayerMatches{
		Finished: []PlayerMatchEntry{
			{Name: "World League", Club: "https://api.chess.com/pub/club/team-usa"},
		},
		InProgress: []PlayerMatchEntry{},
		Registered: []PlayerMatchEntry{},
	}

	srv := newTestServer(t, "/player/erik/matches", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetPlayerMatches(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetPlayerTournaments(t *testing.T) {
	t.Parallel()

	expected := &PlayerTournaments{
		Finished: []PlayerTournamentEntry{
			{Status: "eliminated", Placement: 4, TotalPlayers: 5},
		},
		InProgress: []PlayerTournamentEntry{},
		Registered: []PlayerTournamentEntry{},
	}

	srv := newTestServer(t, "/player/erik/tournaments", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetPlayerTournaments(t.Context(), "erik")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
