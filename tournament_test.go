package chess

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTournament(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *Tournament
		name       string
		urlID      string
		statusCode int
	}{
		{
			name:       "success",
			urlID:      "-33rd-chesscom-quick-knockouts-1401-1600",
			statusCode: http.StatusOK,
			body: Tournament{
				Name:   "33rd Chess.com Quick Knockouts 1401-1600",
				Status: "finished",
			},
			wantResult: &Tournament{
				Name:   "33rd Chess.com Quick Knockouts 1401-1600",
				Status: "finished",
			},
		},
		{
			name:       "not found",
			urlID:      "nonexistent",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/tournament/"+tt.urlID, tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetTournament(t.Context(), tt.urlID)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestGetTournamentRound(t *testing.T) {
	t.Parallel()

	expected := &TournamentRound{
		Groups: []string{"https://api.chess.com/pub/tournament/my-tournament/1/1"},
		Players: []TournamentRoundPlayer{
			{Username: "alice", IsAdvancing: true},
		},
	}

	srv := newTestServer(t, "/tournament/my-tournament/1", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetTournamentRound(t.Context(), "my-tournament", 1)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetTournamentRoundGroup(t *testing.T) {
	t.Parallel()

	expected := &TournamentRoundGroup{
		FairPlayRemovals: []string{},
		Players: []TournamentGroupPlayer{
			{Username: "alice", Points: 2, TieBreak: 6, IsAdvancing: true},
		},
		Games: []CurrentGame{},
	}

	srv := newTestServer(t, "/tournament/my-tournament/1/1", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetTournamentRoundGroup(t.Context(), "my-tournament", 1, 1)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
