package discogs

import (
	"fmt"
	"testing"
)

// TestReleaseGet fetches Release metadata based on the release id
func TestGetRelease(t *testing.T) {

	// Release Title we expect for release id = 23661
	expectedResult := "Smokers Delight"

	client := getClient(t)
	release, _, err := client.Database.GetRelease(23661)
	check(t, err)

	assert(t, release.Title == expectedResult, fmt.Sprintf("Release.Title expected %s, received %s ", expectedResult, release.Title))
}
