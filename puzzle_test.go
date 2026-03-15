package chess

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDailyPuzzle(t *testing.T) {
	t.Parallel()

	expected := &DailyPuzzle{
		Title:       "Puzzle of the Day",
		URL:         "https://www.chess.com/daily-chess-puzzle/2024-01-01",
		FEN:         "r1bqkb1r/pppp1ppp/2n2n2/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R",
		PGN:         "1. e4 e5",
		Image:       "https://images.chesscomfiles.com/uploads/v1/puzzle/example.jpg",
		PublishTime: 1704067200,
	}

	srv := newTestServer(t, "/puzzle", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetDailyPuzzle(t.Context())

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetRandomPuzzle(t *testing.T) {
	t.Parallel()

	expected := &DailyPuzzle{
		Title:       "Random Puzzle",
		URL:         "https://www.chess.com/daily-chess-puzzle/2023-06-15",
		FEN:         "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR",
		PGN:         "1. e4",
		Image:       "https://images.chesscomfiles.com/uploads/v1/puzzle/random.jpg",
		PublishTime: 1686787200,
	}

	srv := newTestServer(t, "/puzzle/random", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetRandomPuzzle(t.Context())

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetDailyPuzzle_RateLimited(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t, "/puzzle", http.StatusTooManyRequests, "")
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	_, err := c.GetDailyPuzzle(t.Context())

	assert.ErrorIs(t, err, ErrRateLimited)
}
