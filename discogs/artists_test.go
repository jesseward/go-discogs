package discogs

import (
	"fmt"
	"testing"
)

// TestGetArtist fetches Artist metadata based on the artist id
func TestGetArtist(t *testing.T) {

	// Artist name we expect for artist id = 4130
	expectedResult := "Blunted Dummies"

	client := getClient(t)
	artist, _, err := client.Database.GetArtist(4130)
	check(t, err)

	assert(t, artist.Name == expectedResult, fmt.Sprintf("Artist.name expected %s, received %s ", expectedResult, artist.Name))
}

// TestArtistGetReleases fetches Artist Releases based on the artist id
func TestGetArtistReleases(t *testing.T) {

	// We expect Pagination.pages == 1 for artist is 4130
	expectedResult := 1

	client := getClient(t)
	releases, _, err := client.Database.GetArtistReleases(4130)
	check(t, err)

	assert(t, releases.Pagination.Pages == expectedResult, fmt.Sprintf("Artist.Pagination.Pages expected %d, received %d", expectedResult, releases.Pagination.Pages))

}
