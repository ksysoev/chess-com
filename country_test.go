package chess

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCountry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *Country
		name       string
		iso        string
		statusCode int
	}{
		{
			name:       "success",
			iso:        "IT",
			statusCode: http.StatusOK,
			body:       Country{Name: "Italy", Code: "IT"},
			wantResult: &Country{Name: "Italy", Code: "IT"},
		},
		{
			name:       "not found",
			iso:        "ZZ",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/country/"+tt.iso, tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetCountry(t.Context(), tt.iso)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestGetCountryPlayers(t *testing.T) {
	t.Parallel()

	expected := &CountryPlayers{
		Players: []string{"alice", "bob", "charlie"},
	}

	srv := newTestServer(t, "/country/IT/players", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetCountryPlayers(t.Context(), "IT")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetCountryClubs(t *testing.T) {
	t.Parallel()

	expected := &CountryClubs{
		Clubs: []string{
			"https://api.chess.com/pub/club/club-italia",
		},
	}

	srv := newTestServer(t, "/country/IT/clubs", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetCountryClubs(t.Context(), "IT")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
