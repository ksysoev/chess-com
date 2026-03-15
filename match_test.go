package chess

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamMatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *TeamMatch
		name       string
		id         int
		statusCode int
	}{
		{
			name:       "success",
			id:         12803,
			statusCode: http.StatusOK,
			body: TeamMatch{
				Name:   "WORLD LEAGUE Round 5",
				Status: "finished",
				Boards: 8,
			},
			wantResult: &TeamMatch{
				Name:   "WORLD LEAGUE Round 5",
				Status: "finished",
				Boards: 8,
			},
		},
		{
			name:       "not found",
			id:         99999,
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/match/12803", tt.statusCode, tt.body)

			if tt.id != 12803 {
				srv.Close()
				srv = newTestServer(t, "/match/99999", tt.statusCode, tt.body)
			}

			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetTeamMatch(t.Context(), tt.id)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestGetTeamMatchBoard(t *testing.T) {
	t.Parallel()

	expected := &MatchBoard{
		BoardScores: map[string]float64{"player1": 0.5, "player2": 1.5},
		Games:       []MatchBoardGame{},
	}

	srv := newTestServer(t, "/match/12803/1", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetTeamMatchBoard(t.Context(), 12803, 1)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetLiveTeamMatch(t *testing.T) {
	t.Parallel()

	expected := &LiveTeamMatch{
		Name:   "Friendly 5+2",
		Status: "finished",
		Boards: 6,
	}

	srv := newTestServer(t, "/match/live/5833", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetLiveTeamMatch(t.Context(), 5833)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetLiveTeamMatchBoard(t *testing.T) {
	t.Parallel()

	expected := &LiveMatchBoard{
		BoardScores: map[string]float64{"stompall": 1.5, "jydra21": 0.5},
		Games:       []LiveMatchBoardGame{},
	}

	srv := newTestServer(t, "/match/live/5833/5", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetLiveTeamMatchBoard(t.Context(), 5833, 5)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
