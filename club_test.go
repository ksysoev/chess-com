package chess

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetClub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body       any
		wantErr    error
		wantResult *Club
		name       string
		urlID      string
		statusCode int
	}{
		{
			name:       "success",
			urlID:      "chess-com-developer-community",
			statusCode: http.StatusOK,
			body: Club{
				Name:         "Chess.com Developer Community",
				ClubID:       57796,
				MembersCount: 54,
			},
			wantResult: &Club{
				Name:         "Chess.com Developer Community",
				ClubID:       57796,
				MembersCount: 54,
			},
		},
		{
			name:       "not found",
			urlID:      "nonexistent-club",
			statusCode: http.StatusNotFound,
			body:       "",
			wantErr:    ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := newTestServer(t, "/club/"+tt.urlID, tt.statusCode, tt.body)
			defer srv.Close()

			c := New(WithBaseURL(srv.URL))

			got, err := c.GetClub(t.Context(), tt.urlID)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestGetClubMembers(t *testing.T) {
	t.Parallel()

	expected := &ClubMembers{
		Weekly: []ClubMember{
			{Username: "alice", Joined: 1600000000},
		},
		Monthly: []ClubMember{
			{Username: "bob", Joined: 1500000000},
		},
		AllTime: []ClubMember{
			{Username: "charlie", Joined: 1400000000},
		},
	}

	srv := newTestServer(t, "/club/my-club/members", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetClubMembers(t.Context(), "my-club")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetClubMatches(t *testing.T) {
	t.Parallel()

	expected := &ClubMatches{
		Finished: []ClubMatchEntry{
			{Name: "League Round 1", TimeClass: "daily", Result: "win"},
		},
		InProgress: []ClubMatchEntry{},
		Registered: []ClubMatchEntry{},
	}

	srv := newTestServer(t, "/club/my-club/matches", http.StatusOK, expected)
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))

	got, err := c.GetClubMatches(t.Context(), "my-club")

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
