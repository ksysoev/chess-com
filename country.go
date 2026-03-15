package chess

import (
	"context"
	"fmt"
)

// GetCountry returns the profile of the country identified by its ISO 3166-1
// alpha-2 code (e.g. "US", "IT", "DE"). Chess.com also supports custom codes
// for regions such as "XE" (England) and "XS" (Scotland).
//
// Endpoint: GET /pub/country/{iso}
func (c *Client) GetCountry(ctx context.Context, iso string) (*Country, error) {
	var country Country

	if err := c.getAndDecode(ctx, "/country/"+iso, &country); err != nil {
		return nil, fmt.Errorf("get country %q: %w", iso, err)
	}

	return &country, nil
}

// GetCountryPlayers returns the list of usernames for recently-active players
// who identify themselves as being in the given country. The ISO code must be
// an uppercase ISO 3166-1 alpha-2 code (e.g. "US").
//
// Endpoint: GET /pub/country/{iso}/players
func (c *Client) GetCountryPlayers(ctx context.Context, iso string) (*CountryPlayers, error) {
	var players CountryPlayers

	if err := c.getAndDecode(ctx, "/country/"+iso+"/players", &players); err != nil {
		return nil, fmt.Errorf("get country players %q: %w", iso, err)
	}

	return &players, nil
}

// GetCountryClubs returns the list of API URLs for clubs associated with the
// given country.
//
// Endpoint: GET /pub/country/{iso}/clubs
func (c *Client) GetCountryClubs(ctx context.Context, iso string) (*CountryClubs, error) {
	var clubs CountryClubs

	if err := c.getAndDecode(ctx, "/country/"+iso+"/clubs", &clubs); err != nil {
		return nil, fmt.Errorf("get country clubs %q: %w", iso, err)
	}

	return &clubs, nil
}
