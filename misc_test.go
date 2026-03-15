package chesscom

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStreamers(t *testing.T) {
	t.Parallel()

	expected := &StreamersList{
		Streamers: []Streamer{
			{
				Username:  "hikaru",
				Avatar:    "https://images.chesscomfiles.com/uploads/v1/user/15448422.jpeg",
				TwitchURL: "https://twitch.tv/gmhikaru",
				URL:       "https://www.chess.com/member/hikaru",
			},
		},
	}

	srv := newTestServer(t, "/streamers", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetStreamers(t.Context())

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetStreamers_RateLimited(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t, "/streamers", http.StatusTooManyRequests, "")
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.GetStreamers(t.Context())

	assert.ErrorIs(t, err, ErrRateLimited)
}

func TestGetLeaderboards(t *testing.T) {
	t.Parallel()

	expected := &Leaderboards{
		LiveBlitz: []LeaderboardEntry{
			{Username: "magnuscarlsen", Score: 3200, Rank: 1},
		},
		Daily: []LeaderboardEntry{
			{Username: "erik", Score: 2800, Rank: 1},
		},
	}

	srv := newTestServer(t, "/leaderboards", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetLeaderboards(t.Context())

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetLeaderboards_RateLimited(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t, "/leaderboards", http.StatusTooManyRequests, "")
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.GetLeaderboards(t.Context())

	assert.ErrorIs(t, err, ErrRateLimited)
}
